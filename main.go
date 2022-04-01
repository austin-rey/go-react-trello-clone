package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/austin-rey/go-react-trello-clone/list"
	"github.com/austin-rey/go-react-trello-clone/organization"
	"github.com/austin-rey/go-react-trello-clone/task"
	"github.com/austin-rey/go-react-trello-clone/user"
)

const basePath = "/api"

func main() {
	fmt.Println("Trello API")
	user.SetupRoutes(basePath)
	organization.SetupRoutes(basePath)
	task.SetupRoutes(basePath)
	list.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}