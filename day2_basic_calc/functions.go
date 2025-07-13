package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/C", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
	return
}
func add(reader *bufio.Reader) {
	inputs := handleInput(reader)
	result := inputs[0] + inputs[1]
	handleNumber(result, "Your addition of a + b")
}
func minus(reader *bufio.Reader) {
	inputs := handleInput(reader)
	result := inputs[0] - inputs[1]
	handleNumber(result, "Your substraction of a - b")
}
func multiply(reader *bufio.Reader) {
	inputs := handleInput(reader)
	result := inputs[0] * inputs[1]
	handleNumber(result, "Your multiplication of a x b")
}
func divide(reader *bufio.Reader) {
	inputs := handleInput(reader)
	result := inputs[0] / inputs[1]
	handleNumber(result, "Your division of a / b")
}
func handleInput(reader *bufio.Reader) []float64 {
	fmt.Print("Enter first number a: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input = strings.TrimSpace(input)
	a, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic("Invailid input does not parse.-")
	}
	fmt.Print("Enter second number b: ")
	input, err = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	b, err := strconv.ParseFloat(input, 64)
	if err != nil {
		panic("Invailid input could not parse.")
	}
	return []float64{a, b}

}
func handleNumber(number float64, title string) {
	result := strings.Split(fmt.Sprintf("%.2f", number), ".")
	if checkHasDecimal(result) {
		fmt.Printf("\n\n%s  = %s\n\n", title, result[0])
		return
	}
	fmt.Printf("\n\n%s  = %.2f", title, number)
}
func checkHasDecimal(number []string) bool {
	return strings.Contains(number[1], "00")
}
func handleChoice(reader *bufio.Reader) string {
	fmt.Printf("\nYour options is +,-,*,/ \n choose any of these options\n\nPlease enter choice: ")
	choice, err := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)
	if err != nil {
		fmt.Printf("\nPlease input vailid choice\n")
		return "error"
	}
	if len(choice) > 1 {
		fmt.Printf("\n❌Please input only one charactor\n")
		return "error"
	}
	return choice

}
