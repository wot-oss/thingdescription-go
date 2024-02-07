package main

import (
	"encoding/json"
	"log"
	"net/url"
	"os"

	"code.siemens.com/saywot/go-thingdescription-structs/thingDescription"
)


func main() { 
	content, err := os.ReadFile("./HotelRoom.jsonld")
	if err != nil {
		log.Fatal(err)
	}
	td,err := thingDescription.UnmarshalThingDescription(content)
	if err != nil {
		log.Fatal(err)
	}
	td.Base = url.URL{}
	td.Events = nil
	td.SecurityDefinitions = nil
	td.Security = nil
	td.Forms = nil

	marshalingJSON ,err := json.Marshal(&td)
	if err != nil {
		log.Fatal(err)
	}
	os.WriteFile("./HotelRoomMarshalled.jsonld", marshalingJSON, 0666)
}
