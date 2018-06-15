package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func Teams(w http.ResponseWriter, r *http.Request) {
	url := "http://api.football-data.org/v1/competitions/467/fixtures?matchday=1"
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

	fmt.Fprint(w, responseString)
}

func main() {
	http.HandleFunc("/", Teams)
	http.ListenAndServe(":8080", nil)
}
