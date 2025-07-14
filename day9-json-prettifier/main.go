package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	inputContent, err := os.ReadFile("json.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	var jsonString any
	if err := json.Unmarshal(inputContent, &jsonString); err != nil {
		fmt.Println(err)
		return
	}
	prettyJson, err := json.MarshalIndent(jsonString, " ", "     ")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Your pretty json is here")
	fmt.Println(string(prettyJson))
}
