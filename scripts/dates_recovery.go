package groupieTracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type IdenxStructure struct {
	Index    []Dates    `json:"index"`
}

type Dates struct {
	Id int `json:"id"`
	Dates []string `json:"dates"`
} 

func DatesRecovery() IdenxStructure {

	var structDates IdenxStructure

	url := "https://groupietrackers.herokuapp.com/api/dates"
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
	err := json.Unmarshal(data, &structDates)
	if err != nil {
		fmt.Println("Error marshal: ", err)
	}

	return structDates
}
