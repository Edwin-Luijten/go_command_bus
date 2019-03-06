package go_command_bus

import (
	"reflect"
	"sort"
	"sync"
)

type HandlerFunc func(command interface{})
type middlewareFunc func(command interface{}, next HandlerFunc)

type middleware struct {
	function middlewareFunc
	priority int
}

type CommandBus struct {
	lock        sync.Mutex
	handlers    map[reflect.Type]HandlerFunc
	middlewares []middleware
}

func New() *CommandBus {
	return &CommandBus{
		handlers:    make(map[reflect.Type]HandlerFunc),
		middlewares: make([]middleware, 0),
	}
}

func (b *CommandBus) RegisterHandler(command interface{}, handler HandlerFunc) {
	b.lock.Lock()

	defer b.lock.Unlock()

	b.handlers[reflect.TypeOf(command)] = handler
}

func (b *CommandBus) RegisterMiddleware(function middlewareFunc, priority int) {
	b.lock.Lock()

	defer b.lock.Unlock()

	b.middlewares = append(b.middlewares, middleware{function: function, priority: priority})

	sort.Sort(sortByPriority(b.middlewares))
}

func (b CommandBus) Handle(command interface{}) {
	b.lock.Lock()

	defer b.lock.Unlock()

	b.getNext(0)(command)
}

func (b CommandBus) GetHandler(command interface{}) HandlerFunc {
	handler, _ := b.handlers[reflect.TypeOf(command)]

	return handler
}

func (b CommandBus) getNext(next int) HandlerFunc {
	if len(b.middlewares) >= (next + 1) {
		return func(command interface{}) {
			middleware := b.middlewares[next]
			middleware.function(command, b.getNext(next+1))
		}
	}

	return func(command interface{}) {
		if handler := b.GetHandler(command); handler != nil {
			handler(command)
		}
	}
}
