package main

import (
	"encoding/json"
	"github.com/totalsecond/quotesapi/handler"
	"io/ioutil"
	"log"
	"os"
)

func dbLoad() {
	jsonFile, err := os.Open("./quotes.json")
	if err != nil {
		log.Fatalln(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &handler.quotes)
}

func dbSave() {
	struct2json, _ := json.Marshal(&quotes)
	err := ioutil.WriteFile("quotes.json", struct2json, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
