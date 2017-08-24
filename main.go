package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	rootURL string = "https://api-2445582011268.apicast.io/games/"
)

// ID is an unsigned 64-bit integer
type ID int

// URL is
type URL string

// Image is a struct that holds the ID to reach the image along with its dimensions
type Image struct {
	URL    URL    `json:"url"`
	ID     string `json:"cloudinary_id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
}

// Video is a struct that holds the name of a video along with its ID.
type Video struct {
	Name string `json:"name"`
	ID   string `json:"video_id"` // Youtube slug
}

func main() {
	g, err := getGames()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(g)
}

func getGames() ([]Game, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", rootURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", APIkey)
	req.Header.Add("Accept", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var g []Game

	err = json.Unmarshal(b, &g)
	if err != nil {
		return nil, err
	}
	return g, nil
}
