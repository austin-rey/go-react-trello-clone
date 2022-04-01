package task

import (
	"fmt"
	"sort"
	"sync"
)

var taskMap = struct {
	sync.RWMutex
	t map[int]Task
}{t: make(map[int]Task)}

func getTasks()[]Task {
	taskMap.RLock()
	tasks := make([]Task, 0, len(taskMap.t))
	for _, value := range taskMap.t {
		tasks = append(tasks, value)
	}
	taskMap.RUnlock()
	return tasks
}

func createTask(task Task)(int, error) {
	nextTaskID := getNextTaskID()
	task.TaskID = nextTaskID
	taskMap.Lock()
	taskMap.t[nextTaskID] = task
	taskMap.Unlock()
	return nextTaskID, nil
}

func getTaskById(taskID int) *Task {
	taskMap.RLock()
	defer taskMap.RUnlock()
	if task, ok:= taskMap.t[taskID]; ok {
		return &task
	}
	return nil
}

func updateTaskById(task Task)(int, error) {
	// Lookup old task and assign it and its ID to own vars
	oldTask := getTaskById(task.TaskID)
	oldTaskId := task.TaskID

	if oldTask == nil {
		return 0, fmt.Errorf("org id [%d] doesn't exist", oldTask.TaskID)
	}
	taskMap.Lock()
	taskMap.t[oldTaskId] = task
	taskMap.Unlock()
	return oldTaskId, nil
}

func deleteTaskByID(taskID int){
	taskMap.Lock()
	defer taskMap.Unlock()
	delete(taskMap.t, taskID)
}

// Utility Functions -------------------------------------------------
func getTaskIds() []int {
	taskMap.RLock()
	taskIds := []int{}
	for key := range taskMap.t {
		taskIds = append(taskIds, key)
	}
	taskMap.RUnlock()
	sort.Ints(taskIds)
	return taskIds
}

func getNextTaskID() int {
	taskIds := getTaskIds()
	fmt.Println(taskIds)

	if len(taskIds) == 0 {
		return 1
	}

	return taskIds[len(taskIds)-1] + 1
}