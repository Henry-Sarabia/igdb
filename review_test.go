package igdb

import (
	"net/http"
	"reflect"
	"strings"
	"testing"
)

const getReviewResp = `
[ {
	"id": 1462,
	"created_at": 1426759830353,
	"updated_at": 1426759830495,
	"username": "brutus86",
	"slug": "almost-perfect",
	"url": "https://www.igdb.com/games/mario-kart-8/reviews/almost-perfect",
	"title": "Almost perfect!",
	"game": 2350,
	"category": 1,
	"likes": 0,
	"views": 250,
	"rating_category": 3,
	"platform": 41,
	"introduction": "Nintendo has made the best Mario Kart-game yet! ",
	"content": "If you want a good looking game with partyfeatures and some of the best racingtracks to this date, then Mario Kart 8 is the game for you. There is almost nothing to negative about the game, except some details Nintendo should take care of. Thanks to the possibillity of downloading ad-ons via eShop we can now enjoy a Mario Kart game longer! \r\n\r\nThere is one BIG downside thou, and thats the battlearenas, there is no one! If you like battle then you should probably hang on to the old games! In Battle on Mario Kart 8 there is just the ordinary tracks and thats a BIG downside for splitscreen-sessions!",
	"conclusion": "Mario Kart 8 introduces downloadable tracks, something we've missed in all these games! It's gourgeous graphic! And it's a lot of fun to play with your friends in splitscreen-mode (or online).",
	"positive_points": "Graphic, expandable, controls",
	"negative_points": "No Battle arenas"
}]
`

const getReviewsResp = `
[{
	"id": 1571,
	"created_at": 1443201603866,
	"updated_at": 1443201605524,
	"username": "ZUPERFLY",
	"slug": "super-mario-maker",
	"url": "https://www.igdb.com/games/super-mario-maker/reviews/super-mario-maker",
	"title": "Super Mario Maker",
	"game": 7339,
	"category": 1,
	"likes": 1,
	"views": 248,
	"rating_category": 3,
	"platform": 41,
	"introduction": "This game is great. Lots of fun building and playing different levels. ",
	"content": "This game is great. I got it when it was released. I personally love games where I get to create things. However, even though I got to. I had to unlock all the parts. The way to unlock the pieces was lame. You have to wait. Not sure if it was because of it being new. Or if it has changed. I don't know. But once you get all the pieces or not even having all the pieces. I can make some really good levels and upload lots of challenging ones. I made a maze level that is a huge hit.  It's great making a level and checking out later that 216 people played it. Commenting on it and giving it stars. Even if they hate it. It's still getting the word out that this level is fun. \r\nAs far as laying levels goes. I'm enjoying it way more than I expected. 100-Man Mario thing is great. Playing level after level. And skipping the ones that are just too crazy to complete. You have to complete your level to upload it. Which is fair. And having the music play over everything is great. My Xbox friends playing Evolve with Mario Theme playing over it. Ruins the scary aspect of their game meanwhile I'm cursing trying to get over a mountain of Goombas and Koopas.\r\n",
	"conclusion": "Overall, this game is probably one of the better Mario games out there. Lots fun to play and create.",
	"positive_points": "-Music\r\n-Creativity",
	"negative_points": "-Unlocking pieces"
},
{
	"id": 65,
	"created_at": 1411679075325,
	"updated_at": 1411679075724,
	"username": "ZUPERFLY",
	"slug": "super-smash-bros-brawl",
	"url": "https://www.igdb.com/games/super-smash-bros-brawl/reviews/super-smash-bros-brawl",
	"title": "Super Smash Bros. Brawl",
	"game": 1628,
	"category": 1,
	"likes": 0,
	"views": 281,
	"rating_category": 3,
	"platform": 5,
	"introduction": "I played this game so much I don't really know where to begin. So many things to do. Yet I've done them all.",
	"content": "This game was as great as Super Smash Bros. can get. The greatest thing this game had from the last two was a really good story. It was really cool. I thought it was the best. Then to add classic mode and all star mode in too. Plus stadium events and games. It was great. \r\n\r\nThe gameplay was great. It flowed nicely. It was clear. The changes they made were good and bad. New characters are always good. But cutting characters is a tough choice. Or changing a characters attacks. Mario's spin move for his down special attack was better than the Brawl version where he pulls FLUDD out and shoots water. It was just better in Melee. But the Smash Balls and Super Smashes was a neat new trick. It gave each character more uniqueness. \r\n\r\nAnd the other weapons are great. I didn't like the timers. I do and always will love the Fan. It's the best weapon in the game. My friends are terrified of me when I get the Fan. The soccer ball was neat. The Team Heal Ball was useless. We rarely played team games. So we never got to really use it. Assist Trophies and Poke Balls are fun. I did like collecting CDs. Trophy collecting is just there. And stickers were useless. They didn't seem to have that much of a purpose and you never know if you're close or not to collecting them all. \r\n\r\nThe stages were okay. A few stages bugged me so I keep them turned off. Custom Stages are really cool. I love custom stuff. So I have several of them. The music for the game was great. However, I personally don't like the Brawl theme as much as the original and Melee theme. \r\n\r\n100 Man Brawl is difficult. You can't turn off items. Which doesn't sound bad until you're at 15 and you have 200% damage. Then a bomb spawns right on top of you and you die instantly and lose all that progress. ",
	"conclusion": "So overall, this game is awesome. Not as awesome as Melee, but still awesome.",
	"positive_points": "-smash balls\r\n-Subspace Emissary\r\n-gameplay\r\n-custom levels",
	"negative_points": "-timers\r\n-Mario"
}]
`

