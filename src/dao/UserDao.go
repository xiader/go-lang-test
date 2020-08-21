package dao

import (
	"appstud.com/github-core/src/models"
	"fmt"
	"sort"
)

var usersStore []models.UserRegistered

func GetUserIfExists(username string) (bool, error) {
	if len(usersStore) > 0 {
		sort.Slice(usersStore, func(i, j int) bool {
			return usersStore[i].Username <= usersStore[j].Username
		})

		idx := sort.Search(len(usersStore), func(i int) bool {
			return string(usersStore[i].Username) >= username
		})

		if usersStore[idx].Username == username {
			fmt.Println("Found:", idx, usersStore[idx])
			return true, nil
		} else {
			fmt.Println("Found noting: ", idx)
		}
	}

	return false, nil
}

func AddUser(username string, password string) models.UserRegistered {
	var newUser = models.UserRegistered{
		Username: username,
		Password: password,
	}
	usersStore = append(usersStore, newUser)

	return newUser
}
