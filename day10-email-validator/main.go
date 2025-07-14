package main

import (
	"bufio"
	"day10-email-validator/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Enter the email: ")
	email, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	email = strings.TrimSpace(email)
	if utils.ValidateEmail(email) {
		fmt.Println(email, " is valid.")
	} else {
		fmt.Println(email, " is not valid.")
	}
}