const searchReviewsResp = `
[{
	"id": 80,
	"created_at": 1416855533948,
	"updated_at": 1416855534080,
	"username": "Tobias",
	"slug": "zelda-reviewathon-number-1",
	"url": "https://www.igdb.com/games/the-legend-of-zelda-ocarina-of-time/reviews/zelda-reviewathon-number-1",
	"title": "Zelda Reviewathon #1",
	"game": 1029,
	"category": 1,
	"likes": 2,
	"views": 263,
	"rating_category": 3,
	"platform": 4,
	"introduction": "Thus begins my reviews of all the relevant Zelda games for consoles. I say relevant because I am simply refusing to acknowledge the latest multiplier atrocity for the WiiU on moral grounds.\r\nThe Legend of Zelda is one of Nintendo's flagships which it wheels out of the closet every few years with another update from the hugely successful 'Ocarina of Time' game for n64 in the late 90s. Since then the franchise has gotten progressively worse in a very uniform pattern I have coined: 'The nostalgia vortex'. The phenomenon occurs when you have such a brain fryingly good time with a game as a child that you would wait untill mum goes to bed before turning it on again with the volume at 1% so you could still hear the jingly noises that made up so much of the fun. Then, as the years go by and you develop curious hair and an interest in the opposite sex that you forget about the game for a while. Eventually after being rejected by the opposite sex and the hair thing has also lost it's novelty, you blow the dust out of those old cartridges and see if the old girl still works. Amazing! It does! And's what's this? It's still LOADS of fun to play. The Vortex is not yet explained by this alone but I will continue it's explanation in the second installment of these reviews.",
	"content": "Ocarina of time (or OoT for those of us who need to save the precious seconds for more pressing matters) is to many of us the game of our childhood. It follows the endeavors of Link (not Zelda!) who is awakened to find that faeries, talking trees and midgets all want him to leave town and go fight dangerous monsters. So off he trots collecting his signature arsenal of weapons and collecting coloured 'things' until he has enough things to fight Ganon who is trying to take over a world seemingly tailored to give minor obstacles to anyone attempting to go anywhere or do anything. You start off as child link and after some fapping about you get enough coloured 'things' to become 'adult link' in the future using the Master Sword and the OoT. As adult link your gaming experience is refreshed by being able to use different abilities and weapons. After a bit more fapping about you become a bit in-game-nostalgic when BAM you are able to return to being child link again. It's win win win all the way home from there.",
	"conclusion": "The game actually has a lot of faults but for 1998 it was mind blowing and special. The nostalgia vortex is kick started with this game and with expectations high for the next Zelda game, the young gamers of the new millennium were about to get their tiny minds blown again... a bit.",
	"positive_points": "There is a glitch where you can actually wear a pot!",
	"negative_points": "There is a glitch where you can actually wear a pot...."
},
{
	"id": 81,
	"created_at": 1416857088212,
	"updated_at": 1416857088404,
	"username": "Tobias",
	"slug": "zelda-reviewathon-number-2",
	"url": "https://www.igdb.com/games/the-legend-of-zelda-majora-s-mask/reviews/zelda-reviewathon-number-2",
	"title": "Zelda reviewathon #2",
	"game": 1030,
	"category": 1,
	"likes": 0,
	"views": 255,
	"rating_category": 3,
	"platform": 4,
	"introduction": "In the first review I started explaining my 'Nostalgia Vortex' theory. It started by explaining that as a child you have a wonderful, irreplaceable experience on, perhaps, one of your first games. Then, after 2 or three years, as you are still bright eyes and innocent on the cusp of pubehood. SEQUEL! Bam, there goes the dynamite. Nintendo knock the ball out of the park with Majora's mask. Majoras mask was a slightly darker, slightly more advanced, slightly more Japanese version of Oot. It keeps the same graphics and many of the same characters as Oot but expands the game enough and ads enough more features for it to be  worthwhile. Or does it? Remember, at this point we are still 10, 15 year old shitlings with nothing better at the time than diablo 2. Would a sequel try to pull that shit these days and we would have a dreamcast!",
	"content": "Majoras mask starts with the actual link from the prequel looking for his fairy or some shit. And immediately at this point all the kids ask themselves 'how did he get out of Hyrule?' I know every last inch of that country and there is no exit. Eventually some Gypsy mugs him and steals his Oot and turns him into a shrub (told you it was more Japanese) forcing Link to follow him to a clock town called Clock Town. The next three days break every rule of time travel pop culture but it works anyway and makes for a lot of fun with all the side quests and collectible masks. Eventually you face off with the Gypsy only to discover he was being controlled by the mask which then proceeds to transform into things and to fight him in the moon which is about to crash into earth... oh, and the moon has a face...",
	"conclusion": "So in conclusion it's actually quite fun but, really, after a lot of weirdness and masks, you're not left with much more than the original prequel.",
	"positive_points": "There is a guy living in a toilet who you can give paper to in return for hearts.",
	"negative_points": "The moons face gives you nightmares."
},
{
	"id": 82,
	"created_at": 1416858766332,
	"updated_at": 1416858766455,
	"username": "Tobias",
	"slug": "zelda-reviewathon-number-3",
	"url": "https://www.igdb.com/games/the-legend-of-zelda-the-wind-waker/reviews/zelda-reviewathon-number-3",
	"title": "Zelda Reviewathon #3",
	"game": 1033,
	"category": 1,
	"likes": 2,
	"views": 267,
	"rating_category": 3,
	"platform": 21,
	"introduction": "A couple of years after Majoras Mask played the N64's swan song, the game cube burst onto the stage to wow us with something or other. As if retroactively listening to me, Nintendo went for a completely unique style or game which some loved, other feared and all sort of ignored. With it's cell shaded graphics and cartoon feel. I personally loved the game and wish Link would do some more nautical stuff. The trouble with this is that now, those little shitlings who opened Oot on Christmas morning all wide eyed and innocent. Or even the ones who used their paper rout money to buy Majoras mask. They are now thinking breathing gamer teenagers who are expecting all the same endorphins that Oot gave them as children. And in an age of GTA vice city, Morrowind and Splinter Cell... it's easy to see why this plucky little game didn't get the response it deserved.  ",
	"content": "Wind Wanker starts with Link being a humble boy on an island in a sparse archipelago populated by freaks, monsters and inefficient pirates captained by the only normal looking blonde female character who is clearly princess Zelda. After some fapping around you meet a boat who can talk, obviously, and you set off to rescue your sister and eventually Zelda. One thing leads to another and you sail all over the shop controlling the wind and finding booty. It's by far the most unique Zelda for consoles so far. One might even go as far as to place it in a separate category as it really doesn't feel very Zeldary at all. However link still has his pictobox, boomerang, green hat that serves no purpose on his grotesquely oversize head, so it feels enough like home to call it Zelda.",
	"conclusion": "In summary Wind Waker is a special game that doesn't really scream special in the grand scheme of things. Yes it hit the right notes and gameplay was fun, but it was also an abrupt change of tone that some gamers might be put off by.",
	"positive_points": "The sailing could get tedious but mostly I enjoyed it.",
	"negative_points": "filling out the map was like homework.\r\n"
}]
`

