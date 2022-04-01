package list

import (
	"fmt"
	"sort"
	"sync"
)

var listMap = struct {
	sync.RWMutex
	l map[int]List
}{l: make(map[int]List)}

func getAllLists()[]List{
	listMap.RLock()
	lists := make([]List, 0, len(listMap.l))
	for _, value := range listMap.l {
		lists = append(lists, value)
	}
	listMap.RUnlock()
	return lists
}

func createList(list List)(int, error){
	nextListId := getNextListID()
	list.ListId = nextListId
	listMap.Lock()
	listMap.l[nextListId] = list
	listMap.Unlock()
	return nextListId, nil
}

func getListByID(listId int)*List{
	listMap.RLock()
	defer listMap.RUnlock()
	if list, ok:= listMap.l[listId]; ok {
		return &list
	}
	return nil
}

func updateListByID(list List)(int, error){
	oldList := getListByID(list.ListId)
	oldListId := list.ListId

	if oldList == nil {
		return 0, fmt.Errorf("list id [%d] doesn't exist", oldList.ListId)
	}
	listMap.Lock()
	listMap.l[oldListId] = list
	listMap.Unlock()
	return oldListId, nil
}

func deleteListByID(listId int){
	listMap.Lock()
	defer listMap.Unlock()
	delete(listMap.l, listId)
}

// func getListTasks() {}

// Utility Functions -------------------------------------------------

func getListIds() []int {
	listMap.RLock()
	listIds := []int{}
	for key := range listMap.l {
		listIds = append(listIds, key)
	}
	listMap.RUnlock()
	sort.Ints(listIds)
	return listIds
}

func getNextListID() int {
	listIds := getListIds()
	fmt.Println(listIds)

	if len(listIds) == 0 {
		return 1
	}

	return listIds[len(listIds)-1] + 1
}