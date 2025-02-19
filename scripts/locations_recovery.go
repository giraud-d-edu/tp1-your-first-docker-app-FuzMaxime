package groupieTracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Locations struct {
	Id        int      `json:"id"`
	Locations []string `json:"locations"`
	Dates     string   `json:"dates"`
}
type IndexStructur struct {
	Index []Locations `json:index`
}

// Location for all artist, no implemeted
func LocationsRecovery() IndexStructur {

	var structLocations IndexStructur

	url := "https://groupietrackers.herokuapp.com/api/locations"
	// Api Request
	req, _ := http.NewRequest("GET", url, nil)
	res, errorRes := http.DefaultClient.Do(req)
	if errorRes != nil {
		fmt.Println("Error Ioutil, ", errorRes)
	}

	defer res.Body.Close()                        // close API
	data, ioutilError := ioutil.ReadAll(res.Body) // read API
	if ioutilError != nil {
		fmt.Println("Error Ioutil, ", ioutilError)
	}
	// Put the data in the struct
	err := json.Unmarshal(data, &structLocations)
	if err != nil {
		fmt.Println("Error marshal: ", err)
	}

	return structLocations
}
