package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sync"
)

var (
	tasks    = []Task{}
	taskID   = 0
	taskLock sync.Mutex
)

type Task struct {
	ID   int
	Name string
}

func main() {
	http.HandleFunc("/", displayTasks)
	http.HandleFunc("/add-task/", addTask)
	http.HandleFunc("/remove-task/", removeTask)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func displayTasks(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.Execute(w, map[string]interface{}{"Tasks": tasks})
}

func addTask(w http.ResponseWriter, r *http.Request) {
	taskLock.Lock()
	defer taskLock.Unlock()

	name := r.PostFormValue("name")
	taskID++
	tasks = append(tasks, Task{ID: taskID, Name: name})

	tmpl := template.Must(template.ParseFiles("index.html"))
	tmpl.ExecuteTemplate(w, "todo-list-element", Task{ID: taskID, Name: name})
}

func removeTask(w http.ResponseWriter, r *http.Request) {
	taskLock.Lock()
	defer taskLock.Unlock()

	id := r.PostFormValue("id")
	// fmt.Println("Received ID:", id)

	for index, task := range tasks {
		if fmt.Sprintf("%d", task.ID) == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}

	// fmt.Println("Current tasks after removal:", tasks)

	w.WriteHeader(http.StatusOK)
}
