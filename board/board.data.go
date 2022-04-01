package board

import (
	"fmt"
	"sort"
	"sync"
)

var boardMap = struct {
	sync.RWMutex
	b map[int]Board
}{b: make(map[int]Board)}

func getBoard(boardId int) *Board {
	boardMap.RLock()
	defer boardMap.RUnlock()
	if board, ok := boardMap.b[boardId]; ok {
		return &board
	}
	return nil
}

// func getBoardLists(boardId int) {

// }

func updateBoard(board Board) (int, error) {
	oldBoard := getBoard(board.BoardID)
	oldBoardId := board.BoardID

	if oldBoard == nil {
		return 0, fmt.Errorf("board id [%d] doesn't exist", oldBoard.BoardID)
	}
	boardMap.Lock()
	boardMap.b[oldBoardId] = board
	boardMap.Unlock()
	return oldBoardId, nil
}

func deleteBoard(boardId int) {
	boardMap.Lock()
	defer boardMap.Unlock()
	delete(boardMap.b, boardId)
}

func createBoard(board Board) (int, error) {
	nextOrgId := getNextBoardID()
	board.BoardID = nextOrgId
	boardMap.Lock()
	boardMap.b[nextOrgId] = board
	boardMap.Unlock()
	return nextOrgId, nil
}

func getAllBoards() []Board {
	boardMap.RLock()
	boards := make([]Board, 0, len(boardMap.b))
	for _, value := range boardMap.b {
		boards = append(boards, value)
	}
	boardMap.RUnlock()
	return boards
}

// Utility Functions -------------------------------------------------

func getBoardIds() []int {
	boardMap.RLock()
	boardIds := []int{}
	for key := range boardMap.b {
		boardIds = append(boardIds, key)
	}
	boardMap.RUnlock()
	sort.Ints(boardIds)
	return boardIds
}

func getNextBoardID() int {
	boardIds := getBoardIds()
	fmt.Println(boardIds)

	if len(boardIds) == 0 {
		return 1
	}

	return boardIds[len(boardIds)-1] + 1
}