package list

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/austin-rey/go-react-trello-clone/cors"
)

const listPath = "list"

func handleList(w http.ResponseWriter, r *http.Request){	
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", listPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	listId, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
		case http.MethodGet:
			list := getListByID(listId)
			if list == nil{
				w.WriteHeader(http.StatusNotFound)
			}

			l, err := json.Marshal(list)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = w.Write(l)
			if err != nil {
				log.Fatal(err)
			}

		case http.MethodPut:
			var list List
			err := json.NewDecoder(r.Body).Decode(&list)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}
			
			if list.ListId != listId {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = updateListByID(list)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		case http.MethodDelete:
			deleteListByID(listId)

		case http.MethodOptions:
			return
			
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleLists(w http.ResponseWriter, r *http.Request)     {
	switch r.Method {
		case http.MethodGet:
			lists := getAllLists()
			l,err := json.Marshal(lists)
			if err != nil {
				log.Fatal(err)
			}

			_,err = w.Write(l)
			if err != nil {
				log.Fatal(err)
			}
		case http.MethodPost:
			var list List

			// Decode user to see if it matches user struct
			err := json.NewDecoder(r.Body).Decode(&list)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			_, err = createList(list)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
	}
}

func handleListTasks(w http.ResponseWriter, r *http.Request) {}

func SetupRoutes(apiBasePath string) {
	listHandler := http.HandlerFunc(handleList)
	listsHandler := http.HandlerFunc(handleLists)
	listTaskHandler := http.HandlerFunc(handleListTasks)

	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, listPath), cors.Middleware(listHandler))
	http.Handle(fmt.Sprintf("%s/%ss", apiBasePath, listPath), cors.Middleware(listsHandler))
	http.Handle(fmt.Sprintf("%s/%s/tasks", apiBasePath, listPath), cors.Middleware(listTaskHandler))
}