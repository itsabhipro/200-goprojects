package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	var name string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	name = strings.TrimSpace(name)
	name = strings.ToTitle(name)
	command := exec.Command("cmd", "/C", "cls")
	command.Stdout = os.Stdout
	command.Run()

	fmt.Printf("\n\nHello, How are you? %s\n\n", name)
}
