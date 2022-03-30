package user

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/austin-rey/go-react-trello-clone/cors"
)

const userPath = "user"

// /api/user/{id}
func handleUser(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", userPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	userID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
		case http.MethodGet:
			// Get user object by ID
			user := getUser(userID)
			if user == nil {
				w.WriteHeader(http.StatusNotFound)
			}

			// Format object into JSON
			u, err := json.Marshal(user)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Add user json to response
			_,err = w.Write(u)
			if err != nil {
				log.Fatal(err)
			}

		case http.MethodPut:
			// Get user json to be updated
			var user User
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// See if url param for users id matched that of the request body
			if user.UserID != userID {
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Update user fields
			_, err = updateUser(user)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

		case http.MethodDelete:
			removeUser(userID)

		case http.MethodOptions:
			return
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

// /api/users
func handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method{
		case http.MethodGet:
			// Get list of users and marshall into json
			userList := getUsers()
			u, err := json.Marshal(userList)
			if err != nil {
				log.Fatal(err)
			}

			// Add json data to respose
			_, err = w.Write(u)
			if err != nil {
				log.Fatal(err)
			}

		case http.MethodPost:
			var user User
		
			// Decode user to see if it matches user struct
			err := json.NewDecoder(r.Body).Decode(&user)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			// Add user to user map
			_, err = addUser(user)
			if err != nil {
				log.Print(err)
				w.WriteHeader(http.StatusBadRequest)
				return
			}

			w.WriteHeader(http.StatusCreated)
		}
}

func SetupRoutes(apiBasePath string) {
	userHandler := http.HandlerFunc(handleUser)
	usersHandler := http.HandlerFunc(handleUsers)
	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, userPath), cors.Middleware(userHandler))
	http.Handle(fmt.Sprintf("%s/%ss", apiBasePath, userPath), cors.Middleware(usersHandler))
}