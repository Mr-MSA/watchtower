package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
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
