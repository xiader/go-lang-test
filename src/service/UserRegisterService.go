package service

import (
	"appstud.com/github-core/src/dao"
	"appstud.com/github-core/src/models"
	"appstud.com/github-core/src/util"
	"errors"
)

func RegisterUser(username string, password string) (models.SuccessfulRegistrationResponse, error) {
	var user models.SuccessfulRegistrationResponse
	if username == "" || password == "" {
		return user, errors.New("please check your username and password")
	}
	var userExists = dao.GetUserIfExists(username)
	if userExists {
		return user, errors.New("user with this username already exists")
	}
	var userToken = util.GenerateRandomString()
	dao.AddUser(username, password)

	return models.SuccessfulRegistrationResponse{
		Username: username,
		Token:    userToken,
	}, nil
}

func GetAllUsers() []models.UsernameResponse {
	var storedUsers = dao.GetAllUsers()
	var allUsernames = make([]models.UsernameResponse, len(storedUsers))
	for i := range storedUsers {
		allUsernames[i].Username = storedUsers[i].Username
	}

	return allUsernames
}
