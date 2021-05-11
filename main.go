package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

//var access = getLanguage("csharp")

func main() {
	response, err := http.Get("access")

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