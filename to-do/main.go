package main

import (
	"fmt"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

type TodoList struct {
	Tasks []Task
}

func (t *TodoList) AddTask(title string) {
	task := Task{
		ID:    len(t.Tasks) + 1,
		Title: title,
		Done:  false,
	}
	t.Tasks = append(t.Tasks, task)
}

func (t *TodoList) MarkTaskCompleted(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Done = true
			break
		}
	}
}

func (t *TodoList) DeleteTask(id int) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
		}
	}
}

func (t *TodoList) EditTask(id int, title string) {
	for i, task := range t.Tasks {
		if task.ID == id {
			t.Tasks[i].Title = title
			break
		}
	}
}

func main() {
	list := TodoList{}
	list.AddTask("Estudar Go!")
	list.AddTask("Estudar cada vez mais!!")

	for _, task := range list.Tasks {
		fmt.Printf("Tarefa %d: %s (Concluída: %t)\n", task.ID, task.Title, task.Done)
	}

	list.MarkTaskCompleted(1)
	list.MarkTaskCompleted(2)

	for _, task := range list.Tasks {
		fmt.Printf("Tarefa %d: %s (Concluída: %t)\n", task.ID, task.Title, task.Done)
	}

	list.DeleteTask(1)

	for _, task := range list.Tasks {
		fmt.Printf("Tarefa %d: %s (Concluída: %t)\n", task.ID, task.Title, task.Done)
	}

	list.EditTask(2, "Titulo novo vai aqui!!")

	for _, task := range list.Tasks {
		fmt.Printf("Tarefa %d: %s (Concluída: %t)\n", task.ID, task.Title, task.Done)
	}
}
