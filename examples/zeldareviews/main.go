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
		fmt.Println("No key provided. Please run: topgames -k YOUR_API_KEY")
		return
	}

	c := igdb.NewClient(key, nil)

	// Search for Zelda reviews
	rev, err := c.Reviews.Search(
		"zelda", // zelda query
		igdb.SetOrder("likes", igdb.OrderDescending), // ordered by most liked
		igdb.SetFields("title", "game", "positive_points", "negative_points"),
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Zelda Reviews Found\n\n")
	for _, v := range rev {
		// Retrieve specific Zelda game using Review.Game ID
		g, err := c.Games.Get(v.Game, igdb.SetFields("name"))
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("%s - %s\n", g.Name, v.Title)
		fmt.Println("Positives:", v.PositivePoints)
		fmt.Println("Negatives:", v.NegativePoints)
		fmt.Println()
	}
}
