package main

import (
	"bytes"
	"encoding/json"
	fmt "fmt"
)

func JSONPlay1() {
	type App struct {
		Target1 struct {
			Id    string `json:"id"`
			Title string `json:"title"`
		} `json: "Target1"`
		Target2 struct {
			Test  string `json:"CurrencyIso"`
			Title string `json:"title"`
		} `json: "Target2"`
	}

	data := []byte(`
		{
			"Target1":
				{ "id":"1000", "title":"Bookish"},
			"Target2":
				{ "CurrencyISO":"10010", "title":"Bookish"}
		}
	`)

	var app App
	_ = json.Unmarshal(data, &app)
	fmt.Print("App: ", app)

	var out bytes.Buffer
	byteData, _ := json.Marshal(app.Target2)

	json.Indent(&out, byteData, "", "-")
	fmt.Println("Marshalled ", out.String())
}

func JSONPlayTest1() {
	/*
		Marshalling to a known structure is not a problem only to generic interface.
	*/

	var myJSON = `
	{
	"data": {
		  "docType":"marble",
		  "name":"marble1",
		  "color":"red",
		  "size": 1000008,
		  "owner": "Joe"
		}
	}`

	marble1 := make(map[string]marble)
	err := json.Unmarshal([]byte(myJSON), &marble1)
	if err != nil {
		fmt.Println(err)
	}

	marblesBytes, err := json.Marshal(marble1)
	if err != nil {
		fmt.Println(err)
	}

	marble2 := make(map[string]marble)
	err = json.Unmarshal(marblesBytes, &marble2)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Unmarshalled to know structure: ", marble2["data"])

}

type marble struct {
	ObjectType string `json:"docType"` //docType is used to distinguish the various types of objects in state database
	Name       string `json:"name"`    //the fieldtags are needed to keep case from bouncing around
	Color      string `json:"color"`
	Size       int    `json:"size"`
	Owner      string `json:"owner"`
}

func JSONPlayTest3() {

	/*	"JSONPlayTest3: Unmarshalling to Generic Interface however using NewDecoder/UseNumber API \n" +
		"This will ensure float-pointing representation done by the JSON package for integer/json-number \n" +
	            "at the time of internal storage is reverted to original format."*/
	var myJSON = `
	{
	"data": {
		  "docType":"marble",
		  "name":"marble1",
		  "color":"red",
		  "size": 1000008,
		  "owner": "Joe"
		}
	}`
	jsonResult := make(map[string]interface{})

	decoder := json.NewDecoder(bytes.NewBuffer([]byte(myJSON)))
	decoder.UseNumber()
	err := decoder.Decode(&jsonResult)
	if err != nil && jsonResult["data"] == nil {
		fmt.Println("Nothing found")
	} else {

		fmt.Println("Unmarshalled to Genric Structure using NewDecoder API: ", jsonResult["data"])
	}
}

func JSONPlayTest2() {
	/*"JSONPlayTest3: Unmarshalling to Generic Interface however using Unmarshall API \n" +
	"This unmarshalling causes problem when used with Generic Interface like below for integer \n" +
	"digits size >= 7. Integer value is represented as with E notation and this will fail when " +
	"unmarshalling to integer");
	*/
	var myJSON = `
	{
	"data": {
		  "docType":"marble",
		  "name":"marble1",
		  "color":"red",
		  "size": 1000008,
		  "owner": "Joe"
		}
	}`

	jsonResult := make(map[string]interface{})
	json.Unmarshal([]byte(myJSON), &jsonResult)

	if jsonResult["data"] == nil {
		fmt.Println("Nothing found")
	} else {

		fmt.Println("Unmarshalled to Generic Interface: ", jsonResult["data"])
	}
}
