package main

import (
	"Interview/handlers"

	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	e.GET("/needcab", handlers.Cabdetails)
	e.POST("/bookcab", handlers.Getacab)
	e.GET("/mybooking", handlers.Booking)
	e.Logger.Fatal(e.Start(":8008"))

}
