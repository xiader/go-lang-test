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
	"time"
)

const githubGetAllEvents = "https://api.github.com/events"
const githubGetAllEventsByUser = "https://api.github.com/users/%s/events"

var client = &http.Client{Timeout: time.Minute}

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
		Repo: response.Repo{
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
