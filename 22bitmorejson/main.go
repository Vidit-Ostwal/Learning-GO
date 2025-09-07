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
	fmt.Println("Welcome to Json in Golang")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcocourse := []course{
		{"ReactJS Bootcamp", 299, "Learningcode.in", "adfd", []string{"web-dev", "sd"}},
		{"ReactJS Bootcamp", 199, "Learningcode.in", "adfd", []string{"web-dev", "sd"}},
		{"ReactJS Bootcamp", 399, "Learningcode.in", "adfd", nil},
	}

	// Package this data in JSON

	final_json, err := json.MarshalIndent(lcocourse, "", "\t")

	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", final_json)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
                "coursename": "ReactJS Bootcamp",
                "Price": 299,
                "website": "Learningcode.in",
                "tags": [
                        "web-dev",
                        "sd"
                ]
    }
	`)

	var lcocourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON was valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("JSON was not valid")
	}

	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData)

	for k, v := range myOnlineData {
		fmt.Printf("Key is %v and value is %v and Type is : %T\n", k, v, v)
	}

}
