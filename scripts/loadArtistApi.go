package groupieTracker

import (
	"strconv"
)

var sort string = ""
var items string = ""

var firstItem int = 0
var lastItem int = 0

// Load Api with sort specification.
func LoadArtistApi(sortchoice string, itemschoice string, scrollArtists string) AllAPI {

	var data = AllAPI{Artist: artistsRecovery()}
	parserData := AllAPI{}

	if sortchoice != "" {
		sort = sortchoice
	} else if sortchoice == "default" {
		sort = ""
	}

	if itemschoice != "" {
		items = itemschoice
		firstItem = 0
		lastItem = 0
	} else if itemschoice == "default" {
		items = ""
	}

	// function to sort
	Sorts(sort, data.Artist, 0, len(data.Artist)-1)

	if (items == "4") || (items == "13") || (items == "26") {
		// convert string to int + error handling
		nbItems, _ := strconv.Atoi(items)

		lastItem += nbItems
		if lastItem > nbItems {
			lastItem -= nbItems
		}

		if (scrollArtists == "prev") && (firstItem > 0) && (lastItem > nbItems) {
			firstItem = firstItem - nbItems
			lastItem = lastItem - nbItems

		} else if (scrollArtists == "next") && (lastItem < 52) {
			firstItem = firstItem + nbItems
			lastItem = lastItem + nbItems
		}

		// adds the indexed content
		for i := firstItem; i < lastItem; i++ {
			parserData.Artist = append(parserData.Artist, data.Artist[i])
		}
		return parserData

	}
	return data
}
