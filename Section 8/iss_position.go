package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type Location struct {
	Position Position `json:"iss_position"`
	Message  string   `json:"message"`
}

type Position struct {
	Longitude string `json:"longitude"`
	Latitude  string `json:"latitude"`
}

// get Internation Space Station Position
func GetISSPosition() (*Location, error) {
	const url = "http://api.open-notify.org/iss-now.json"

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	l := Location{}
	err = json.Unmarshal(body, &l)
	if err != nil {
		return nil, err
	}

	return &l, nil

}

func main() {
	resp, err := GetISSPosition()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Longitude: %s\n", resp.Position.Longitude)
	fmt.Printf("Latitude: %s\n", resp.Position.Latitude)

}
