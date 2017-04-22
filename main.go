package main

import (
	"log"
	"net/http"

	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/shogo-ma/todo_app/model"
)

var collection *mgo.Collection

func getTodo(c echo.Context) error {
	todo_id := c.Param("id")

	todo := new(model.Todo)
	err := collection.Find(bson.M{
		"todoid": bson.M{
			"$eq": todo_id,
		},
	}).One(&todo)

	if err != nil {
		return err
	}

	return c.JSON(http.StatusCreated, todo)
}

// DBにpostする
func postTodo(c echo.Context) error {
	data := new(model.Todo)
	if err := c.Bind(data); err != nil {
		return err
	}

	data.TodoID = MakeRandomString()

	err := collection.Insert(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.NoContent(http.StatusOK)
}

func deleteTodo(c echo.Context) error {
	todo_id := c.Param("id")

	err := collection.Remove(bson.M{
		"todoid": bson.M{
			"$eq": todo_id,
		},
	})

	if err != nil {
		log.Fatal(err)
		return err
	}

	return c.NoContent(http.StatusOK)
}

func checkedTodo(c echo.Context) error {
	todo_id := c.Param("id")

	todo := new(model.Todo)
	err := collection.Find(bson.M{
		"todoid": bson.M{
			"$eq": todo_id,
		},
	}).One(&todo)

	if err != nil {
		return err
	}

	// update query
	query := bson.M{
		"todoid": bson.M{
			"$eq": todo_id,
		},
	}

	change := bson.M{
		"$set": bson.M{
			"todoid": todo_id,
			"text":   todo.Text,
			"status": true,
		}}

	err = collection.Update(query, change)
	if err != nil {
		return err
	}

	return c.NoContent(http.StatusOK)
}

func getTodos(c echo.Context) error {
	var todos []model.Todo
	if err := collection.Find(nil).All(&todos); err != nil {
		log.Fatal(err)
		return err
	}

	return c.JSON(http.StatusOK, todos)
}

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		log.Println("dbとの接続に失敗")
		return
	}

	defer session.Close()

	collection = session.DB("sample_db").C("todo")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.Static("/", "public/views")
	e.Static("/static", "assets")

	e.POST("/api/v1/todo", postTodo)
	e.GET("/api/v1/todos", getTodos)

	e.GET("/api/v1/todo/:id", getTodo)
	e.DELETE("/api/v1/todo/:id", deleteTodo)

	e.PUT("/api/v1/checked/:id", checkedTodo)
	e.Logger.Fatal(e.Start(":8080"))
}
