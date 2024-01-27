package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
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
	router.POST("/newTodo", addTodos)
	router.GET("/todos/:id", getTodo)
	router.PATCH("/todos/:id", toggleTodoStatus)
	router.Run("localhost:8081")

}

func getTodos(context *gin.Context) {
	//to json
	context.IndentedJSON(http.StatusOK, todos)
}

func addTodos(context *gin.Context) {
	var newTodo todo

	if err := context.BindJSON(&newTodo); err != nil {
		return
	}

	todos = append(todos, newTodo)

	context.IndentedJSON(http.StatusCreated, newTodo)
}

func toggleTodoStatus(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}

	todo.Completed = !todo.Completed

	ctx.IndentedJSON(http.StatusOK, todo)
}

func getTodo(ctx *gin.Context) {
	id := ctx.Param("id")

	todo, err := getTodoById(id)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Todo not found"})
		return
	}
	ctx.IndentedJSON(http.StatusOK, todo)
}

func getTodoById(id string) (*todo, error) {
	for i, t := range todos {
		if t.ID == id {
			return &todos[i], nil
		}
	}
	return nil, errors.New("Todo not found")
}
