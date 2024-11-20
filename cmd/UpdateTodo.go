package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Grisela/todo-cli/todos"
)

func UpdateTodo () {
	var id string
	var selectedMenu int16
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Which one do you want to update?\n1. Title\n2. Description\n3. Both\n4. cancel")
		fmt.Scan(&selectedMenu)
		if selectedMenu != 4 {
		fmt.Println("Type the id")
		fmt.Scan(&id)
		}
		switch selectedMenu {
		case 1:
			fmt.Println("Please fill the title")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)
			var description string
			todos.EditTodo(id, &title, &description)
			fmt.Println("item edited")
		case 2:
			fmt.Println("Please fill the description")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			var title string
			todos.EditTodo(id, &title, &description)
			fmt.Println("item edited")
		case 3:
			fmt.Println("Please fill the title")
			title, _ := reader.ReadString('\n')
			title = strings.TrimSpace(title)

			fmt.Println("Please fill the description")
			description, _ := reader.ReadString('\n')
			description = strings.TrimSpace(description)
			todos.EditTodo(id, &title, &description)
		case 4:
			fmt.Println("cancelling...")
			return
		default:
			fmt.Println("Invalid Selection!")
		}
	}
}