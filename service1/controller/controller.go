package controller

import (
	"io/ioutil"
	"net/http"

	"github.com/labstack/echo"
)

// Index return html text
func Index(c echo.Context) error {
	return c.HTML(http.StatusOK, "<h1>Hello world form app1</h1>")
}

// GetApp2 get data from app2
func GetApp2(c echo.Context) error {
	res, err := http.Get("http://localhost:3000")
	if err != nil {
		return c.String(500, err.Error())
	}

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return c.String(500, err.Error())
	}

	// To string
	resStr := string(data)

	return c.String(http.StatusOK, resStr)
}
