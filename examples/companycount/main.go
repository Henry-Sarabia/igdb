package main

import (
	"flag"
	"fmt"

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
		fmt.Println("No key provided. Please run: companycount -k YOUR_CLIENT_ID -t YOUR_APP_ACCESS_TOKEN")
		return
	}
	if token == "" {
		fmt.Println("No token provided. Please run: companycount -k YOUR_CLIENT_ID -t YOUR_APP_ACCESS_TOKEN")
		return
	}

	c := igdb.NewClient(key, token, nil)

	// Count number of US companies
	USA, err := c.Companies.Count(igdb.SetFilter("country", igdb.OpEquals, "840"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Count number of UK companies
	UK, err := c.Companies.Count(igdb.SetFilter("country", igdb.OpEquals, "826"))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Count number of JP companies
	JP, err := c.Companies.Count(igdb.SetFilter("country", igdb.OpEquals, "392"))
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Print("Number of Companies Based In Specific Countries:\n\n")
	fmt.Println("USA -", USA)
	fmt.Println("UK -", UK)
	fmt.Println("Japan -", JP)

	return
}
