# A Command Bus for GO

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/edwin-luijten/go_command_bus?style=flat-square) 
[![Go Reference](https://pkg.go.dev/badge/github.com/Edwin-Luijten/go_command_bus.svg)](https://pkg.go.dev/github.com/Edwin-Luijten/go_command_bus)
[![Build Status](https://travis-ci.com/Edwin-Luijten/go_command_bus.svg?branch=master)](https://travis-ci.com/Edwin-Luijten/go_command_bus) 
[![Maintainability](https://api.codeclimate.com/v1/badges/ff5d37cbc59ef9a174a5/maintainability)](https://codeclimate.com/github/Edwin-Luijten/go_command_bus/maintainability) 
[![Test Coverage](https://api.codeclimate.com/v1/badges/ff5d37cbc59ef9a174a5/test_coverage)](https://codeclimate.com/github/Edwin-Luijten/go_command_bus/test_coverage)
## Installation
``` go get github.com/edwin-luijten/go_command_bus ```

## Usage

### Commands
```go
import (
    commandbus "github.com/edwin-luijten/go_command_bus"
)

func main() {
    bus := commandbus.New()
    
    bus.RegisterHandler(&RegisterUserCommand{}, func(command interface{}) {
        cmd := command.(*RegisterUserCommand)
        
        fmt.Sprintf("Registered %s", cmd.Username)
    })
    
    bus.Handle(&RegisterUserCommand{
    	Email: "jj@email.com",
    	Password: "secret",
    })
}
```

### Middleware's

A middleware has a handler and a priority.  
Where priority of 0 is the least amount of priority.  
```go
import (
    commandbus "github.com/edwin-luijten/go_command_bus"
)

func main() {
    bus := commandbus.New()
    
    bus.RegisterMiddleware(func(command interface{}, next HandlerFunc) {
        
    	// your logic here
    
        next(command)
    
        // or here
    }, 1) 
}
```