package user

import (
	"fmt"
	"sort"
	"sync"
)

var userMap = struct {
	sync.RWMutex
	u map[int]User
}{u : make(map[int]User)}

func init() {

}

func getUser(userId int) (*User) {
	userMap.RLock();
	defer userMap.RUnlock()
	if user, ok := userMap.u[userId]; ok {
		return &user
	}
	return nil
}

func addUser(user User) (int, error) {
	nextId := getNextUserID()
	user.UserID = nextId
	userMap.Lock()
	userMap.u[nextId] = user
	userMap.Unlock()
	return nextId, nil
}

func updateUser(user User) (int, error) {
	oldUser := getUser(user.UserID)
	oldUserId := user.UserID

	if oldUser == nil {
		return 0, fmt.Errorf("product id [%d] doesn't exist", user.UserID)
	}

	userMap.Lock()
	userMap.u[oldUserId] = user
	userMap.Unlock()
	return oldUserId, nil
}

func removeUser(userId int) {
	userMap.Lock()
	defer userMap.Unlock()
	delete(userMap.u, userId)
}

func getUsers() []User {
	userMap.RLock()
	users := make([]User,0,len(userMap.u))
	for _, value := range userMap.u {
		users = append(users, value)
	}
	userMap.RUnlock()
	return users
}

// Utility Functions -------------------------------------------------

func getUserIds() []int {
	userMap.RLock()
	userIds := []int{}
	for key := range userMap.u {
		userIds = append(userIds, key)
	}
	userMap.RUnlock()
	sort.Ints(userIds)
	return userIds
}

func getNextUserID() int {
	userIds := getUserIds()
	fmt.Println(userIds)

	if len(userIds) == 0 {
		return 1
	}

	return userIds[len(userIds)-1] + 1
}