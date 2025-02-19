package groupieTracker

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

type AllAPI struct {
	Artist        []Artists
	Relation      []string
	ValueArtist   Artist
	YoutubeSearch string
	SpotifyLink   string
	DeezerLink    string
	SpotifyMusic  string
	WikiLink      string
}

// Home Page
func Home(w http.ResponseWriter, r *http.Request) {
	// recovery of the selected value
	sortSelector := r.FormValue("select-sort")
	itemsSelector := r.FormValue("select-items")

	// recovery of the prev and next value
	scrollArtists := r.FormValue("scroll-artists")

	allData := LoadArtistApi(sortSelector, itemsSelector, scrollArtists)
	RenderTemplates(w, "index", allData)
}

// Function for details for 1 artist
func Details(w http.ResponseWriter, r *http.Request) {
	isInData := false
	var finalId int
	// Forms retrieval
	id := r.FormValue("data")
	search := r.FormValue("search")
	finalSearch := strings.ToUpper(search)
	// Retrieval of artist info via id
	if len(id) == 1 || len(id) == 2 {
		newId, _ := strconv.Atoi(id)
		newId--
		finalId = newId
		isInData = true
	} else {
		// Load Artist data
		artistsData := LoadArtistApi("", "default", "")
		// Search by  names
		for j := 0; j != len(artistsData.Artist); j++ {
			for i := 0; i != len(artistsData.Artist[j].Name); i++ {
				upperCaseArtise := strings.ToUpper(artistsData.Artist[j].Name)
				if upperCaseArtise == finalSearch {
					isInData = true
					// Recupertion of id
					finalId = artistsData.Artist[j].Id - 1
				}
			}
		}
	} // Case of error in the search
	if !isInData {
		allData := AllAPI{}
		RenderTemplates(w, "error", allData)
	} else {
		// Send data in struct
		relationData := RelationRecovery(finalId)
		Value := ArtistByIdRecoverry(finalId)
		spotifyLink := LinkMusic(finalId)
		// Data formatting for link
		allLink := ArtistSocialLink(finalId)
		// send data to html page
		allData := AllAPI{ValueArtist: Value, Relation: relationData, YoutubeSearch: allLink[0], SpotifyLink: allLink[1], DeezerLink: allLink[2], SpotifyMusic: spotifyLink, WikiLink: allLink[3]}
		RenderTemplates(w, "details", allData)
	}
}

// About page
func About(w http.ResponseWriter, r *http.Request) {
	allData := AllAPI{}
	RenderTemplates(w, "about", allData)
}

// Execution of server
func RenderTemplates(w http.ResponseWriter, tmpl string, allData AllAPI) {
	t, err := template.ParseFiles("../../assets/templates/" + tmpl + ".html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	t.Execute(w, allData)
}

// Return socail link for artist
func ArtistSocialLink(id int) []string {
	var artistNamesWIthoutSpaces []string
	var artistNamesWIthoutSpacesForWiki []string
	var spotifyANdDezerrLink []string
	var linkArray []string
	Value := ArtistByIdRecoverry(id)
	// Data formatting for link
	artistName := strings.ToLower(Value.Name)
	// Change space into an other caracter for different link
	for v := 0; v < len(artistName); v++ {
		artistNamesWIthoutSpaces = append(artistNamesWIthoutSpaces, string(artistName[v]))
		artistNamesWIthoutSpacesForWiki = append(artistNamesWIthoutSpacesForWiki, string(artistName[v]))
		spotifyANdDezerrLink = append(spotifyANdDezerrLink, string(artistName[v]))
		if artistNamesWIthoutSpaces[v] == " " {
			artistNamesWIthoutSpaces[v] = "+"
		}
		if artistNamesWIthoutSpacesForWiki[v] == " " {
			artistNamesWIthoutSpacesForWiki[v] = "_"
		}
		if spotifyANdDezerrLink[v] == " " {
			spotifyANdDezerrLink[v] = "%20"
		}

	}
	// Formatting the string for ling
	for i := 0; i < len(artistNamesWIthoutSpacesForWiki); i++ {
		artistNamesWIthoutSpacesForWiki[0] = strings.ToUpper(artistNamesWIthoutSpacesForWiki[0])
		if artistNamesWIthoutSpacesForWiki[i] == "_" {
			artistNamesWIthoutSpacesForWiki[i+1] = strings.ToUpper(artistNamesWIthoutSpacesForWiki[i+1])
		}
	}
	// Create final link
	artistName = strings.Join(artistNamesWIthoutSpaces, "")
	artistNameWiki := strings.Join(artistNamesWIthoutSpacesForWiki, "")
	sportifyAndDezerString := strings.Join(spotifyANdDezerrLink, "")
	linkWinki := "https://fr.wikipedia.org/wiki/" + artistNameWiki
	link := "https://www.youtube.com/results?search_query=" + artistName
	spotifyLink := "https://open.spotify.com/search/" + sportifyAndDezerString
	deezerLink := "https://www.deezer.com/search/" + sportifyAndDezerString
	linkArray = append(linkArray, link, spotifyLink, deezerLink, linkWinki)
	return linkArray
}
