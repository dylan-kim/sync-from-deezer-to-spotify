package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

type DeezerResponse struct {
	Tracks []DeezerTrack `json:"data"`
}

type DeezerTrack struct {
	Title        string
	DeezerArtist DeezerArtist `json:"artist"`
}

type DeezerArtist struct {
	Name string
}

func getTracksFromDeezerPlaylist() []DeezerTrack {
	deezerClient := http.Client{Timeout: time.Duration(4) * time.Second}
	deezerUrl = fmt.Sprintf(deezerUrl, strconv.Itoa(deezerPlaylistId), strconv.Itoa(deezerLimit), deezerAccessToken)
	res, err := deezerClient.Get(deezerUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	deezerResponse := DeezerResponse{}
	json.Unmarshal(body, &deezerResponse)
	return deezerResponse.Tracks
}
