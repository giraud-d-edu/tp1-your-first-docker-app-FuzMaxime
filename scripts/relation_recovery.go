package groupieTracker

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Relation struct {
	ID             int                 `json:"id"`
	DatesLocations map[string][]string `json:"datesLocations"`
}

// Relation API for 1 artist
func RelationRecovery(indexValue int) []string {
	indexValue ++
	var relationData []string
	var dateArray []string
	var locationArray []string

	// --------------- Pull API by URL ---------------
	intValue := strconv.Itoa(indexValue)
	url := "https://groupietrackers.herokuapp.com/api/relation/"
	finalUrl := url + intValue
	resp, errRespons := http.Get(finalUrl)
	if errRespons != nil {
		ErrManagement(errRespons)
	}
	// --------------- Close API ---------------
	defer resp.Body.Close()

	// --------------- Read API ---------------
	body, errIoutils := ioutil.ReadAll(resp.Body)
	if errIoutils != nil {
		ErrManagement(errIoutils)
	}

	// --------------- Data into struct ---------------
	var relation Relation
	errMarshal := json.Unmarshal(body, &relation)
	if errMarshal != nil {
		ErrManagement(errMarshal)
	}

	// --------------- Sorting the data ---------------
	for location, dates := range relation.DatesLocations {
		relationData = append(relationData, dates[0], location)
	}
	//  --------------- Separation of dates and location ---------------
	for i := 0; i < len(relationData); i++ {
		if i%2 == 0 {
			dateArray = append(dateArray, relationData[i])
			}else{
				locationArray = append(locationArray, relationData[i])
			}
		}
		for i := 0 ; i < len(dateArray); i++ {
			for v := i +1 ; v < len(locationArray) ; v++ {
				// --------------- Verification of the year ---------------
				if dateArray[v][6:len(dateArray[v])] < dateArray[i][6:len(dateArray[i])] {  
					dateArray[i] , dateArray[v] = dateArray[v] , dateArray[i]
					locationArray[i], locationArray[v] = locationArray[v] , locationArray[i]
				}
				// --------------- Verification of the month ---------------
				if dateArray[v][3:5] < dateArray[i][3:5] && dateArray[v][6:len(dateArray[v])] == dateArray[i][6:len(dateArray[i])]{  // trie par mois si l'année est pareil
						dateArray[i] , dateArray[v] = dateArray[v] , dateArray[i]
						locationArray[i], locationArray[v] = locationArray[v] , locationArray[i]
				}// --------------- Verification of the days ---------------
				if dateArray[i][0:2]> dateArray[v][0:2] && dateArray[v][3:5] == dateArray[i][3:5] && dateArray[v][6:len(dateArray[v])] == dateArray[i][6:len(dateArray[i])] {	// trie par jour si mois et année pareil
						dateArray[i] , dateArray[v] = dateArray[v] , dateArray[i]
						locationArray[i], locationArray[v] = locationArray[v] , locationArray[i]			
					}
				}		
			}			
		var finalArrayByDates []string
		for i := 0 ; i < len(locationArray); i++{
			finalArrayByDates = append(finalArrayByDates, locationArray[i], dateArray[i])
		}	
	return finalArrayByDates						
}			
					
func ErrManagement(err error) error {
	fmt.Println(err)
	return err
}
