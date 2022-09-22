package main

import (
	"echo/handler"
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	log "github.com/sirupsen/logrus"
)

//테스트

type user struct {
	Name       string      `json:"name"`
	Age        int         `json:"age"`
	Data       string      `json:"data"`
	Background interface{} `json:"background"`
}

const (
	G_API_WEB = "/api/web"
)

func init() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	e := echo.New()
	log.Println("echo start")
	e.Use(middleware.Logger())
	echoGroup := e.Group(G_API_WEB)

	echoGroup.PUT("/dashboard/background", handler.PutBackgrounChange)

	e.GET("/", func(c echo.Context) error {
		fmt.Println("GET /")
		return c.String(http.StatusOK, "Hello, echo!")
	})

	e.POST("/query", handler.Query)

	e.Logger.Fatal(e.Start(":3000"))
}
