package board

type Board struct {
	BoardID          int    `json:"boardId"`
	BoardName        string `json:"boardName"`
	BoardDescription string `json:"boardDescription"`
}