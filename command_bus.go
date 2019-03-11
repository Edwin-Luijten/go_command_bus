// Package commandbus provides an easy way to implement the command bus pattern.
package commandbus

import (
	"reflect"
	"sort"
	"sync"
)

// HandlerFunc provides command handler logic
type HandlerFunc func(command interface{})

type middlewareFunc func(command interface{}, next HandlerFunc)

type middleware struct {
	function middlewareFunc
	priority int
}

// CommandBus ...
type CommandBus struct {
	lock        *sync.Mutex
	handlers    map[reflect.Type]HandlerFunc
	middlewares []middleware
}

// New creates a new instance that returns a CommandBus
func New() *CommandBus {
	return &CommandBus{
		lock:        &sync.Mutex{},
		handlers:    make(map[reflect.Type]HandlerFunc),
		middlewares: make([]middleware, 0),
	}
}

// RegisterHandler allows you to register a command and it's handler.
func (b *CommandBus) RegisterHandler(command interface{}, handler HandlerFunc) {
	b.lock.Lock()

	defer b.lock.Unlock()

	b.handlers[reflect.TypeOf(command)] = handler
}

// RegisterMiddleware allows you to register middleware's.
// Use the priority argument to control the order of execution.
func (b *CommandBus) RegisterMiddleware(function middlewareFunc, priority int) {
	b.lock.Lock()

	defer b.lock.Unlock()

	b.middlewares = append(b.middlewares, middleware{function: function, priority: priority})

	sort.Sort(sortByPriority(b.middlewares))
}

// Handle allows you to trigger a command.
func (b CommandBus) Handle(command interface{}) {
	b.lock.Lock()

	defer b.lock.Unlock()

	b.getNext(0)(command)
}

// GetHandler returns the command handler that is registered to a given command.
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
