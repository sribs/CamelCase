package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/sribs/CamelCase/converter"
	"github.com/sribs/CamelCase/dictionaryapi"
	"github.com/sribs/CamelCase/jsonops"
)

var path = "CamelCase.json"
var api = dictionaryapi.DictionaryAPI{
	"", "", "",
}

// GetAllCamelCase : Get all words and CamelCase conversions
func GetAllCamelCase(w http.ResponseWriter, r *http.Request) {
	jsonMap := jsonops.GetFromJSON(path)
	json.NewEncoder(w).Encode(jsonMap)
}
// GetCamelCase : Get string and its CamelCase version if exist
func GetCamelCase(w http.ResponseWriter, r *http.Request) {
	jsonMap := jsonops.GetFromJSON(path)
	params := mux.Vars(r)

	if params != nil && jsonMap[params["string"]] != "" {
		json.NewEncoder(w).Encode(map[string]string{params["string"]: jsonMap[params["string"]]})
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"failed": "Non Existent key or Incomplete Parameters"})
}
// AddCamelCase : Add a new string and perform CamelCase conversion and append it to the JSON
func AddCamelCase(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	if params != nil {
		min, err := strconv.Atoi(params["min"])
		if err != nil {
			json.NewEncoder(w).Encode(map[string]string{"failed": "Error in conversion"})
			return
		}
		camelCase := converter.CamelCaseDP(api, params["string"], min)
		valueMap := map[string]string{params["string"]: camelCase}
		jsonops.AppendToJSON(path, valueMap)
		json.NewEncoder(w).Encode(valueMap)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"failed": "Error"})
}
// DeleteCamelCase : Delete from CamelCase JSON file if the string exists
func DeleteCamelCase(w http.ResponseWriter, r *http.Request) {
	jsonMap := jsonops.GetFromJSON(path)
	params := mux.Vars(r)

	if params != nil && jsonMap[params["string"]] != "" {
		json.NewEncoder(w).Encode(map[string]string{params["string"]: jsonMap[params["string"]]})
		delete(jsonMap, params["string"])
		jsonops.UpdateToJSON(path, jsonMap)
		return
	}
	json.NewEncoder(w).Encode(map[string]string{"failed": "Non Existent key"})
}

func main() {
	if len(os.Args) >= 2 {
		api.APIURL = os.Args[1]
		if len(os.Args) == 4 {
			api.AppID = os.Args[2]
			api.AppKey = os.Args[3]
		}
	} else {
		log.Fatal("Arguments Error: Insufficient Arguments Passed")
	}
	router := mux.NewRouter()
	router.HandleFunc("/camelcase", GetAllCamelCase).Methods("GET")
	router.HandleFunc("/camelcase/{string}", GetCamelCase).Methods("GET")
	router.HandleFunc("/camelcase/{string}/{min}", AddCamelCase).Methods("POST")
	router.HandleFunc("/camelcase/{string}", DeleteCamelCase).Methods("DELETE")
	log.Fatal(http.ListenAndServe("0.0.0.0:80", router))
}
