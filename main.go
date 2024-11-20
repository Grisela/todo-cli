package main

import (
	"fmt"

	"github.com/Grisela/todo-cli/cmd"
)

func main() {

	var selectMenu int16

	for {
		fmt.Println("Type your desired menu:")
	fmt.Printf("1. Add Todo\n2. Edit Todo\n3. Delete Todo\n4. View Todo\n5. Exit\n")
	fmt.Scan(&selectMenu)

	switch selectMenu {
	case 1:
		cmd.CreateTodo()
	case 2:
		cmd.UpdateTodo()
	case 3:
		cmd.DeleteTodo()
	case 4:
		cmd.ViewTodo()
	case 5:
		fmt.Println("Exiting...")
		return
	default:
		fmt.Println("Invalid Selection!")
	}
	}

}