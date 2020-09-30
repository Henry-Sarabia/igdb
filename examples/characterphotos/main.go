package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/Henry-Sarabia/igdb/v2"
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
		fmt.Println("No key provided. Please run: characterphotos -k YOUR_CLIENT_ID -t YOUR_APP_ACCESS_TOKEN")
		return
	}
	if token == "" {
		fmt.Println("No token provided. Please run: characterphotos -k YOUR_CLIENT_ID -t YOUR_APP_ACCESS_TOKEN")
		return
	}

	c := igdb.NewClient(key, token, nil)

	// Retrieve human character photos
	ch, err := c.Characters.Index(
		igdb.SetLimit(20),
		igdb.SetFields("name", "mug_shot"),
		igdb.SetFilter("species", igdb.OpEquals, "1"),        // only humans
		igdb.SetFilter("mug_shot", igdb.OpNotEquals, "null"), // only characters with images
		igdb.SetOrder("created_at", igdb.OrderDescending),    // most recently created
	)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("The 20 Newest Character Photos:\n\n")
	for _, v := range ch {
		mugshot, err := c.CharacterMugshots.Get(v.MugShot, igdb.SetFields("image_id")) // retrieve mugshot ID
		if err != nil {
			log.Fatal(err)
		}

		if mugshot.ImageID == "" {
			continue
		}

		img, err := mugshot.SizedURL(igdb.Size1080p, 1) // resize to largest image available
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s - %s\n", v.Name, img)
	}

	return
}
