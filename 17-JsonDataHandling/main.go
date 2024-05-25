package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	MovieName   string   `json:"name"`
	Rating      float32  `json:"rating"`
	Tags        []string `json:"tags,omitempty"`
	ReleaseDate string   `json:"-"`
}

func main() {

	fmt.Println("Json in GOlang")
	DecodeJSON()

}

func EncodeJSON() {
	movies := []Movie{
		{MovieName: "The Matrix", Rating: 8.7, Tags: []string{"Action", "Sci-Fi"}, ReleaseDate: "1999-03-31"},
		{MovieName: "Inception", Rating: 8.8, Tags: []string{"Action", "Adventure", "Sci-Fi"}, ReleaseDate: "2010-07-16"},
		{MovieName: "The Shawshank Redemption", Rating: 9.3, Tags: []string{"Drama"}, ReleaseDate: "1994-10-14"},
		{MovieName: "Pulp Fiction", Rating: 8.9, Tags: []string{"Crime", "Drama"}, ReleaseDate: "1994-10-14"},
		{MovieName: "The Godfather", Rating: 9.2, Tags: nil, ReleaseDate: "1972-03-24"},
	}

	jsonData, err := json.MarshalIndent(movies, "", "\t")
	checkNilErr(err)

	fmt.Println(string(jsonData))
}

func DecodeJSON() {
	jsonDataFromWeb := []byte(`
	{
		"name": "The Matrix", 
		"rating": 8.7, 
		"tags": ["Action", "Sci-Fi"]
		}
		`)

	var tempJsonData Movie
	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json Data is Valid! :)")
		json.Unmarshal(jsonDataFromWeb, &tempJsonData)
		fmt.Printf("%#v \n", tempJsonData) // # - keys, v - value
	} else {
		fmt.Println("Json Data is not Valid! :(")
	}

	// Alternate Method

	var tempJSONData2 map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &tempJSONData2)
	fmt.Printf("%#v", tempJsonData)
}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
