package router

import (
	"github.com/khivuksergey/webserver"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"github.com/swaggo/swag"
)

const (
	swaggerName = "webserver-swagger"
	swaggerPath = "/swagger/*"
)

type Router struct {
	*echo.Echo
}

func NewEchoRouter(cfg *webserver.RouterConfig) *Router {
	r := &Router{echo.New()}

	if cfg != nil {
		if cfg.UseLogger {
			r.Use(echoMiddleware.Logger())
		}
		if cfg.UseRecovery {
			r.Use(echoMiddleware.Recover())
		}
	}

	return r
}

func (r *Router) UseMiddleware(middleware ...echo.MiddlewareFunc) *Router {
	r.Use(middleware...)
	return r
}

func (r *Router) UseHealthCheck() *Router {
	r.GET("/health", health)
	return r
}

func (r *Router) UseSwagger(swaggerInfo *swag.Spec, swaggerConfig *webserver.SwaggerConfig) *Router {
	if swaggerInfo == nil {
		panic("cannot use nil swagger")
	}
	path := configureSwagger(swaggerInfo, swaggerConfig)
	r.GET(path, echoSwagger.WrapHandler)
	return r
}

func configureSwagger(info *swag.Spec, config *webserver.SwaggerConfig) string {
	defer swag.Register(swaggerName, info)

	path := swaggerPath

	if config == nil {
		return path
	}

	setIfNotEmpty(&path, &config.RoutePath)
	setIfNotEmpty(&info.Version, &config.Version)
	setIfNotEmpty(&info.Host, &config.Host)
	setIfNotEmpty(&info.BasePath, &config.BasePath)
	setIfNotEmpty(&info.Title, &config.Title)
	setIfNotEmpty(&info.Description, &config.Description)

	if len(config.Schemes) > 0 {
		info.Schemes = config.Schemes
	}

	return path
}

func setIfNotEmpty(destination, source *string) {
	if source != nil && len(*source) > 0 {
		*destination = *source
	}
}
