package main

import (
	"log"
	"net/http"
	"os"
	"time"
)

const (
	deezerPlaylistId  = 1234567891
	deezerLimit       = 560
	spotifyPlaylistId = "abcdefghijklmnop"
	spotifyLimit      = 10
)

var (
	deezerAccessToken  = os.Getenv("DEEZER_ACCESS_TOKEN")
	deezerUrl          = "https://api.deezer.com/playlist/%s/tracks?limit=%s&access_token=%s"
	spotifyToken       = os.Getenv("SPOTIFY_TOKEN")
	spotifySearchUrl   = "https://api.spotify.com/v1/search?q=%s&type=track&limit=%s"
	spotifyAddItemsUrl = "https://api.spotify.com/v1/playlists/%s/tracks?uris=%s"
)

func main() {
	deezerClient := http.Client{Timeout: time.Duration(4) * time.Second}
	spotifyClient := http.Client{Timeout: time.Duration(10) * time.Second}
	syncDeezerPlaylistToSpotify(deezerClient, spotifyClient)
}

func syncDeezerPlaylistToSpotify(deezerClient http.Client, spotifyClient http.Client) {
	spotifyTracksUris := []string{}
	notFoundDeezerTracks := []DeezerTrack{}
	deezerTracks := getTracksFromDeezerPlaylist(deezerClient)

	for _, deezerTrack := range deezerTracks {
		spotifyResponse := getSpotifyTrackFor(spotifyClient, deezerTrack.Title, deezerTrack.DeezerArtist.Name)
		if len(spotifyResponse.SpotifyTracks.SpotifyItems) > 0 {
			spotifyTracksUris = append(spotifyTracksUris, spotifyResponse.SpotifyTracks.SpotifyItems[0].Uri)
		} else {
			notFoundDeezerTracks = append(notFoundDeezerTracks, deezerTrack)
		}
	}
	log.Println("Deezer tracks not found in Spotify:", notFoundDeezerTracks)
	// Updating by batch otherwise it won't work because of status code: 414 (URI too long)
	for i := 0; i < len(spotifyTracksUris); i = i + 90 {
		spotifyTracksUrisSliced := spotifyTracksUris[i : i+90]
		req, err := createSpotifyAddItemsRequest(spotifyTracksUrisSliced)
		response, err := spotifyClient.Do(req)
		if err != nil {
			log.Println("Error while adding track to a playlist using Spotify API:", err)
		}
		log.Println("Adding items to the playlist with status code:", response.StatusCode)
	}
	log.Println("Synchronization done")
}
