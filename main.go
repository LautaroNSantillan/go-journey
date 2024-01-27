package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json: "title"`
	Completed bool   `json: "completed"`
}

var todos = []todo{
	{ID: "1", Item: "todo1", Completed: false},
	{ID: "2", Item: "todo2", Completed: true},
	{ID: "3", Item: "todo3", Completed: false},
}

func main() {
	//Create n run server
	router := gin.Default()
	router.GET("/todos", getTodos)
	router.Run("localhost:8081")

}

func getTodos(context *gin.Context) {
	//to json
	context.IndentedJSON(http.StatusOK, todos)
}
