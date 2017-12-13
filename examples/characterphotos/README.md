# Character Photos Example

This example makes use of the igdb client's Character service to retrieve 
images of the 20 most recently added characters from the IGDB. The example
demonstrates the use of the following key functions:
* Characters.List()
* SetLimit()
* SetFields()
* SetFilter()
* SetOrder()
* SizedURL()

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

Now, make sure you are in the same containing directory and use
the following command with your API key.

```
characterphotos -k YOUR_API_KEY
```

If your API key is valid and no unforseen errors occur, you should see some
output that resembles the following:

```
Ashley Thorne - https://images.igdb.com/igdb/image/upload/t_1080p/dwxdcqu80xrrdnsa9mxp.jpg
Dr. Sarah Brook - https://images.igdb.com/igdb/image/upload/t_1080p/yrdhbindmfzwuewwjwng.jpg
Princess Zelda - https://images.igdb.com/igdb/image/upload/t_1080p/q64bgyomnbk5spyikkex.jpg
Nicola - https://images.igdb.com/igdb/image/upload/t_1080p/kl5irzci1bj4ymvnzhjl.jpg
Mitsuru Kirijo - https://images.igdb.com/igdb/image/upload/t_1080p/rjnvgdl0aifmjrtrjgj9.jpg
Mr. Ekoda - https://images.igdb.com/igdb/image/upload/t_1080p/ttw7ocki33zjbhvngwrb.jpg
Officer Kurosawa - https://images.igdb.com/igdb/image/upload/t_1080p/zrbtld8pjyxedhanu11b.jpg
Isako Toriumi - https://images.igdb.com/igdb/image/upload/t_1080p/evd5usxhpjngl4rfsd7q.jpg
Natsuki Moriyama - https://images.igdb.com/igdb/image/upload/t_1080p/nljtysljky3sag5s5g44.jpg
Takeharu Kirijo - https://images.igdb.com/igdb/image/upload/t_1080p/ilqbigrpm338gkpft3g8.jpg
Shuji Ikutsuki - https://images.igdb.com/igdb/image/upload/t_1080p/ixy4q3ee8zdlm8esqakx.jpg
Chidori Yoshino - https://images.igdb.com/igdb/image/upload/t_1080p/ok9whyfiguoukcoghito.jpg
Takaya Sakaki - https://images.igdb.com/igdb/image/upload/t_1080p/jlzuiw93jxoyfa0oare7.jpg
Ken Amada - https://images.igdb.com/igdb/image/upload/t_1080p/ak5yzxhcgvnndiuuxsm9.jpg
Shinjiro Aragaki - https://images.igdb.com/igdb/image/upload/t_1080p/utnsmywxm5alnaclp2qy.jpg
Fuuka Yamagishi - https://images.igdb.com/igdb/image/upload/t_1080p/g0tg6ki63fmosh3aj4ji.jpg
Akihiko Sanada - https://images.igdb.com/igdb/image/upload/t_1080p/ysotmcmuusyqi6u4xhnm.jpg
Junpei Iori - https://images.igdb.com/igdb/image/upload/t_1080p/ih2ey8d0y66orat7tmuy.jpg
Yukari Takeba - https://images.igdb.com/igdb/image/upload/t_1080p/u3wddae7y6faxamrpoty.jpg
Heroine - https://images.igdb.com/igdb/image/upload/t_1080p/mds3odgd72u5x3vgnkdu.jpg
```

If you follow the link to one of the images, it will show a photo of the 
respective character like so:
<img src="https://images.igdb.com/igdb/image/upload/t_1080p/u3wddae7y6faxamrpoty.jpg">