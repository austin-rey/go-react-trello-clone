package board

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/austin-rey/go-react-trello-clone/cors"
)

const boardPath = "board"

func handleBoard(w http.ResponseWriter, r *http.Request) {
	urlPathSegments := strings.Split(r.URL.Path, fmt.Sprintf("%s/", boardPath))
	if len(urlPathSegments[1:]) > 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	boardID, err := strconv.Atoi(urlPathSegments[len(urlPathSegments)-1])
	if err != nil {
		log.Print(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	switch r.Method {
	case http.MethodGet:
		organization := getBoard(boardID)
		if organization == nil {
			w.WriteHeader(http.StatusNotFound)
		}

		o, err := json.Marshal(organization)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = w.Write(o)
		if err != nil {
			log.Fatal(err)
		}

	case http.MethodPut:
		var board Board
		err := json.NewDecoder(r.Body).Decode(&board)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		if board.BoardID != boardID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = updateBoard(board)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

	case http.MethodDelete:
		deleteBoard(boardID)

	case http.MethodOptions:
		return

	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func handleBoards(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		boards := getAllBoards()
		o, err := json.Marshal(boards)
		if err != nil {
			log.Fatal(err)
		}

		_, err = w.Write(o)
		if err != nil {
			log.Fatal(err)
		}
	case http.MethodPost:
		var board Board

		// Decode user to see if it matches user struct
		err := json.NewDecoder(r.Body).Decode(&board)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		_, err = createBoard(board)
		if err != nil {
			log.Print(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

func handleBoardLists(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func SetupRoutes(apiBasePath string) {
	boardHandler := http.HandlerFunc(handleBoard)
	boardsHandler := http.HandlerFunc(handleBoards)
	boardListsHandler := http.HandlerFunc(handleBoardLists)

	http.Handle(fmt.Sprintf("%s/%s/", apiBasePath, boardPath), cors.Middleware(boardHandler))
	http.Handle(fmt.Sprintf("%s/%ss", apiBasePath, boardPath), cors.Middleware(boardsHandler))
	http.Handle(fmt.Sprintf("%s/%s/list", apiBasePath, boardPath), cors.Middleware(boardListsHandler))
}