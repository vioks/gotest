package main

import (
	"echo/handler"
	"fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"

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
	e.Use(middleware.Logger())
	echoGroup := e.Group(G_API_WEB)

	echoGroup.PUT("/dashboard/background", handler.PutBackgrounChange)

	e.GET("/", func(c echo.Context) error {
		fmt.Println("GET /")
		return c.String(http.StatusOK, "Hello, echo!")
	})

	e.POST("/query", handler.Query)

	e.GET("/path/:name", func(c echo.Context) error {
		name := c.Param("name")
		return c.String(200, name)
	})

	e.POST("/form", func(c echo.Context) error {
		name := c.FormValue("name")
		age, _ := strconv.Atoi(c.FormValue("age"))
		return c.JSON(http.StatusOK, user{Name: name, Age: age})
	})

	//Bind()도 구조체 필드들이 대문자여야 가능
	e.POST("/json", func(c echo.Context) error {
		user := user{}
		if err := c.Bind(&user); err != nil {
			return c.String(http.StatusBadRequest, err.Error())
		}
		if user.Data == "" {
			fmt.Println("Data 공백")
			user.Data = "board"
		}

		fmt.Println(reflect.TypeOf(user.Background))

		// user.Background = string(bd)

		return c.JSON(http.StatusOK, user)
	})

	e.GET("/json/data", func(c echo.Context) error {
		user := &user{
			Name: "won",
			Age:  27,
			Data: "json 데이터 전송 테스트",
		}

		return c.JSON(http.StatusOK, user)
	})

	e.GET("/json/pretty", func(c echo.Context) error {
		user := &user{
			Name: "won",
			Age:  27,
			Data: "json 데이터 전송 테스트",
		}

		return c.JSONPretty(http.StatusOK, user, "	")
	})

	e.Logger.Fatal(e.Start(":3000"))
}
