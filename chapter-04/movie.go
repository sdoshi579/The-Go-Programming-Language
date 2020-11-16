package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string
	Year   int  `json:"released"`
	Color  bool `json:"color,omitempty"`
	Actors []string
}

func main() {
	var movies = []Movie{
		{
			Title:  "Don",
			Year:   1982,
			Color:  false,
			Actors: []string{"Amitabh", "Hema", "Jaya"},
		},
		{
			Title:  "Don 2",
			Year:   2012,
			Color:  true,
			Actors: []string{"Sharukh", "Priyanka", "Boman"},
		},
	}

	data, _ := json.Marshal(movies)
	fmt.Printf("%s\n", data)

	data, _ = json.MarshalIndent(movies, " ", "     ")
	fmt.Printf("%s\n", data)
}
