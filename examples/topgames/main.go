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

	// Composing options set to retrieve top 5 popular results
	byPop := igdb.ComposeOptions(
		igdb.SetLimit(5),
		igdb.SetFields("name"),
		igdb.SetOrder("popularity", igdb.OrderDescending),
		igdb.SetFilter("version_parent", igdb.OpNotExists),
	)

	// Retrieve PS4 inter-console exclusives
	PS4, err := c.Games.List(nil,
		byPop, // top 5 popular results
		igdb.SetFilter("platforms", igdb.OpIn, "48"),    // only PS4 games
		igdb.SetFilter("platforms", igdb.OpNotIn, "49"), // filter out XB1
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Retrieve XB1 inter-console exclusives
	XBOX, err := c.Games.List(nil,
		byPop, // top 5 popular results
		igdb.SetFilter("platforms", igdb.OpIn, "49"),    // only XB1 games
		igdb.SetFilter("platforms", igdb.OpNotIn, "48"), // filter out PS4
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Top 5 PS4 Games")
	for _, v := range PS4 {
		fmt.Println(v.Name)
	}

	fmt.Println("\nTop 5 XBOX Games")
	for _, v := range XBOX {
		fmt.Println(v.Name)
	}
}
