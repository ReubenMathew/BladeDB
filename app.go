package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	// echo instance
	e := echo.New()

	e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	/* Routing */
	e.GET("/", home)
	e.GET("/get", get)
	e.POST("/put", post)

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
		Key: key,
	}
	return c.JSON(http.StatusOK, kvp)
}

func post(c echo.Context) error {
	key := c.FormValue("key")
	value := c.FormValue("value")

	insert(key, value) //placeholder

	returnString := "key: " + key + " value: " + value
	return c.String(http.StatusAccepted, returnString)
}

func insert(key string, value string) {
	// replace with cluster insert
}
