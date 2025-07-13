package utils

import "fmt"

func DisplayResult(title string, input float64, unit string, result float64, rUnit string) {
	fmt.Printf("\n---------Result------------\n\n")
	fmt.Printf("🌡️ %s 🌡️\nYou've entered: %f%s\nConverting that into the toasty world of Fahrenheit...\n\n🎉 Result: %f%s = %f%s\n\nStay curious, stay cool (or warm 😉), and keep converting with confidence!\n\n", title, input, unit, input, unit, result, rUnit)
}

func DisplayWelcome() {
	fmt.Printf("\n🌟 Welcome to UnitEase – Your Handy Unit Converter 🌟\nConvert with confidence!Whether it's 🌡️ Temperature, ⚖️ Mass, 📏 Distance, or ⚡ Electricity units, I've got you covered.\n\n👉 Just follow the prompts and get instant, accurate conversions – no stress, no guesswork.\nLet's get converting! 🚀\n\n")
}

func DisplayUnitOptions() {
	fmt.Printf("\n\n--------Conversion Options-----------\n")
	fmt.Printf("\nTemperature: 1\nMass: 2\nDistance: 3\nElectricity: 4\n\nChoose any opens☝️: ")
}

func DisplayConversionOption(options []string) {
	fmt.Printf("\n\n-------Actions--------\n")
	for index, option := range options {
		fmt.Printf("%s: \t%d\n", option, index+1)
	}
	fmt.Printf("\nChoose actions ☝️: ")
}
