package groupieTracker

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// Spotify music
func LinkMusic(id int) string {
	// stRingID := strconv.Itoa(id)

	data, error := ioutil.ReadFile("../../assets/extern/linkMusic.txt")
	if error != nil {
		fmt.Println(error)
	}

	value := strings.Split(string(data), "\n")

	return value[id*2]
	// if string(value[id*2+1]) == string(stRingID) {
	// 	return value[id*2]
	// } else {
	// 	return "4cOdK2wGLETKBW3PvgPWqT"
	// }
}
