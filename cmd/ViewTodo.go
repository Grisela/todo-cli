package cmd

import (
	"fmt"
	"strings"

	"github.com/Grisela/todo-cli/todos"
)

func ViewTodo () {
	maxIDLen, maxTitleLen, maxDescLen := len("ID"), len("Title"), len("Description")
	for _, todo := range todos.TodoList {
		if len(todo.Id) > maxIDLen {
			maxIDLen = len(todo.Id)
		}
		if len(todo.Title) > maxTitleLen {
			maxTitleLen = len(todo.Title)
		}
		if len(todo.Description) > maxDescLen {
			maxDescLen = len(todo.Description)
		}
	}

	headerFormat := fmt.Sprintf("%%-%ds | %%-%ds | %%-%ds\n", maxIDLen, maxTitleLen, maxDescLen)
	rowFormat := fmt.Sprintf("%%-%ds | %%-%ds | %%-%ds\n", maxIDLen, maxTitleLen, maxDescLen)

	fmt.Printf(headerFormat, "ID", "Title", "Description")
	fmt.Println(strings.Repeat("-", maxIDLen+maxTitleLen+maxDescLen+6))


	for _, p := range todos.TodoList {
		fmt.Printf(rowFormat, p.Id, p.Title, p.Description)
	}
}