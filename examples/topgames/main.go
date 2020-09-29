package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Henry-Sarabia/igdb"
)

var key string
var token string

func init() {
	flag.StringVar(&key, "k", "", "Client-ID")
	flag.StringVar(&token, "t", "", "AppAccessToken")
	flag.Parse()
}

func main() {
	if key == "" {
		fmt.Println("No key provided. Please run: topgames -k YOUR_CLIENT_ID -t YOUR_APP_ACCESS_TOKEN")
		return
	}
	if token == "" {
		fmt.Println("No token provided. Please run: topgames -k YOUR_CLIENT_ID -t YOUR_APP_ACCESS_TOKEN")
		return
	}

	c := igdb.NewClient(key, token, nil)

	// Composing options set to retrieve top 5 popular results
	byPop := igdb.ComposeOptions(
		igdb.SetLimit(5),
		igdb.SetFields("name", "cover"),
		igdb.SetOrder("hypes", igdb.OrderDescending),
		igdb.SetFilter("category", igdb.OpEquals, "0"),
		igdb.SetFilter("cover", igdb.OpNotEquals, "null"),
	)

	// Retrieve PS4 inter-console exclusives
	PS4, err := c.Games.Index(
		byPop, // top 5 popular results
		igdb.SetFilter("platforms", igdb.OpEquals, "48"), // only PS4 games
	)
	if err != nil {
		log.Fatal(err)
	}

	// Retrieve XB1 inter-console exclusives
	XBOX, err := c.Games.Index(
		byPop, // top 5 popular results
		igdb.SetFilter("platforms", igdb.OpEquals, "49"), // only XB1 games
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Top 5 PS4 Games:")
	for _, game := range PS4 {
		cover, err := c.Covers.Get(game.Cover, igdb.SetFields("image_id")) // retrieve cover IDs
		if err != nil {
			log.Fatal(err)
		}
		img, err := cover.SizedURL(igdb.Size1080p, 1) // resize to largest image available
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s - %s\n", game.Name, img)
	}

	fmt.Println("\nTop 5 XBOX Games:")
	for _, game := range XBOX {
		cover, err := c.Covers.Get(game.Cover, igdb.SetFields("id", "image_id")) // retrieve cover IDs
		if err != nil {
			log.Fatal(err)
		}
		img, err := cover.SizedURL(igdb.Size1080p, 1) // resize to largest image available
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s - %s\n", game.Name, img)
	}

	return
}
