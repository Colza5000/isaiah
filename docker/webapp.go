package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// type WorldCupTeams struct {
// 	Count int `json:"count"`
// 	Teams []struct {
// 		CrestURL         string      `json:"crestUrl"`
// 		ID               int         `json:"id"`
// 		Name             string      `json:"name"`
// 		ShortName        interface{} `json:"shortName"`
// 		SquadMarketValue interface{} `json:"squadMarketValue"`
// 	} `json:"teams"`
// }

// func main() {
// 	resp, err := http.Get("http://api.football-data.org/v1/competitions/467/teams")
// 	defer resp.Body.Close()
// 	if err != nil {
// 		println(err)
// 	}
//
// 	body, err := ioutil.ReadAll(resp.Body)
//
// 	if err != nil {
// 		println(err)
// 	}
//
// 	var football WorldCupTeams
// 	err = json.Unmarshal(body, &football)
// 	if err != nil {
// 		fmt.Printf("%s", err)
// 	}
//
// 	for _, v := range football {
// 		fmt.Printf("%s ", v.Teams)
// 	}
// }

// func main() {
func Teams(w http.ResponseWriter, r *http.Request) {
	url := "http://api.football-data.org/v1/competitions/467/teams"
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	responseString := string(responseData)

	// fmt.Println(responseString)
	fmt.Fprint(w, responseString)
}

func main() {
	http.HandleFunc("/", Teams)
	http.ListenAndServe(":8080", nil)
}
