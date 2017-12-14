# Zelda Reviews Example

This example makes use of the igdb cleint's Reviews service to search the IGDB
for the most liked reviews that match the "zelda" search query. The example
demonstrates the use of the following key functions:
* Reviews.Search
* Games.Get
* SetOrder
* SetFields

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
zeldareviews -k YOUR_API_KEY
```

If your API key is valid and no unforseen errors occur, you should see some
output that resembles the following:

```
Zelda Reviews Found

The Legend of Zelda: Ocarina of Time - Zelda Reviewathon #1
Positives: There is a glitch where you can actually wear a pot!
Negatives: There is a glitch where you can actually wear a pot....

The Legend of Zelda: The Wind Waker - Zelda Reviewathon #3
Positives: The sailing could get tedious but mostly I enjoyed it.
Negatives: filling out the map was like homework.


The Legend of Zelda: Majora's Mask - Zelda reviewathon #2
Positives: There is a guy living in a toilet who you can give paper to in return for hearts.
Negatives: The moons face gives you nightmares.

The Legend of Zelda: Twilight Princess - Zelda reviewathon #4
Positives: Pretty environments and characters. Wolf form was awesome.
Negatives: Tossing goats

The Legend of Zelda: Skyward Sword - Zelda reviewathon #5
Positives: flying was amusing.
Negatives: The map was so small.
```