# Top PS4 and XB1 Games Example

This example makes use of the igdb client's Games service to retrieve the
top 5 most popular inter-console exclusive games and their respective cover 
art for the PS4 and the Xbox One. The example demonstrates the use of the
following key functions:
* ComposeOptions
* Games.List
* SetLimit
* SetFields
* SetOrder
* SetFilter
* SizedURL

### Installation
If you do not have Go installed yet, you can find installation instructions 
[here](https://golang.org/doc/install).

To build this example, navigate to the containing directory and use the Go 
build tool as follows.

```
go build
```
topg
This command will have built the example into an executable and can now be run
from the command line.

### Usage
To run this example you will need a valid IGDB API key. If you do not have one
yet, sign up [here](https://api.igdb.com/signup).

Now, make sure you are in the same containing directory and use the following 
command with your API key.

```
topgames -k YOUR_API_KEY
```

If your API key is valid and no unforeseen errors occur, you should see some
output that resembles the following:

```
Top 5 PS4 Games:
Dead Island 2 - https://images.igdb.com/igdb/image/upload/t_1080p/lumgkti6rht3evlbu8xw.jpg
Pyre - https://images.igdb.com/igdb/image/upload/t_1080p/zl02iwvbyyp28wquk8br.jpg
The Last of Us: Part II - https://images.igdb.com/igdb/image/upload/t_1080p/murzziwvvmzacglj5ch5.jpg
Horizon: Zero Dawn - https://images.igdb.com/igdb/image/upload/t_1080p/ayacax1kkj9f76z0oepo.jpg
Lost Sphear - https://images.igdb.com/igdb/image/upload/t_1080p/rgxz637m1rac8ao7g3pi.jpg

Top 5 XBOX Games:
Hello Neighbor - https://images.igdb.com/igdb/image/upload/t_1080p/zsvyzrqpbuvtfnpdagfp.jpg
Ashen - https://images.igdb.com/igdb/image/upload/t_1080p/iwbh1uu5zbreqwrk7hlr.jpg
PLAYERUNKNOWN'S BATTLEGROUNDS - https://images.igdb.com/igdb/image/upload/t_1080p/lvoic2oakbklg2dytgpa.jpg
Ori and the Blind Forest - https://images.igdb.com/igdb/image/upload/t_1080p/uqpeoercdo8rjwbsn5b3.jpg
Cuphead - https://images.igdb.com/igdb/image/upload/t_1080p/gtzfjc9pipa6s7v7m68g.jpg
```

If you follow the link to one of the images, it will show the cover art of the
appropriate game, like so:
<img src="https://images.igdb.com/igdb/image/upload/t_1080p/uqpeoercdo8rjwbsn5b3.jpg">