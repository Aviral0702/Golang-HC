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
	// encodeJson()
	decodeJSON()
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

func decodeJSON() {
	jsonFromWeb := []byte(`
		{
			"coursename": "golang",
			"Price": 299,
			"website": "udemy",
			"tags": ["go", "programming", "development"]
		}
	`)
	var mycourse course
	jsonValid := json.Valid(jsonFromWeb)

	if jsonValid {
		fmt.Println("Json was valid")
		json.Unmarshal(jsonFromWeb, &mycourse)
		fmt.Printf("%+v\n", mycourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}

	json.Unmarshal(jsonFromWeb, &myOnlineData)
	fmt.Printf("The type of data is %#v\n", &myOnlineData)

	for key, value := range myOnlineData {
		fmt.Printf("Key: %v, Value: %v and type is: %T\n", key, value, value)
	}
}
