package commandbus

import (
	"bytes"
	"fmt"
	"testing"
)

type Command1 struct {
	Message string
}
type Command2 struct{}

func ExampleCommandBus_Commands() {
	bus := New()

	bus.RegisterHandler(&Command1{}, func(command interface{}) {
		cmd := command.(*Command1)

		fmt.Println(cmd.Message)
		// Output: yay
	})

	bus.Handle(&Command1{
		Message: "yay",
	})
}

func ExampleCommandBus_Middlewares() {
	bus := New()

	bus.RegisterHandler(&Command1{}, func(command interface{}) {
		cmd := command.(*Command1)

		fmt.Println(cmd.Message)
		// Output: no
	})

	bus.RegisterMiddleware(func(command interface{}, next HandlerFunc) {
		switch command.(type) {
		case Command1:
			command.(*Command1).Message = "no"
		}

		next(command)
	}, 1)

	bus.Handle(&Command1{
		Message: "yay",
	})
}

func TestGetRegisteredHandler(t *testing.T) {
	bus := New()

	test := 0

	bus.RegisterHandler(&Command1{}, func(command interface{}) {
		test = 1
	})

	bus.RegisterHandler(&Command2{}, func(command interface{}) {
		test = 2
	})

	bus.GetHandler(&Command1{})(nil)

	if test != 1 {
		t.Log("Wrong handler")
		t.Fail()
	}

	bus.GetHandler(&Command2{})(nil)

	if test != 2 {
		t.Log("Wrong handler")
		t.Fail()
	}
}

func TestHandler(t *testing.T) {
	var buff bytes.Buffer

	bus := New()

	bus.RegisterHandler(&Command1{}, func(command interface{}) {
		cmd := command.(*Command1)
		buff.WriteString(cmd.Message)
	})

	bus.Handle(&Command1{
		Message: "yay",
	})

	if buff.String() != "yay" {
		t.Log("Command failed")
		t.Fail()
	}
}

func TestMiddleware(t *testing.T) {
	var buff bytes.Buffer

	bus := New()

	bus.RegisterMiddleware(func(command interface{}, next HandlerFunc) {
		buff.WriteString("1")

		next(command)

		buff.WriteString("1")
	}, 0)

	command := &Command1{}
	bus.RegisterHandler(command, func(command interface{}) {
		buff.WriteString("yay")
	})

	bus.Handle(command)

	if buff.String() != "1yay1" {
		t.Log("Middleware failed")
		t.Fail()
	}
}

func TestPrioritizedMiddlewares(t *testing.T) {
	var buff bytes.Buffer

	bus := New()

	// higher priority
	bus.RegisterMiddleware(func(command interface{}, next HandlerFunc) {
		buff.WriteString("1")

		next(command)

		buff.WriteString("1")
	}, 1)

	// lower priority
	bus.RegisterMiddleware(func(command interface{}, next HandlerFunc) {
		buff.WriteString("2")

		next(command)

		buff.WriteString("2")
	}, 0)

	command := &Command1{}
	bus.RegisterHandler(command, func(command interface{}) {
		buff.WriteString("yay")
	})

	bus.Handle(command)

	if buff.String() != "12yay21" {
		t.Log("Middleware failed")
		t.Fail()
	}
}
