package conv

import (
	"day3-unit-converter/utils"
	"fmt"
	"math/big"
)

type Temperature struct {
	Unit float64
}

func (temp *Temperature) CToF() float64 {
	frac, _ := big.NewRat(9, 5).Float64()
	return frac*temp.Unit + 32
}
func (temp *Temperature) FToC() float64 {
	frac, _ := big.NewRat(9, 5).Float64()
	return frac * (temp.Unit - 32)
}
func (temp *Temperature) CelToKel() float64 {
	return temp.Unit + 273.15
}
func (temp *Temperature) KelToCel() float64 {
	return temp.Unit - 273.15
}
func (temp *Temperature) FahToKel() float64 {
	return temp.FToC() + 273.15
}

func (temp *Temperature) HandleTempOptions() {
	utils.DisplayConversionOption([]string{"Celciuse to Fariegnhight", "Fariegnhight to celcius", "Celcius to Kelvin", "Kelvin to Celcius", "Fahrenheit to Kelvin"})
	subChoice := utils.HandleChoice()
	switch subChoice {
	case "1":
		inputs := []map[string]*float64{
			{"Celciuse": &temp.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Celsius to Fahrenheit Converter", *inputs[0]["Celciuse"], "°C", temp.CToF(), "°F")
		fmt.Scanln()
		return
	case "2":
		inputs := []map[string]*float64{
			{"Fahrenheit": &temp.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Fahrenheit to Celcius Converter", *inputs[0]["Fahrenheit"], "°F", temp.FToC(), "°C")
		fmt.Scanln()
		return
	case "3":
		inputs := []map[string]*float64{
			{"Celcius": &temp.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Celcius to Kelvin Converter", *inputs[0]["Celcius"], "°C", temp.CelToKel(), "K")
		fmt.Scanln()
		return
	case "4":
		inputs := []map[string]*float64{
			{"Kelvin": &temp.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Kelvin to celcius Converter", *inputs[0]["Kelvin"], "K", temp.KelToCel(), "°C")
		fmt.Scanln()
		return
	case "5":
		inputs := []map[string]*float64{
			{"Fahrenheit": &temp.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Fahrenheit to Kelvin Converter", *inputs[0]["Fahrenheit"], "°F", temp.FahToKel(), "K")
		fmt.Scanln()
		return
	}
}
