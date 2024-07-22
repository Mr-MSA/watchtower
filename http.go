package main

import (
	"encoding/base64"
	"encoding/json"
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

	// set content-type
	if reqbody != "" {
		if isJSON(reqbody) {
			req.Header.Set("Content-Type", "application/json")
		} else {
			req.Header.Set("Content-Type", "text/plain")
		}
	}

	// set credentials
	creds := base64.StdEncoding.EncodeToString([]byte(envVariable("Username") + ":" + envVariable("Password")))
	req.Header.Set("Authorization", "basic "+creds)

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

func isJSON(s string) bool {
	var js interface{}
	err := json.Unmarshal([]byte(s), &js)
	return err == nil
}
