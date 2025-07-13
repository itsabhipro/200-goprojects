package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	isRunning := true
	for isRunning {
		fmt.Printf("\nWelcome to our basic calculator\n\n")
		choice := strings.TrimSpace(handleChoice(reader))
		switch choice {
		case "+":
			add(reader)
			reader.ReadString('\n')
			clearScreen()
			continue
		case "-":
			minus(reader)
			reader.ReadString('\n')
			clearScreen()
			continue
		case "*":
			multiply(reader)
			reader.ReadString('\n')
			clearScreen()
			continue
		case "/":
			divide(reader)
			reader.ReadString('\n')
			clearScreen()
			continue
		case "e":
			isRunning = false
		default:
			fmt.Printf("\nPlease choose one of above choices\n")
			reader.ReadString('\n')
			//fmt.Println("press any key to continue....")
			clearScreen()
			continue
		}
	}
}
