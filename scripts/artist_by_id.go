package groupieTracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Artist struct {
	Id           int      `json:"id"`
	Image        string   `json:"image"`
	Name         string   `json:"name"`
	Members      []string `json:"members"`
	CreationDate int      `json:"creationDate"`
	FirstAlbum   string   `json:"firstAlbum"`
}

func ArtistByIdRecoverry(id int) Artist {
	id++
	var structArtists Artist

	urlApi := "https://groupietrackers.herokuapp.com/api/artists/"
	idInString := strconv.Itoa(id)
	url := urlApi + idInString
	// Api Request
	req, _ := http.NewRequest("GET", url, nil)
	res, errorRes := http.DefaultClient.Do(req)
	if errorRes != nil {
		fmt.Println("Error Ioutil, ", errorRes)
	}
	defer res.Body.Close()                        // CLose api
	data, ioutilError := ioutil.ReadAll(res.Body) // Read response
	if ioutilError != nil {
		fmt.Println("Error Ioutil, ", ioutilError)
	}
	//Put the data in the struct
	err := json.Unmarshal(data, &structArtists)
	if err != nil {
		fmt.Println("Error marshal: ", err)
	}
	return structArtists
}
