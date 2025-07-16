package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter word: ")
	word, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error from input: ", err)
		return
	}
	word = strings.TrimSpace(word)
	temp := ""
	for i := len(word) - 1; i >= 0; i-- {
		temp = temp + string(word[i])
	}
	if strings.EqualFold(word, temp) {
		fmt.Println("Yes, this word is palindrome")
	} else {
		fmt.Println("No, This word is not palindrome")
	}

}
