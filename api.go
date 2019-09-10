package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func testApi() {
	e := echo.New()
	e.GET("/hello", helloWorld)
	e.GET("/users/:name", getUserName)
	e.GET("/show", show)
	e.POST("/save", save)
	e.POST("/users", saveUser)
	e.POST("/send", sendMessage)
	e.Logger.Fatal(e.Start(":1323"))
}

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func getUserName(c echo.Context) error {
	name := c.Param("name")
	return c.String(http.StatusOK, name)
}

func show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

func save(c echo.Context) error {
	// Get name and email
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusOK, u)
}

type Message struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

type Response struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:message`
	Stusts  string `json:status`
}

func sendMessage(c echo.Context) error {
	m := new(Message)
	if error := c.Bind(m); error != nil {
		return error
	}
	r := new(Response)
	r.Name = m.Name
	r.Email = m.Email
	r.Message = m.Message
	r.Stusts = "success"
	return c.JSON(http.StatusOK, r)
}
