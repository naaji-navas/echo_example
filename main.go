package main

import (

	"github.com/labstack/echo"
	"github.com/naaji-navas/echo_example/cmd/api/handlers"
)


func main() {
	e := echo.New()
	e.GET("/health-check",handlers.HealthCheckHandler)
	e.GET("posts",handlers.PostIndexHandler)
	e.GET("posts/:id",handlers.PostSingleHandler)

	e.Logger.Fatal(e.Start(":1323"))

}