func TestReviewTypeIntegrity(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping test requiring communication with external server")
	}

	c := NewClient()

	r := Review{}
	typ := reflect.ValueOf(r).Type()

	err := c.validateStruct(typ, ReviewEndpoint)
	if err != nil {
		t.Error(err)
	}
}

func TestGetReview(t *testing.T) {
	ts, c := testServerString(http.StatusOK, getReviewResp)
	defer ts.Close()

	r, err := c.GetReview(1462)
	if err != nil {
		t.Error(err)
	}

	et := "Almost perfect!"
	at := r.Title
	if at != et {
		t.Errorf("Expected title '%s', got '%s'", et, at)
	}

	eURL := URL("https://www.igdb.com/games/mario-kart-8/reviews/almost-perfect")
	aURL := r.URL
	if eURL != aURL {
		t.Errorf("Expected URL '%s', got '%s'", eURL, aURL)
	}

	ev := 250
	av := r.Views
	if av != ev {
		t.Errorf("Expected view count %d, got %d", ev, av)
	}
}

func TestGetReviews(t *testing.T) {
	ts, c := testServerString(http.StatusOK, getReviewsResp)
	defer ts.Close()

	ids := []int{1571, 65}
	r, err := c.GetReviews(ids)
	if err != nil {
		t.Error(err)
	}

	el := 2
	al := len(r)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	ec := 1
	ac := r[0].Category
	if ac != ec {
		t.Errorf("Expected date category %d, got %d", ec, ac)
	}

	erc := 3
	arc := r[0].RatingCategory
	if arc != erc {
		t.Errorf("Expected rating category %d, got %d", erc, arc)
	}

	ep := 41
	ap := r[0].Platform
	if ap != ep {
		t.Errorf("Expected platform %d, got %d", ep, ap)
	}

	eu := "ZUPERFLY"
	au := r[1].Username
	if au != eu {
		t.Errorf("Expected username '%s', got '%s'", eu, au)
	}

	ePos := "-smash balls\r\n-Subspace Emissary\r\n-gameplay\r\n-custom levels"
	aPos := r[1].PositivePoints
	if aPos != ePos {
		t.Errorf("Expected positive points '%s', got '%s'", ePos, aPos)
	}

	eNeg := "-timers\r\n-Mario"
	aNeg := r[1].NegativePoints
	if aNeg != eNeg {
		t.Errorf("Expected negative points '%s', got '%s'", eNeg, aNeg)
	}
}

