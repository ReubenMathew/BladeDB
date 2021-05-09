package main

import (
	"fmt"
	"net/http"

	"github.com/ReubenMathew/BladeDB/model"
	"github.com/ReubenMathew/BladeDB/raft"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	fmt.Println("Hello World")
	r := raft.NewRaft(2)
	fmt.Println(r.GetID())
	s := model.NewServer(1)
	fmt.Println(s)
	// echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	/* Routing */
	e.GET("/", home)
	e.GET("/get", get)

	e.Logger.Fatal(e.Start(":1323"))
}

func home(c echo.Context) error {
	return c.String(http.StatusOK, "This is BladeDB!")
}

// key-value pair
type KVP struct {
	Key   string `json:"key" xml:"key"`
	Value string `json:"value" xml:"value"`
}

func get(c echo.Context) error {

	key := c.QueryParam("key")
	kvp := &KVP{
		Key:   key,
		Value: "This is a value that maps to the key",
	}
	return c.JSON(http.StatusOK, kvp)
}
