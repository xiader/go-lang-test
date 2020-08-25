package service

import (
	"appstud.com/github-core/src/models/github/model"
	"appstud.com/github-core/src/models/github/response"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const githubGetAllEvents = "https://api.github.com/events"

func GetGitHubFeed() ([]response.Event, error) {
	var eventsModel []model.Event
	var eventsResponse []response.Event
	var request = githubGetAllEvents
	var respError = getResponse(request, &eventsModel)
	if respError != nil {
		return eventsResponse, respError
	}

	eventsResponse = make([]response.Event, len(eventsModel))
	for ind, elem := range eventsModel {
		resp := transformGitHubEventModel2CustomServerEventResponse(elem)
		eventsResponse[ind] = resp
	}

	return eventsResponse, nil
}

func transformGitHubEventModel2CustomServerEventResponse(inputModel model.Event) response.Event {
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

func getResponse(url string, responseModel interface{}) error {
	resp, requestError := http.Get(url)
	if requestError != nil {
		return requestError
	}
	defer resp.Body.Close()

	body, responseReadError := readResponseBody(resp)
	if responseReadError != nil {
		return responseReadError
	}

	return json.Unmarshal(body, responseModel)
}

func readResponseBody(resp *http.Response) ([]byte, error) {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
