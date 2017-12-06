package main

import (
	"flag"
	"fmt"

	"github.com/Henry-Sarabia/igdb"
)

var key string

func init() {
	flag.StringVar(&key, "k", "", "API key")
	flag.Parse()
}

func main() {
	if key == "" {
		fmt.Println("No key provided. Please run: characterphotos -k YOUR_API_KEY")
		return
	}

	c := igdb.NewClient(key, nil)

	// Retrieve human character photos
	ch, err := c.Characters.List(nil,
		igdb.SetLimit(20),
		igdb.SetFields("name", "mug_shot"),
		igdb.SetFilter("species", igdb.OpEquals, "1"),     // only humans
		igdb.SetFilter("mug_shot", igdb.OpExists),         // only characters with images
		igdb.SetOrder("created_at", igdb.OrderDescending), // most recently created
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("The 20 Newest Character Photos:\n\n")
	for _, v := range ch {
		img, err := v.Mugshot.SizedURL(igdb.Size1080p, 1) // resize to largest image available
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("%s - %s\n", v.Name, img)
	}

	return
}
