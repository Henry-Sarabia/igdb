package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/Henry-Sarabia/igdb"
)

const (
	rootURL string = "https://api-2445582011268.apicast.io/games/"
)

func main() {
	g, err := getGames()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(g)
}

func getGames() ([]igdb.Game, error) {
	client := http.DefaultClient
	req, err := http.NewRequest("GET", rootURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("user-key", igdb.APIkey)
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

	var g []igdb.Game

	err = json.Unmarshal(b, &g)
	if err != nil {
		return nil, err
	}
	return g, nil
}
