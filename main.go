package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shogo-ma/todo_app/controller"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public/views")
	e.Static("/static", "assets")

	e.POST("/api/v1/todo", controller.PostTodo)
	e.GET("/api/v1/todos", controller.GetTodos)

	e.GET("/api/v1/todo/:id", controller.GetTodo)
	e.DELETE("/api/v1/todo/:id", controller.DeleteTodo)

	e.PUT("/api/v1/checked/:id", controller.CheckedTodo)
	e.Logger.Fatal(e.Start(":8080"))
}
