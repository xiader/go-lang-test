package service

import (
	"appstud.com/github-core/src/models/github/model"
	"appstud.com/github-core/src/models/github/response"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

const githubGetAllEvents string = "https://api.github.com/events"
const githubGetAllEventsByUser string = "https://api.github.com/users/%s/events"
const gitHubGetUserInfo string = "https://api.github.com/users/"

var client = &http.Client{Timeout: time.Minute}

func GetGitHubPublicUserInfo(username string) (interface{}, error) {
	var gitHubUserResponse response.User
	if username == "" {
		return gitHubUserResponse, errors.New("please check username in path params")
	}
	var userInfoRequest = gitHubGetUserInfo + username

	var publicGitHubUserUser, userError = getGitHubPublicUser(userInfoRequest)
	if userError != nil {
		return nil, userError
	}
	user := publicGitHubUserUser.(model.User)
	var publicRepos, reposError = getGitHubPublicUserRepos(user.ReposURL)
	var publicGists, gistsError = getGitHubPublicUserGists(user.GistsURL)
	var followers, followersError = getGitHubUserFollowers(user.FollowersURL)
	var following, followingError = getGitHubUserFollowing(user.FollowingURL)

	if reposError != nil {
		return nil, reposError
	}
	if gistsError != nil {
		return nil, gistsError
	}
	if followersError != nil {
		return nil, followersError
	}
	if followingError != nil {
		return nil, followingError
	}

	//repositories := publicRepos.([]response.PublicRepo)
	//gists := publicGists.([]response.PublicGist)
	//userFollowers := followers.([]response.BriefUserInfo)
	//userFollowing := following.([]response.BriefUserInfo)

	var details = response.Details{
		PublicRepos: publicRepos.([]response.PublicRepo),
		PublicGists: publicGists.([]response.PublicGist),
		Followers:   followers.([]response.BriefUserInfo),
		Following:   following.([]response.BriefUserInfo),
	}
	gitHubUserResponse = response.User{
		ID:      user.ID,
		Login:   user.Login,
		Avatar:  user.AvatarURL,
		Details: details,
	}

	return gitHubUserResponse, nil
}

func getGitHubPublicUser(getGitHubUserInfoURL string) (interface{}, error) {
	var responseBody, requestError = handleResponse(getGitHubUserInfoURL)
	if requestError != nil {
		return nil, requestError
	}

	var gitHubUser model.User
	unmarshallError := parseResponseBodyIntoModel(responseBody.([]byte), &gitHubUser)

	if unmarshallError != nil {
		return nil, unmarshallError
	}

	return gitHubUser, nil
}

func handleResponse(url string) (interface{}, error) {
	var responseBody, requestError = getResponseBody(url)
	var serverErrorResponse model.ErrorMessage
	_ = parseResponseBodyIntoModel(responseBody, &serverErrorResponse)
	if serverErrorResponse != (model.ErrorMessage{}) {

		return nil, errors.New(serverErrorResponse.Message + "; documentation URL: " + serverErrorResponse.DocumentationURL)
	}

	return responseBody, requestError
}

func getGitHubUserFollowing(userFollowingURL string) (interface{}, error) {
	urlGetAllFollowing := userFollowingURL[:strings.IndexByte(userFollowingURL, '{')]
	var responseBody, requestError = handleResponse(urlGetAllFollowing)
	if requestError != nil {
		return nil, requestError
	}

	var gitHubUserFollowing []model.User
	unmarshallError := parseResponseBodyIntoModel(responseBody.([]byte), &gitHubUserFollowing)
	if unmarshallError != nil {
		unmarshalErrorWithCustomMessage := fmt.Errorf("%w; Custom message: error in parsing user gists", unmarshallError)
		return nil, unmarshalErrorWithCustomMessage
	}

	var following = make([]response.BriefUserInfo, len(gitHubUserFollowing))
	for i, element := range gitHubUserFollowing {
		following[i] = response.BriefUserInfo{
			Login:  element.Login,
			ID:     element.ID,
			Url:    element.URL,
			Avatar: element.AvatarURL,
		}
	}

	return following, nil
}

func getGitHubUserFollowers(userFollowersURL string) (interface{}, error) {
	var responseBody, requestError = handleResponse(userFollowersURL)
	if requestError != nil {
		return nil, requestError
	}

	var gitHubUserFollowers []model.User
	unmarshallError := parseResponseBodyIntoModel(responseBody.([]byte), &gitHubUserFollowers)
	if unmarshallError != nil {
		unmarshalErrorWithCustomMessage := fmt.Errorf("%w; Custom message: error in parsing user gists", unmarshallError)
		return nil, unmarshalErrorWithCustomMessage
	}

	var followers = make([]response.BriefUserInfo, len(gitHubUserFollowers))
	for i, element := range gitHubUserFollowers {
		followers[i] = response.BriefUserInfo{
			Login:  element.Login,
			ID:     element.ID,
			Url:    element.URL,
			Avatar: element.AvatarURL,
		}
	}

	return followers, nil
}

func getGitHubPublicUserGists(userGistsURL string) (interface{}, error) {
	urlGetAllGists := userGistsURL[:strings.IndexByte(userGistsURL, '{')]
	var responseBody, requestError = handleResponse(urlGetAllGists)
	if requestError != nil {
		return nil, requestError
	}

	var gitHubUserGists []model.Gist
	unmarshallError := parseResponseBodyIntoModel(responseBody.([]byte), &gitHubUserGists)
	if unmarshallError != nil {
		unmarshalErrorWithCustomMessage := fmt.Errorf("%w; Custom message: error in parsing user gists", unmarshallError)
		return nil, unmarshalErrorWithCustomMessage
	}

	var gists = make([]response.PublicGist, len(gitHubUserGists))
	for i, element := range gitHubUserGists {
		gists[i] = response.PublicGist{
			ID:  element.ID,
			Url: element.URL,
		}
	}

	return gists, nil
}

func getGitHubPublicUserRepos(userReposURL string) (interface{}, error) {
	var responseBody, requestError = handleResponse(userReposURL)
	if requestError != nil {
		return nil, requestError
	}

	var gitHubUserRepos []model.Repo
	unmarshallError := parseResponseBodyIntoModel(responseBody.([]byte), &gitHubUserRepos)

	if unmarshallError != nil {
		unmarshalErrorWithCustomMessage := fmt.Errorf("%w; Custom message: error in parsing user repositories", unmarshallError)
		return nil, unmarshalErrorWithCustomMessage
	}

	var repos = make([]response.PublicRepo, len(gitHubUserRepos))
	for i, element := range gitHubUserRepos {
		repos[i] = response.PublicRepo{
			Name: element.Name,
			Url:  element.URL,
		}
	}

	return repos, nil
}

func GetGitHubFeed() ([]response.Event, error) {
	var eventsModel []model.Event
	var eventsResponse []response.Event
	var responseBody, requestError = getResponseBody(githubGetAllEvents)
	if requestError != nil {
		return nil, requestError
	}
	unmarshallError := parseResponseBodyIntoModel(responseBody, &eventsModel)
	if unmarshallError != nil {
		return nil, unmarshallError
	}
	eventsResponse = convertGitHubEventModels2CustomServerEventResponses(eventsResponse, eventsModel)

	return eventsResponse, nil
}

func GetGitHubFeedForUser(username string) (interface{}, error) {
	var eventsModel []model.Event
	var eventsResponse []response.Event
	if username == "" {
		return nil, errors.New("provide a github username in request parameters")
	}
	var request = fmt.Sprintf(githubGetAllEventsByUser, username)
	var responseBody, requestError = getResponseBody(request)
	if requestError != nil {
		return nil, requestError
	}

	var serverErrorResponse model.ErrorMessage
	_ = parseResponseBodyIntoModel(responseBody, &serverErrorResponse)
	if serverErrorResponse != (model.ErrorMessage{}) {
		return serverErrorResponse, nil
	}

	unmarshallError := parseResponseBodyIntoModel(responseBody, &eventsModel)

	if unmarshallError != nil {
		return nil, unmarshallError
	}

	eventsResponse = convertGitHubEventModels2CustomServerEventResponses(eventsResponse, eventsModel)

	return eventsResponse, nil
}

func convertGitHubEventModels2CustomServerEventResponses(eventsResponse []response.Event, eventsModel []model.Event) []response.Event {
	eventsResponse = make([]response.Event, len(eventsModel))
	for ind, elem := range eventsModel {
		resp := mapGitHubEventModel2CustomServerEventResponse(elem)
		eventsResponse[ind] = resp
	}

	return eventsResponse
}

func mapGitHubEventModel2CustomServerEventResponse(inputModel model.Event) response.Event {
	responseEvent := response.Event{
		Type: inputModel.Type,
		Actor: response.Actor{
			Id:    inputModel.Actor.Id,
			Login: inputModel.Actor.Login,
		},
		Repo: response.EventRepo{
			Id:   inputModel.Repo.Id,
			Name: inputModel.Repo.Name,
		},
	}

	return responseEvent
}

func getResponseBody(url string) ([]byte, error) {
	resp, err := makeRequest2GitHub(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, responseReadError := readResponseBody(resp.Body)
	if responseReadError != nil {
		return nil, responseReadError
	}

	return body, nil
}

func parseResponseBodyIntoModel(body []byte, responseModel interface{}) error {

	return json.Unmarshal(body, responseModel)
}

func makeRequest2GitHub(url string) (*http.Response, error) {
	req, requestError := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.github.v3+json")
	resp, responseError := client.Do(req)
	if requestError != nil {
		return nil, requestError
	}
	if responseError != nil {
		return nil, responseError
	}

	return resp, nil
}

func readResponseBody(responseBody io.Reader) ([]byte, error) {
	body, err := ioutil.ReadAll(responseBody)
	if err != nil {
		return nil, err
	}

	return body, nil
}
