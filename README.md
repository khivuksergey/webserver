# webserver
Simple go webserver with built-in graceful shutdown.

## Features
- Built-in server graceful shutdown with ability to shut down other necessary components, such as database connections.
- Configurable router implementation with [Echo](https://github.com/labstack/echo).
- Simple console logger implementation.

## How to use
### Initialize and run server
```go
import (
    "github.com/khivuksergey/webserver"
    "github.com/khivuksergey/webserver/router"
    "github.com/labstack/echo/v4"
    "net/http"
    "os"
)

func main() {
    // Create HTTP router
    r := router.NewEchoRouter()
    r.GET("/hello", func (c echo.Context) error {
        return c.String(http.StatusOK, "world!")
    })
    
    // Initialize server with default config and pass created router
    server := webserver.NewServer(r)
    
    // Create a channel of type os.Signal
    quit := make(chan os.Signal, 1)
    
    // Pass created Server and channel to RunServer function
    if err := webserver.RunServer(server, quit); err != nil {
        panic(err)
    }
}
```

### Configure server
```go
// Create server with provided HTTP handler
server := webserver.NewServer(router).
    // Add configuration
    WithConfig(&config.Server)
    // Add logger
    AddLogger(logger).
    // Add handlers to shut down components gracefully
    AddStopHandlers(
        NewStopHandler("Postgres", db.Close),
        NewStopHandler("KafkaConsumer", kafka.Close),
    )
```
#### Server configuration
```go
type ServerConfig struct {
    Port        int
    StopTimeout time.Duration
}
```

### Configure router
```go
// Create default echo router
e := NewEchoRouter().
    // Optionally add echo's logger and recovery middleware
    WithConfig(&config.Router).
    // Add custom middleware
    UseMiddleware(errorHandler, rateLimiter).
    // Add simple health check endpoint
    UseHealthCheck().
    // Register Swagger endpoint with pre-generated SwaggerInfo and optional swaggerConfig
    UseSwagger(docs.SwaggerInfo, &config.Swagger)
```
#### Router configuration
```go
type RouterConfig struct {
    UseLogger   bool
    UseRecovery bool
}
```

#### Swagger configuration
```go
type SwaggerConfig struct {
    RoutePath   string
    Version     string
    Host        string
    BasePath    string
    Schemes     []string
    Title       string
    Description string
}
```