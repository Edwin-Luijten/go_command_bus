# A Command Bus for GO

## Installation
``` go get github.com/edwin-luijten/go-command-bus@v0.0 ```

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