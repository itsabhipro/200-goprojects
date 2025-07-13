package main

import (
	"day3-unit-converter/conv"
	"day3-unit-converter/utils"
	"fmt"
)

func main() {
	isRunning := true
	utils.ClearScreen()
	for isRunning {
		utils.DisplayWelcome()
		utils.DisplayUnitOptions()
		choiceStr := utils.HandleChoice()
		switch choiceStr {
		case "1":
			temp := conv.Temperature{}
			temp.HandleTempOptions()
			continue
		case "2":
			mass := conv.Mass{}
			mass.HandleMassOptions()
			continue
		case "3":
			distance := conv.Distance{}
			distance.HandleDistanceOptions()
			continue
		case "4":
			elec := conv.Electricity{}
			elec.HandleElectricOption()
			continue
		case "e":
			isRunning = false
		default:
			fmt.Println("Invailid choice")
			continue
		}
	}

}
