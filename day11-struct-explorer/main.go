package main

import (
	"day11-struct-explorer/explorer/struc"
	"day11-struct-explorer/utils"
	"fmt"
	"strings"
)

type User struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Password string
}
type Test struct {
	Name    string   `json:"name"`
	Emails  []string `json:"email"`
	Authors []User   `json:"authors"`
}

func main() {
	test := Test{}
	isRunning := true

	for isRunning {
		utils.DisplayBanner()
		selection := utils.ScanfString("")
		if !utils.ValidateSelection(selection) {
			fmt.Println("Invailid selection, Please select valid.")
			continue
		}
		switch selection {
		case "1":
			inspectedContent := struc.InsepectStruct(test)
			fmt.Printf("-------- Your inspected Struct is here ---------\n\n")
			fmt.Println(inspectedContent)
			choice := utils.ScanfString("Do you want to write inspected struct into file (y-yes/anykey-No)")
			if strings.ToLower(choice) == "y" {
				err := utils.WriteContentToFile(inspectedContent)
				if err != nil {
					fmt.Println("Error to writing into file: ", err)
					continue
				}
			}
			continue
		case "2":
			confirm := utils.ScanfString("Do you want write into json file? (y-yes/press any key)")
			if strings.ToLower(confirm) == "y" {
				confirm := utils.ScanfString("Do you want to write pretty json.(y-yes)")
				if strings.ToLower(confirm) == "y" {
					confirm := utils.ScanfString("Do you want to write seperate file.(y-yes)")
					if strings.ToLower(confirm) == "y" {
						filename := utils.ScanfString("Enter your filename: ")
						err := struc.WritePrettyJson(test, filename)
						if err != nil {
							fmt.Println("Error to write json: ", err)
							continue
						}
					}

				}
				err := struc.WriteStructIntoJson(test)
				if err != nil {
					fmt.Println("Error from json: ", err)
					continue
				}
			}
			continue
		case "3":
		case "4":
			isRunning = false
			return
		}
	}

}
