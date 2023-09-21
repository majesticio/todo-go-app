package main

import (
	"embed"
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

// Embed the index.html file
//
//go:embed index.html
var content embed.FS

type Task struct {
	ID   int
	Name string
}

func main() {
	http.HandleFunc("/", displayTasks)
	http.HandleFunc("/add-task/", addTask)
	http.HandleFunc("/remove-task/", removeTask)
	http.HandleFunc("/clear-tasks/", clearTasks)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func displayTasks(w http.ResponseWriter, r *http.Request) {
	data, err := content.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.New("index.html").Parse(string(data)))
	tmpl.Execute(w, map[string]interface{}{"Tasks": tasks})
}

func addTask(w http.ResponseWriter, r *http.Request) {
	taskLock.Lock()
	defer taskLock.Unlock()

	name := r.PostFormValue("name")
	taskID++
	tasks = append(tasks, Task{ID: taskID, Name: name})
	data, err := content.ReadFile("index.html")
	if err != nil {
		http.Error(w, "Failed to load template", http.StatusInternalServerError)
		return
	}
	tmpl := template.Must(template.New("index.html").Parse(string(data)))
	tmpl.ExecuteTemplate(w, "todo-list-element", Task{ID: taskID, Name: name})
}

func removeTask(w http.ResponseWriter, r *http.Request) {
	taskLock.Lock()
	defer taskLock.Unlock()

	id := r.PostFormValue("id")

	for index, task := range tasks {
		if fmt.Sprintf("%d", task.ID) == id {
			tasks = append(tasks[:index], tasks[index+1:]...)
			break
		}
	}

	w.WriteHeader(http.StatusOK)
}

func clearTasks(w http.ResponseWriter, r *http.Request) {
	taskLock.Lock()
	defer taskLock.Unlock()

	tasks = []Task{} // Clear the tasks slice

	w.WriteHeader(http.StatusOK)
}
