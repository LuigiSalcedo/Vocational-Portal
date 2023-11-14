package main

import (
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/", "static")

	log.Fatal(e.Start(":8088")) // Initializing server
}
