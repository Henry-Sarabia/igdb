# IGDB 

[![GoDoc](https://godoc.org/github.com/Henry-Sarabia/igdb?status.svg)](https://godoc.org/github.com/Henry-Sarabia/igdb) [![Go Report Card](https://goreportcard.com/badge/github.com/Henry-Sarabia/igdb)](https://goreportcard.com/report/github.com/Henry-Sarabia/igdb) 

<img align="right" src="https://raw.githubusercontent.com/Henry-Sarabia/igdb/master/img/igdbicon.png ">

Communicate with the [Internet Game Database API](https://api.igdb.com/) quickly and easily
with the `igdb` [Go](https://golang.org/) package. With the `igdb` package, you can retrieve
extensive information on any number of video games, characters, companies, reviews, media,
and [much more](https://igdb.github.io/api/endpoints/). Every IGDB API endpoint is supported!

If you would like to help the Go `igdb` project, please submit a pull request - it's always
greatly appreciated.

## Installation

If you do not have Go installed yet, you can find installation instructions 
[here](https://golang.org/doc/install).

To pull the most recent version of `igdb`, use `go get`.

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

Before using the `igdb` package, you need to have an IGDB API key. If you do
not have a key yet, you can sign up [here](https://api.igdb.com/signup).

Create a client with your API key to start communicating with the IGDB API.

```go
client, err := igdb.NewClient("YOUR_API_KEY", nil)
```

If you need to use a preconfigured HTTP client, simply pass its address to the
NewClient function.

```go
client, err := igdb.NewClient("YOUR_API_KEY", &customClient)
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
function as you would expect.

Service functions by themselves allow you to retrieve a considerable amount of
information from the IGDB but sometimes you need more control over the results
being returned. For this reason, the `igdb` package provides a set of 
flexible functional options for customizing a service function's API call.

### Functional Options

The `igdb` package uses what are called functional options to apply different 
query parameters to service function's API call. Functional options themselves
are merely first order functions that are passed to a service function.

Let's walk through a few different functional option examples.

To set the limit of the amount of results returned from an API call, pass 
SetLimit to the service function.
```go
revs, err := client.Reviews.Search("mario", SetLimit(25))
```
As you can see, you simply need to pass the functional option as an argument 
to the service function.

To offset the results returned from an API call, pass SetOffset to the service
function.
```go
revs, err := client.Reviews.Search("mario", SetOffset(10))
```
SetOffset is used to iterate through a large set of results that cannot be 
retrieved in a single API call. In this case, the first 10 results are ignored
so we effectively iterated through to the next several results by 10.

To set the order of the results returned from an API call, pass SetOrder much
in the same way as the previous examples.
```go
revs, err := client.Reviews.Search("mario", SetOrder("views", igdb.OrderDescending))
```
SetOrder is used to specify in what order we want the results to be retrieved 
in and by what criteria. Here, SetOrder will retrieve the results with the 
most views first.

The rest of the functional options are not unlike the examples we covered and 
are further described in the [documentation](https://godoc.org/github.com/Henry-Sarabia/igdb#FuncOption).
