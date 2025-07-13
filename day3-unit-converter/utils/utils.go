package utils

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
)

func ClearScreen() {
	cmd := exec.Command("cmd", "/C", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func FormatNumber(number float64) string {
	result := strings.Split(fmt.Sprintf("%.2f", number), ".")
	if checkHasDecimal(result) {
		return result[0]
	}
	return fmt.Sprintf("%.2f", number)
}
func checkHasDecimal(number []string) bool {
	return strings.Contains(number[1], "00")
}

func HandleInput(inputs *[]map[string]*float64) {
	reader := bufio.NewReader(os.Stdin)
	for _, input := range *inputs {
		keys := slices.Collect(maps.Keys(input))
		fmt.Printf("Enter %s: ", keys[0])
		strVal, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Could not read from screen: ", err.Error())
			continue
		}
		strVal = strings.TrimSpace(strVal)
		floatVal, err := strconv.ParseFloat(strVal, 64)
		if err != nil {
			fmt.Println("Invailid data passed: ", err.Error())
			continue
		}
		*input[keys[0]] = floatVal
	}
}

func HandleChoice() string {
	reader := bufio.NewReader(os.Stdin)
	choiceStr, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}
	choiceStr = strings.TrimSpace(choiceStr)
	if len(choiceStr) > 1 {
		fmt.Println("Please Enter only one charactor")
		return ""
	}
	return choiceStr
}
