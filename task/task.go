package task

type Task struct {
	TaskID         int    `json:"taskID"`
	TaskName       string `json:"taskName"`
	TaskDesc       string `json:"taskDesc"`
	CreatedAtDate  string `json:"createdAtDate"`
	CompleteByDate string `json:"completeByDate"`
	BoardID        int    `json:"boardID"`
	ListID         int    `json:"listID"`
}

// Future improvements will include comments by users