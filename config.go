package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func ReadJSON(filename string) map[string]interface{} {

	fileHandler, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening structure")
		os.Exit(2)
	}

	fileData, _ := ioutil.ReadAll(fileHandler)

	var out map[string]interface{}
	json.Unmarshal(fileData, &out)

	return out
}

func envVariable(key string) string {
	// read and parse configurations
	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatal(err)
	}

	err = godotenv.Load(homedir + "/.watch-client/.env")
	if err != nil {
		fmt.Printf("Error: %s", err)
		os.Exit(2)
	}

	val := os.Getenv(key)
	return val
}
func downloadFile(filepath string, fileurl string) (err error) {

	// Create blank file
	file, err := os.Create(filepath)
	if err != nil {
		return err
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fileurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	defer file.Close()

	return nil
}
