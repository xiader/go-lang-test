package service

import (
	"appstud.com/github-core/src/dao"
	"appstud.com/github-core/src/models"
	"appstud.com/github-core/src/util"
	"errors"
)

func RegisterUser(username string, password string) (models.UsernameTokenResponse, error) {
	var user models.UsernameTokenResponse
	if username == "" || password == "" {
		return user, errors.New("please check your username and password")
	}
	var userExists = dao.GetUserIfExists(username)
	if userExists {
		return user, errors.New("user with this username already exists")
	}
	var userToken = util.GenerateRandomString()
	dao.AddUser(username, password)

	return models.UsernameTokenResponse{
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

func LoginUser(username string, password string) (models.UsernameTokenResponse, error) {
	var user models.UsernameTokenResponse
	if username == "" || password == "" {
		return user, errors.New("please check your username and password")
	}
	var authenticatedUser, err = dao.AuthenticateUser(username, password)
	if err != nil {
		return user, err
	}
	var userToken = util.GenerateRandomString()
	dao.AddToken2User(userToken, authenticatedUser)
	user = models.UsernameTokenResponse{
		Username: authenticatedUser.Username,
		Token:    userToken,
	}

	return user, nil
}

func GetConnectedUser(token string) (models.UsernameResponse, error) {
	var user models.UsernameResponse
	if token == "" {
		return user, errors.New("no token found in a request")
	}
	var loggedInUser, err = dao.GetActiveUsersByToken(token)
	if err != nil {
		return user, err
	}
	user = models.UsernameResponse{Username: loggedInUser.Username}

	return user, nil
}