func TestSearchReviews(t *testing.T) {
	ts, c := testServerString(http.StatusOK, searchReviewsResp)
	defer ts.Close()

	r, err := c.SearchReviews("zelda")
	if err != nil {
		t.Error(err)
	}

	el := 3
	al := len(r)
	if al != el {
		t.Errorf("Expected length of %d, got %d", el, al)
	}

	eID := 80
	aID := r[0].ID
	if aID != eID {
		t.Errorf("Expected ID %d, got %d", eID, aID)
	}

	es := "zelda-reviewathon-number-1"
	as := r[0].Slug
	if as != es {
		t.Errorf("Expected slug '%s', got '%s'", es, as)
	}

	eIn := "In the first review I started explaining my 'Nostalgia Vortex' theory."
	aIn := r[1].Introduction
	if !strings.Contains(aIn, eIn) {
		t.Errorf("Expected Introduction to contain '%s', got '%s'", eIn, aIn)
	}

	eCont := "I know every last inch of that country and there is no exit."
	aCont := r[1].Content
	if !strings.Contains(aCont, eCont) {
		t.Errorf("Expected Content to contain '%s', got '%s'", eCont, aCont)
	}

	eConc := "but it was also an abrupt change of tone that some gamers might be put off by."
	aConc := r[2].Conclusion
	if !strings.Contains(aConc, eConc) {
		t.Errorf("Expected Conclusion to contain '%s', got '%s'", eConc, aConc)
	}

	elc := 2
	alc := r[2].Likes
	if alc != elc {
		t.Errorf("Expected Likes count %d, got %d", elc, alc)
	}
}
