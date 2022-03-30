package main

import (
	"fmt"
	"log"
	"net/http"

	user "github.com/austin-rey/go-react-trello-clone/user"
)

const basePath = "/api"

func main() {
	fmt.Println("Trello API")
	user.SetupRoutes(basePath)
	log.Fatal(http.ListenAndServe(":5000", nil))
}