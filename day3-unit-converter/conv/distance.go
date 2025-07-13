package conv

import (
	"day3-unit-converter/utils"
	"fmt"
)

type Distance struct {
	Unit float64
}

func (dis *Distance) KiloToMiles() float64 {
	return dis.Unit * 0.621371
}

func (dis *Distance) MilesToKilo() float64 {
	return dis.Unit * 1.60934
}

func (dis *Distance) MeterToFeet() float64 {
	return dis.Unit * 3.28084
}

func (dis *Distance) FeetToMeter() float64 {
	return dis.Unit * 0.3048
}

func (dis *Distance) InchesToCentimeter() float64 {
	return dis.Unit * 2.54
}

func (dis *Distance) HandleDistanceOptions() {
	options := []string{"Kilometer to Miles", "Miles to Kilometer", "Meter to Feet", "Feet to Meter", "Inches to Centimeter"}
	utils.DisplayConversionOption(options)
	subChoice := utils.HandleChoice()
	switch subChoice {
	case "1":
		inputs := []map[string]*float64{
			{"Kilometer": &dis.Unit},
		}
		utils.HandleInput(&inputs)
		if dis.Unit < 0 {
			fmt.Println("Distance coludn't negative")
			return
		}
		utils.DisplayResult("Kilometer to Miles Converter", *inputs[0]["Kilometer"], "KM", dis.KiloToMiles(), "Miles")
		fmt.Scanln()
		return
	case "2":
		inputs := []map[string]*float64{
			{"Miles": &dis.Unit},
		}
		utils.HandleInput(&inputs)
		if dis.Unit < 0 {
			fmt.Println("Distance coludn't negative")
			return
		}
		utils.DisplayResult("Miles to Kilometer Converter", *inputs[0]["Miles"], "Miles", dis.MilesToKilo(), "KM")
		fmt.Scanln()
		return
	case "3":
		inputs := []map[string]*float64{
			{"Meter": &dis.Unit},
		}
		utils.HandleInput(&inputs)
		if dis.Unit < 0 {
			fmt.Println("Distance coludn't negative")
			return
		}
		utils.DisplayResult("Meter to Feet Converter", *inputs[0]["Meter"], "M", dis.MeterToFeet(), "FT")
		fmt.Scanln()
		return
	case "4":
		inputs := []map[string]*float64{
			{"Feet": &dis.Unit},
		}
		utils.HandleInput(&inputs)
		if dis.Unit < 0 {
			fmt.Println("Distance coludn't negative")
			return
		}
		utils.DisplayResult("Feet to Meter Converter", *inputs[0]["Feet"], "FT", dis.FeetToMeter(), "M")
		fmt.Scanln()
		return
	case "5":
		inputs := []map[string]*float64{
			{"Inches": &dis.Unit},
		}
		utils.HandleInput(&inputs)
		if dis.Unit < 0 {
			fmt.Println("Distance coludn't negative")
			return
		}
		utils.DisplayResult("Inches to Centimeter Converter", *inputs[0]["Inches"], "In", dis.InchesToCentimeter(), "CM")
		fmt.Scanln()
		return
	}
}
