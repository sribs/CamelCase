package converter

import (
	"strings"

	"github.com/sribs/CamelCase/dictionaryapi"
)

//ConvertCamelCase : Function to convert string to CamelCase String using Dictionary API
func ConvertCamelCase(d dictionaryapi.DictionaryAPI, word string) string {
	var wordSize = len(word)
	if wordSize == 0 {
		return ""
	}
	camelCase := ""
	dpBool := make([]bool, wordSize+1)
	var matchedIndex []int
	matchedIndex = append(matchedIndex, -1)

	for index := 0; index < wordSize; index++ {
		msize := len(matchedIndex)
		flag := false
		for jndex := msize - 1; jndex >= 0; jndex-- {
			substr := word[matchedIndex[jndex]+1 : matchedIndex[jndex]+1+index-matchedIndex[jndex]]
			if d.IsAWord(substr) {
				flag = true
				print(camelCase)
				camelCase += strings.Title(substr)
				break
			}
		}
		if flag {
			println("Entered")
			dpBool[index] = true
			matchedIndex = append(matchedIndex, index)
			println(len(matchedIndex))
		}
	}
	return camelCase
}

var cache = make(map[string][]string)

// CamelCaseDP : A DP based function to convert to CamelCase
func CamelCaseDP(d dictionaryapi.DictionaryAPI, word string, minLength int) string {
	list := camelCaseDPUtil(d, word, minLength)
	var camelCase string
	for index := 0; index < len(list); index++ {
		camelCase += strings.Title(list[index])
	}
	return camelCase
}
func camelCaseDPUtil(d dictionaryapi.DictionaryAPI, word string, minLength int) []string {
	if val, ok := cache[word]; ok && len(word) == len(val) {
		return val
	}
	if word == "" {
		return []string{}
	}
	for length := minLength; length <= len(word); length++ {
		if d.IsAWord(word[:length]) {
			remaining := camelCaseDPUtil(d, word[length:], minLength)
			if remaining != nil {
				result := append([]string{word[:length]}, remaining...)
				cache[word] = result
				return result
			}
		}
	}
	cache[word] = nil
	return nil
}
