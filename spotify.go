package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type SpotifyResponse struct {
	SpotifyTracks SpotifyTrack `json:"tracks"`
}

type SpotifyTrack struct {
	SpotifyItems []SpotifyItem `json:"items"`
}

type SpotifyItem struct {
	Uri string
}

func getSpotifyTrackFor(spotifyClient http.Client, title string, artist string) SpotifyResponse {
	req, err := createSpotifySearchRequest(title, artist)
	res, _ := spotifyClient.Do(req)
	if err != nil {
		log.Println("Error while calling Spotify Search API:", err)
	}

	body, _ := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	spotifyResponse := SpotifyResponse{}
	json.Unmarshal(body, &spotifyResponse)
	return spotifyResponse
}

func createSpotifySearchRequest(track string, artist string) (*http.Request, error) {
	spotifyQuery := url.QueryEscape(fmt.Sprintf("track:%s artist:%s", track, artist))
	spotifyUrl := fmt.Sprintf(spotifySearchUrl, spotifyQuery, strconv.Itoa(spotifyLimit))
	req, err := http.NewRequest("GET", spotifyUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", spotifyToken))
	return req, err
}

func createSpotifyAddItemsRequest(trackUris []string) (*http.Request, error) {
	spotifyTrackUris := url.QueryEscape(strings.Join(trackUris, ","))
	spotifyUrl := fmt.Sprintf(spotifyAddItemsUrl, spotifyPlaylistId, spotifyTrackUris)
	req, err := http.NewRequest("POST", spotifyUrl, nil)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", spotifyToken))
	return req, err
}
