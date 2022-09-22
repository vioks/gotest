package main

import (
	"os"

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
	e.Logger.Fatal(e.Start(":3000"))
}
