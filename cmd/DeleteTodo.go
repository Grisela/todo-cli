package cmd

import (
	"fmt"

	"github.com/Grisela/todo-cli/todos"
)

func DeleteTodo () {
	var id string
	fmt.Println("Enter the item id to be deleted:")
	fmt.Scan(&id)
	todos.DeleteTodo(id)
	fmt.Println("Delete Success")
}