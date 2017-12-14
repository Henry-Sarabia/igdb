# Company Count Example

This example makes use of the igdb client's Company service to count the number
of video game Companies based in the US, UK, and Japan. The example demonstrates 
the use of the following key functions:
* Companies.Count
* SetFilter

### Installation
If you do not have Go installed yet, you can find installation instructions 
[here](https://golang.org/doc/install).

To build this example, navigate to the containing directory and use the Go 
build tool as follows.

```
go build
```

This command will have built the example into an executable and can now be run
from the command line.

### Usage
To run this example you will need a valid IGDB API key. If you do not have one
yet, sign up [here](https://api.igdb.com/signup).

Now, make sure you are in the same containing directory and use the following
command with your API key.

```
companycount -k YOUR_API_KEY
```

If your API key is valid and no unforseen errors occur, you should see some
output that resembles the following:

```
Number of Companies Based In Specific Countries:

USA - 438
UK - 173
Japan - 202
```