package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Grisela/todo-cli/todos"
)



func CreateTodo () {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please fill the title")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)

	fmt.Println("Please fill the description")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	todos.AddTodo(title, description)
	fmt.Println("item added")
}