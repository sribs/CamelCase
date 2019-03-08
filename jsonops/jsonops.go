package jsonops

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
)

// GetFromJSON : Function to Get from JSON file
func GetFromJSON(path string) map[string]string {
	jsonFile, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer jsonFile.Close()
	byteArr, _ := ioutil.ReadAll(jsonFile)
	jsonMap := make(map[string]string)
	json.Unmarshal(byteArr, &jsonMap)
	return jsonMap
}

// UpdateToJSON : Function to update JSON contents
func UpdateToJSON(path string, jsonMap map[string]string) {
	mappedJSON, _ := json.Marshal(jsonMap)
	err := ioutil.WriteFile(path, mappedJSON, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// AppendToJSON : function to append new json to the original json file
func AppendToJSON(path string, camelMap map[string]string) {
	jsonMap := GetFromJSON(path)
	for key, value := range camelMap {
		jsonMap[key] = value
	}
	UpdateToJSON(path, jsonMap)
}
