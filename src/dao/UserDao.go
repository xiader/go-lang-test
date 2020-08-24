package dao

import (
	"appstud.com/github-core/src/models"
	"errors"
	"fmt"
	"sort"
)

var usersStore []models.UserRegistered
var usersLoggedIn = make(map[string]models.UserRegistered)

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

func AuthenticateUser(username string, password string) (models.UserRegistered, error) {
	var existingUser = getUserByUsername(username)
	//check on existence
	if (models.UserRegistered{}) == existingUser {
		return existingUser, errors.New("user is not registered")
	}
	if existingUser.Password != password {
		return existingUser, errors.New("password is invalid")
	}

	return existingUser, nil
}

func getUserByUsername(username string) models.UserRegistered {
	var empty models.UserRegistered
	index := sort.Search(len(usersStore), func(i int) bool {
		return (usersStore[i].Username) == username
	})
	if index < len(usersStore) && usersStore[index].Username == username {
		empty = usersStore[index]
	}

	return empty
}

func AddToken2User(token string, registered models.UserRegistered) {
	usersLoggedIn[token] = registered
}

func GetActiveUsersByToken(token string) (models.UserRegistered, error) {
	userByToken, ok := usersLoggedIn[token]
	if ok == false {
		return userByToken, errors.New("no users active by token " + token)
	}

	return userByToken, nil
}
