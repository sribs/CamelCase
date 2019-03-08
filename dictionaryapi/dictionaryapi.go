package dictionaryapi

import (
	"log"
	"net/http"
)

// DictionaryAPI : API Attributes for Storing AppID, AppKey and APIURL
type DictionaryAPI struct {
	AppID  string
	AppKey string
	APIURL string
}

func statusCodeEval(statusCode int) bool {
	if statusCode >= 400 {
		return false
	}
	return true
}

// IsAWord : A function which queries the API to check if a word exists in dictionary
func (d *DictionaryAPI) IsAWord(word string) bool {

	if d.AppID == "" && d.AppKey == "" {
		response, err := http.Get(d.APIURL + word)
		if err != nil {
			log.Fatal(err)
		}
		return statusCodeEval(response.StatusCode)
	}
	client := &http.Client{}

	request, err := http.NewRequest("GET", d.APIURL+word, nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Add("app_id", d.AppID)
	request.Header.Add("app_key", d.AppKey)

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	return statusCodeEval(response.StatusCode)
}
