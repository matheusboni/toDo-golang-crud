package main

import (
	"fmt"
	goEnv "github.com/Netflix/go-env"
	"github.com/gin-gonic/gin"
	"log"
	"toDo-golang-crud/config"
	"toDo-golang-crud/handler"
	"toDo-golang-crud/repository"
)

var env config.Env

func init()  {
	if _, err := goEnv.UnmarshalFromEnviron(&env); err != nil {
		log.Fatalf("Error decoding environment variables, %v", err)
	}
}

func main() {
	fmt.Println("Starting server...")

	e := gin.New()
	e.Use(gin.Logger())
	e.Use(gin.Recovery())

	config.SetUpDB(&env)
	router := handler.RouteHandler{
		Repository: &repository.DBConnection{DB: config.GetDB()},
	}

	e = router.SetupRouter(e)

	var port = env.ApplicationPort
	var err error

	if port == "" {
		err = e.Run()
	} else {
		err = e.Run(port)
	}

	if err != nil {
		fmt.Println("Error while starting server...")
	}
}