package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"toDo-golang-crud/model"
	"toDo-golang-crud/repository"
)

const (
	EmptyPath = ""
	http500 = "INTERNAL_ERROR"
	http400 = "BAD_REQUEST"
	http404 = "NOT_FOUND"
)

type RouteHandler struct {
	Repository *repository.DBConnection
}

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
		c.JSON(http.StatusBadRequest, HandleError(http400, err.Error()))
		return
	}

	saved, err := r.Repository.Save(&body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, HandleError(http500, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, saved)
}

func (r *RouteHandler) GetOne(c *gin.Context) {
	id := c.Param("id")
	current, err, is404 := r.Repository.GetById(id)

	if err != nil {

		var httpStatus = http.StatusInternalServerError
		var code = http500
		if is404 {
			httpStatus = http.StatusNotFound
			code = http404
		}

		c.JSON(httpStatus, HandleError(code, err.Error()))
		return
	}

	c.JSON(http.StatusOK, current)
}

func (r *RouteHandler) GetAll(c *gin.Context) {

	todos, err, is404 := r.Repository.GetALl()

	if err != nil {

		var httpStatus = http.StatusInternalServerError
		var code = http500
		if is404 {
			httpStatus = http.StatusNotFound
			code = http404
		}

		c.JSON(httpStatus, HandleError(code, err.Error()))
		return
	}

	c.JSON(http.StatusOK, todos)
}

func (r *RouteHandler) Put(c *gin.Context) {
	var id = c.Param("id")
	var body model.ToDo

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, HandleError(http400, err.Error()))
		return
	}

	current, err, is404 := r.Repository.GetById(id)

	if err != nil {

		var httpStatus = http.StatusInternalServerError
		var code = http500
		if is404 {
			httpStatus = http.StatusNotFound
			code = http404
		}

		c.JSON(httpStatus, HandleError(code, err.Error()))
		return
	}

	body.Id = current.Id
	updated, err := r.Repository.Put(&body)

	if err != nil {
		c.JSON(http.StatusInternalServerError, HandleError(http500, err.Error()))
	}

	c.JSON(http.StatusOK, updated)
}

func (r *RouteHandler) Delete(c *gin.Context) {
	err := r.Repository.Delete(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, HandleError(http500, err.Error()))
		return
	}

	c.Status(http.StatusNoContent)
}


func HandleError(code string, msg string) map[string]interface{} {
	return map[string]interface{}{"code": code, "message": msg}
}