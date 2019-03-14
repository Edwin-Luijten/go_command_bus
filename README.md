# A Command Bus for GO

[![Build Status](https://travis-ci.com/Edwin-Luijten/go_command_bus.svg?branch=master)](https://travis-ci.com/Edwin-Luijten/go_command_bus) 
[![Maintainability](https://api.codeclimate.com/v1/badges/ff5d37cbc59ef9a174a5/maintainability)](https://codeclimate.com/github/Edwin-Luijten/go_command_bus/maintainability) 
[![Test Coverage](https://api.codeclimate.com/v1/badges/ff5d37cbc59ef9a174a5/test_coverage)](https://codeclimate.com/github/Edwin-Luijten/go_command_bus/test_coverage) 
[![GoDoc](https://godoc.org/github.com/Edwin-Luijten/go_command_bus?status.svg)](https://godoc.org/github.com/Edwin-Luijten/go_command_bus)  

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