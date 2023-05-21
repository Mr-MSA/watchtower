package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func MakeHttpRequest(url string, flags intelArgs, reqbody string) string {

	client := &http.Client{}

	req, err := http.NewRequest(flags.Method, url, strings.NewReader(reqbody))
	if err != nil {
		log.Fatal(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)

}
