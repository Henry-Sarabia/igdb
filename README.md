# IGDB 

[![GoDoc](https://godoc.org/github.com/Henry-Sarabia/igdb?status.svg)](https://pkg.go.dev/github.com/Henry-Sarabia/igdb?tab=doc) [![Build Status](https://travis-ci.org/Henry-Sarabia/igdb.svg?branch=master)](https://travis-ci.org/Henry-Sarabia/igdb) [![Coverage Status](https://coveralls.io/repos/github/Henry-Sarabia/igdb/badge.svg?branch=master)](https://coveralls.io/github/Henry-Sarabia/igdb?branch=master) [![Go Report Card](https://goreportcard.com/badge/github.com/Henry-Sarabia/igdb)](https://goreportcard.com/report/github.com/Henry-Sarabia/igdb) [![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go)  

<img align="right" src="https://raw.githubusercontent.com/Henry-Sarabia/igdb/master/img/gopherigdb.png">

Communicate with the [Internet Game Database API](https://www.igdb.com/api) quickly and easily
with the **igdb** package. With the **igdb** client, you can retrieve
extensive information on any number of video games, characters, companies, media, artwork
and [much more](https://api-docs.igdb.com/#endpoints). Every IGDB API endpoint is supported!

If you would like to help the Go **igdb** project, please submit a pull request - it's always
greatly appreciated.

## Installation

If you do not have [Go](https://golang.org/) installed yet, you can find installation instructions 
[here](https://golang.org/doc/install). Please note that the package requires Go version
1.13 or later for module support.

To pull the most recent version of **igdb**, use `go get`.

```
go get github.com/Henry-Sarabia/igdb
```

Then import the package into your project as you normally would.

```go
import "github.com/Henry-Sarabia/igdb"
```

Now you're ready to Go.

## Usage

### Creating A Client

Before using the **igdb** package, you need to have an IGDB API key. If you do
not have a key yet, you can sign up [here](https://api.igdb.com/signup).

Create a client with your API key to start communicating with the IGDB API.

```go
client, err := igdb.NewClient("YOUR_API_KEY", nil)
```

If you need to use a preconfigured HTTP client, simply pass its address to the
`NewClient` function.

```go
client, err := igdb.NewClient("YOUR_API_KEY", &custom)
```

### Services

The client contains a distinct service for working with each of the IGDB API
endpoints. Each service has a set of service functions that make specific API
calls to their respective endpoint.  

To start communicating with the IGDB, choose a service and call its service
function. Take the Games service for example.

To search for a Game, use the Search service function.
```go
games, err := client.Games.Search("zelda")
```

To retrieve several Games by their IGDB ID, use the List service function.
```go
games, err := client.Games.List([]int{7346, 1721, 2777})
```

The rest of the service functions work much the same way; they are concise and
behave as you would expect. The [documentation](https://godoc.org/github.com/Henry-Sarabia/igdb#pkg-examples)
contains several examples on how to use each service function.

Service functions by themselves allow you to retrieve a considerable amount of
information from the IGDB but sometimes you need more control over the results
being returned. For this reason, the **igdb** package provides a set of 
flexible functional options for customizing a service function's API query.

### Functional Options

The **igdb** package uses what are called functional options to apply different 
query parameters to service function's API call. Functional options themselves
are merely first order functions that are passed to a service function.

Let's walk through a few different functional option examples.

To set the limit of the amount of results returned from an API query, pass 
SetLimit to the service function.
```go
games, err := client.Games.Search("megaman", SetLimit(15))
```
As you can see, you simply need to pass the functional option as an argument 
to the service function.

To offset the results returned from an API call, pass SetOffset to the service
function.
```go
games, err := client.Games.Search("megaman", SetOffset(15))
```
SetOffset is used to iterate through a large set of results that cannot be 
retrieved in a single API call. In this case, the first 15 results are ignored
so we effectively iterated through to the next set of results by 15.

To set the order of the results returned from an API call, pass SetOrder much
in the same way as the previous examples.
```go
games, err := client.Games.Search("megaman", SetOrder("popularity", igdb.OrderDescending))
```
SetOrder is used to specify in what order you want the results to be retrieved 
in and by what criteria. Here, SetOrder will retrieve the results with the 
highest popularity first.

The remaining functional options are not unlike the examples we covered and 
are further described in the [documentation](https://godoc.org/github.com/Henry-Sarabia/igdb#Option).

### Functional Option Composition

More often than not, you will need to set more than one option for an API query.
Fortunately, this functionality is supported through variadic functions and
functional option composition.

First, service functions are variadic so you can pass in any number of 
functional options.
```go
chars, err := client.Characters.Search(
    "mario",
    SetFields("id", "name", "games"),
    SetFilter("gender", "1"),
    SetLimit(5), 
    )
```
This API call will search the Characters endpoint using the query "mario",
filter out any character that does not have a gender code of 1 (which in this
case represents male), retrieve the id, name, and games fields, and return
only up to 5 of these results.

Second, the **igdb** package provides a `ComposeOptions` function which takes any 
number of functional options as its parameters, composes them into a single
functional option, and returns that composed functional option.
```go
popularOpt := igdb.ComposeOptions(
    igdb.SetLimit(5),
    igdb.SetFields("name"),
	igdb.SetOrder("popularity", igdb.OrderDescending),
)
```
This call to ComposeOptions creates a single functional option that will allow
you to retrieve the names of the top 5 most popular games when passed to the
appropriate service function.

Functional option composition allows you to create custom functional options
that can be reused in different API calls.

Taking the previous example, this can be done in the following way.
```go
PS4, err := c.Games.Index(
		popularOpt,
		igdb.SetFilter("platforms", igdb.OpEquals, "48"),    // filter out games not on PS4
    )
    
XBOX, err := c.Games.Index(
		popularOpt, 
		igdb.SetFilter("platforms", igdb.OpEquals, "49"),    // filter out games not on XB1
    )
```
This example has two service function calls that each utilize the previously
composed functional option in the same way but for different platforms. The 
first function retrieves the top 5 most popular PS4 games while the second
function retrieves the top 5 most popular XB1 games.

Functional option composition reduces duplicate code and helps keep your code
DRY. You can even compose newly composed functional options for even more
finely grained control over similar API calls.

## Examples

The repository contains several example mini-applications that demonstrate
how one might use the **igdb** package.

* [Mini Applications](https://github.com/Henry-Sarabia/igdb/tree/master/examples)
* [Documentation Examples](https://godoc.org/github.com/Henry-Sarabia/igdb#pkg-examples)

If you have used the **igdb** package for a project and would like to have it
featured here as a reference for new users, please submit an issue and I'll be
sure to add it.

## Contributions

If you would like to contribute to this project, please adhere to the following
guidelines.

* Submit an issue describing the problem.
* Fork the repo and add your contribution.
* Add appropriate tests.
* Run go fmt, go vet, and golint.
* Prefer idiomatic Go over non-idiomatic code.
* Follow the basic Go conventions found [here](https://github.com/golang/go/wiki/CodeReviewComments).
* If in doubt, try to match your code to the current codebase.
* Create a pull request with a description of your changes.

Again, contributions are greatly appreciated!

## Special Thanks

<img align="right" src="https://github.com/Henry-Sarabia/igdb/blob/master/img/gopherthanks.png">

* You for your interest
* John for the IGDB Gopher
* Peter for the "Thank You" Gopher
* Dave Cheney for his [article](https://dave.cheney.net/2014/10/17/functional-options-for-friendly-apis)
on functional options
* The [DiscordGo](https://github.com/bwmarrin/discordgo) and Go [Spotify](https://github.com/zmb3/spotify)
projects for inspiring me to create my own open source package for others to enjoy
* The [Awesome Go](https://github.com/avelino/awesome-go) project for so many
references to admire
* The awesome people in the IGDB community who are always open to questions
