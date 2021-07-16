package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"toDo-golang-crud/handler"
)

func main() {
	fmt.Println("Starting server...")

	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	router := handler.RouteHandler{}

	e = router.SetupRouter(e)

	err := e.Run()

	if err != nil {
		return
	}
}