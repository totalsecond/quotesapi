package db

import (
	"encoding/json"
	"github.com/totalsecond/quotesapi/handler"
	"io/ioutil"
	"log"
	"os"
)

// Load loads the specified json database file
func Load(file string) {
	jsonFile, err := os.Open("./quotes.json")
	if err != nil {
		log.Fatalln(err)
	}

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &handler.Quotez)
}

// Save converts the in-memory struct to json and saves it to our database file
func Save(file string) {
	struct2json, _ := json.Marshal(&handler.Quotez)
	err := ioutil.WriteFile(file, struct2json, 0644)
	if err != nil {
		log.Fatalln(err)
	}
}
