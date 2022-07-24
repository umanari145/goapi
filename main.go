package main

import (
	"net/http"

	"github.com/joho/godotenv"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		logrus.Fatal("Error loading .env")
	}

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":8080"))
}
