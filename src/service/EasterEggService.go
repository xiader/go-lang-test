package service

import (
	"appstud.com/github-core/src/models"
	"math"
)

func GetEasterEggs() []models.EasterEggResponse {
	return easterEggs
}

var easterEggs = []models.EasterEggResponse{
	{
		Name:      "My mom is in love with me",
		Version:   "1.0",
		Timestamp: -446723100,
	},
	{
		Name:      "I go to the future and my mom end up with the wrong guy",
		Version:   "2.0",
		Timestamp: 1445470140,
	},
	{
		Name:      "I go to the past and you will not believe what happens next",
		Version:   "3.0",
		Timestamp: math.MinInt64 + 1,
	},
}
