package main

import (
	"day11-struct-explorer/explorer/struc"
	"day11-struct-explorer/utils"
	"fmt"
	"strings"
)

func main() {
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
			name := utils.ScanfString("Enter struct name")
			if name == "" || name == " " {
				fmt.Println("Please input valid name")
				continue
			}
			gStruct, err := structRegistry[name]
			if !err {
				fmt.Println("Struct not Found")
				continue
			}
			inspectedContent := struc.InsepectStruct(gStruct)
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
			name := utils.ScanfString("Enter struct name")
			test, ok := structRegistry[name]
			if !ok {
				fmt.Println("Struct not found.")
				continue
			}
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
			struc.ExploreFile()
		case "4":
			isRunning = false
			return
		}
	}

}
