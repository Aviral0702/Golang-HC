package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON class")
	encodeJson()
}

func encodeJson() {
	newCourse := []course{
		{
			Name:     "golang",
			Price:    299,
			Platform: "udemy",
			Password: "abc123",
			Tags:     []string{"go", "programming", "development"},
		},
		{
			Name:     "reactjs",
			Price:    199,
			Platform: "udemy",
			Password: "cde123",
			Tags:     []string{"reactjs", "development"},
		},
		{
			Name:     "JavaScript",
			Price:    599,
			Platform: "udemy",
			Password: "gullyboy",
			Tags:     []string{"JS", "Web", "development"},
		},
	}
	finalJson, err := json.MarshalIndent(newCourse, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println("Final JSON is here: ", string(finalJson))
}
