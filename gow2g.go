package gow2g

import (
	"net/http"
	"strings"
)

var httpClient = &http.Client{
	CheckRedirect: func(req *http.Request, via []*http.Request) error {
		return nil
	},
}

type W2GRoom struct {
	ID string
}

func (room W2GRoom) URL() string {
	return "https://www.watch2gether.com/rooms/" + room.ID
}

func GenerateW2GRoom() (*W2GRoom, error) {
	if response, err := httpClient.Post("https://www.watch2gether.com/rooms/create", "text/html; charset=utf-8", nil); err == nil {
		roomId := strings.Replace(response.Request.URL.EscapedPath(), "/rooms/", "", -1)

		return &W2GRoom{
			ID: roomId,
		}, nil
	} else {
		return nil, err
	}
}
