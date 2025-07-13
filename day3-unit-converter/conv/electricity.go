package conv

import (
	"day3-unit-converter/utils"
	"fmt"
)

type Electricity struct {
	Voltage    float64
	Current    float64
	Resistance float64
	Time       float64
}

func (e *Electricity) OhmLawFindVoltage() float64 {
	return e.Current * e.Resistance
}

func (e *Electricity) OhmLawFindCurrent() float64 {
	return e.Voltage / e.Resistance
}

func (e *Electricity) OhmLawFindResistance() float64 {
	return e.Voltage / e.Current
}

func (e *Electricity) Power() float64 {
	return e.Voltage * e.Current
}

func (e *Electricity) EnergyConsumption() float64 {
	//power is in watt and time in hours
	return (e.Power() * e.Time) / 1000
}

func (e *Electricity) HandleElectricOption() {
	utils.DisplayConversionOption([]string{"Volt finder by OHM's Law", "Current finder by OHM's Law", "Resitance finder", "Power finder", "Power consumption finder"})
	subChoice := utils.HandleChoice()
	switch subChoice {
	case "1":
		inputs := []map[string]*float64{
			{"Current in AMPs": &e.Current},
			{"Resistance in OHM": &e.Resistance},
		}
		utils.HandleInput(&inputs)
		utils.DisplayResult("Voltage Finder", *inputs[0]["Current in AMPs"], "AMP", e.OhmLawFindVoltage(), "Volts")
		fmt.Scanln()
		return
	}
}
