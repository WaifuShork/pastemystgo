package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

// https://paste.myst.rs/api/v2/data/language?name=c
// 4945060a0ddde6baee3c175af51e16b3
func main() {
	// http://api.openweathermap.org/data/2.5/weather?id=London&appid=75945ef3cc25a44a1cf4796c0d5c7397
	access := `http://api.openweathermap.org/data/2.5/weather?id=London&appid=75945ef3cc25a44a1cf4796c0d5c7397`
	response, err := http.Get(access)

	if err != nil {
		log.Fatal(err)
	}

	//log.Print(response.Body)
	
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Print(string(bytes))
}