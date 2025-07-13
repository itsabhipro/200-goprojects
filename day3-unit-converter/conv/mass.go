package conv

import (
	"day3-unit-converter/utils"
	"fmt"
)

type Mass struct {
	Unit float64
}

func (m *Mass) KilogramToPounds() float64 {
	return m.Unit * 2.20462
}

func (m *Mass) PoundsToKilogram() float64 {
	return m.Unit * 0.453592
}

func (m *Mass) GramsToOunces() float64 {
	return m.Unit * 0.035274
}

func (m *Mass) OuncesToGrams() float64 {
	return m.Unit * 28.3495
}

func (m *Mass) TonnesToKilogram() float64 {
	return m.Unit * 1000
}

func (m *Mass) HandleMassOptions() {
	options := []string{"Kilogram to Pounds", "Pounds to Kilogram", "Grams to Ounces", "Ounces to Grams", "Tonnes to Kilogram"}
	utils.DisplayConversionOption(options)
	subChoice := utils.HandleChoice()
	switch subChoice {
	case "1":
		inputs := []map[string]*float64{
			{"Kilogram": &m.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Kilogram to Pounds Converter", *inputs[0]["Kilogram"], "KG", m.KilogramToPounds(), "LB")
		fmt.Scanln()
		return
	case "2":
		inputs := []map[string]*float64{
			{"Pound": &m.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Pounds to Kilogram Converter", *inputs[0]["Pound"], "LB", m.PoundsToKilogram(), "KG")
		fmt.Scanln()
		return
	case "3":
		inputs := []map[string]*float64{
			{"Grams": &m.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Grams to Ounces Converter", *inputs[0]["Grams"], "G", m.GramsToOunces(), "OZ")
		fmt.Scanln()
		return
	case "4":
		inputs := []map[string]*float64{
			{"Ounces": &m.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Ounces to Grams Converter", *inputs[0]["Ounces"], "OZ", m.OuncesToGrams(), "G")
		fmt.Scanln()
		return
	case "5":
		inputs := []map[string]*float64{
			{"Tonnes": &m.Unit},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Ton to Kilogram Converter", *inputs[0]["Tonnes"], "TON", m.TonnesToKilogram(), "KG")
		fmt.Scanln()
		return
	}
}
