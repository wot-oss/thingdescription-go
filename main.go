package main

import (
	"log"
	"net/url"
	"os"

	"github.com/go-json-experiment/json"
	"github.com/web-of-things-open-source/thingdescription-go/thingDescription"
)

func main() {
	content, err := os.ReadFile("./HotelRoom.jsonld")
	if err != nil {
		log.Fatal(err)
	}

	var tdtest map[string]interface{}
	json.Unmarshal(content, &tdtest)
	if err != nil {
		log.Fatal(err)
	}

	td, err := thingDescription.UnmarshalThingDescription(content)
	if err != nil {
		log.Fatal(err)
	}
	td.Base = url.URL{}
	td.Events = nil
	td.SecurityDefinitions = nil
	td.Security = nil
	td.Forms = nil

	marshalingJSON, err := json.Marshal(&td)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("./HotelRoomMarshalled.jsonld", marshalingJSON, 0666)
}
