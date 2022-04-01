package task

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/austin-rey/go-react-trello-clone/cors"
)

const taskPath = "task"

// /api/task/{id}
func handleTask(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", taskPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	taskID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
		case http.MethodGet:
			task := getTaskById(taskID)
			if task == nil {
				w.WriteHeader(http.StatusBadRequest)
			}
			
			t,err := json.Marshal(task)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = w.Write(t)
			if err != nil {
				log.Fatal(err)
			}
		case http.MethodPut:
			var task Task
			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			
			if task.TaskID != taskID {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = updateTaskById(task)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		case http.MethodDelete:
			deleteTaskByID(taskID)
	}
}

// /api/tasks
func handleTasks(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodGet:
			tasks := getTasks()
			t, err := json.Marshal(tasks)

			if err != nil {
				log.Fatal(err)
			}

			_, err = w.Write(t)
			if err != nil {
				log.Fatal(err)
			}
		case http.MethodPost:
			var task Task

			err := json.NewDecoder(r.Body).Decode(&task)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = createTask(task)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			w.WriteHeader(http.StatusCreated)
	}
}

func SetupRoutes(apiBasePath string) {
	taskHandler := http.HandlerFunc(handleTask)
	tasksHandler := http.HandlerFunc(handleTasks)

	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, taskPath), cors.Middleware(taskHandler))
	http.Handle(fmt.Sprintf("%s/%ss", apiBasePath, taskPath), cors.Middleware(tasksHandler))
}