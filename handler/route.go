package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"toDo-golang-crud/model"
	"toDo-golang-crud/repository"
)

const EmptyPath = ""

type RouteHandler struct {}

func (r *RouteHandler) SetupRouter(e *gin.Engine) *gin.Engine {
	root := e.Group("/todo-api/todos")

	root.GET(EmptyPath, r.GetAll)
	root.POST(EmptyPath, r.Create)
	root.GET("/:id", r.GetOne)
	root.PUT("/:id", r.Put)
	root.DELETE("/:id", r.Delete)

	return e
}

func (r *RouteHandler) Create(c *gin.Context) {
	var body model.ToDo
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, repository.Save(&body))
}

func (r *RouteHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	current := repository.Get(id)

	if current.Id == "" {
		c.JSON(http.StatusNotFound, _HandleNotFound(id))
		return
	}

	c.JSON(http.StatusOK, current)
}

func (r *RouteHandler) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, repository.GetALl())
}

func (r *RouteHandler) Put(c *gin.Context) {
	var id = c.Param("id")
	var body model.ToDo

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	current := repository.Get(id)

	if current.Id != "" {
		body.Id = current.Id
	} else {
		c.JSON(http.StatusNotFound, _HandleNotFound(id))
		return
	}

	c.JSON(http.StatusCreated, repository.Put(&body))
}

func (r *RouteHandler) Delete(c *gin.Context) {
	repository.Delete(c.Param("id"))
	c.Status(http.StatusNoContent)
}

func _HandleNotFound(id string) map[string]interface{} {
	return map[string]interface{}{"code": "NOT_FOUND", "message": fmt.Sprintf("ToDo with id: %s was not found", id)}
}