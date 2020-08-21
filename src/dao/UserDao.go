package dao

import (
	"appstud.com/github-core/src/models"
	"fmt"
	"sort"
)

var usersStore []models.UserRegistered

func GetUserIfExists(username string) bool {
	usersStoredCount := len(usersStore)
	if usersStoredCount > 0 {
		sort.Slice(usersStore, func(i, j int) bool {
			return usersStore[i].Username <= usersStore[j].Username
		})

		idx := sort.Search(usersStoredCount, func(i int) bool {
			return (usersStore[i].Username) == username
		})

		if idx < usersStoredCount && usersStore[idx].Username == username {
			fmt.Println("Found:", idx, usersStore[idx])
			return true
		} else {
			fmt.Println("Found noting: ", idx)
		}
	}

	return false
}

func AddUser(username string, password string) models.UserRegistered {
	var newUser = models.UserRegistered{
		Username: username,
		Password: password,
	}
	usersStore = append(usersStore, newUser)

	return newUser
}

func GetAllUsers() []models.UserRegistered {
	return usersStore
}
