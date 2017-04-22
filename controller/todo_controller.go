package controller

import (
	"crypto/rand"
	"encoding/binary"
	"log"
	"net/http"
	"strconv"

	"gopkg.in/mgo.v2/bson"

	"github.com/labstack/echo"
	"github.com/shogo-ma/todo_app/db"
	"github.com/shogo-ma/todo_app/model"
)

func MakeRandomString() string {
	var n uint64
	binary.Read(rand.Reader, binary.LittleEndian, &n)
	return strconv.FormatUint(n, 36)
}

func GetTodo(c echo.Context) error {
	todo_id := c.Param("id")

	database, err := db.Init("localhost", "sample_db")
	if err != nil {
		return err
	}

	collection := database.C("todo")

	todo := new(model.Todo)
	err = collection.Find(bson.M{
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
func PostTodo(c echo.Context) error {
	database, err := db.Init("localhost", "sample_db")
	if err != nil {
		return err
	}

	collection := database.C("todo")

	data := new(model.Todo)
	if err = c.Bind(data); err != nil {
		return err
	}

	data.TodoID = MakeRandomString()

	err = collection.Insert(data)
	if err != nil {
		log.Println(err)
		return err
	}

	return c.NoContent(http.StatusOK)
}
func DeleteTodo(c echo.Context) error {
	todo_id := c.Param("id")

	database, err := db.Init("localhost", "sample_db")
	if err != nil {
		return err
	}

	collection := database.C("todo")

	err = collection.Remove(bson.M{
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

func CheckedTodo(c echo.Context) error {
	todo_id := c.Param("id")

	database, err := db.Init("localhost", "sample_db")
	if err != nil {
		return err
	}

	collection := database.C("todo")

	todo := new(model.Todo)
	err = collection.Find(bson.M{
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

func GetTodos(c echo.Context) error {
	var todos []model.Todo

	database, err := db.Init("localhost", "sample_db")
	if err != nil {
		return err
	}

	collection := database.C("todo")
	if err := collection.Find(nil).All(&todos); err != nil {
		log.Fatal(err)
		return err
	}

	return c.JSON(http.StatusOK, todos)
}
