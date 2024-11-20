package todos

import (
	"fmt"
	"slices"

	"github.com/google/uuid"
)

type Todo struct {
	Id			string
	Title		string
	Description	string
}

var TodoList []Todo

func AddTodo(title, description string) {
	TodoList = append(TodoList, Todo{
		Id:          uuid.New().String(),
		Title:       title,
		Description: description,
	})
}

func EditTodo(id string, title, description *string){
	idx := slices.IndexFunc(TodoList, func(c Todo) bool {return c.Id == id})
	if idx == -1 {
        return
    }
	if title != nil {
		TodoList[idx].Title = *title
	}
	if description != nil {
		TodoList[idx].Description = *description
	}
}

func DeleteTodo(id string){
	idx := slices.IndexFunc(TodoList, func(c Todo) bool { return c.Id == id })
	if idx == -1 {
        return
    }
	TodoList = slices.Delete(TodoList, idx, idx+1)
	fmt.Println(TodoList)
}