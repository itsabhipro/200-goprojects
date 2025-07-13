package main

import (
	"bufio"
	"day4-file-line-counter/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter the path: ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	text = strings.TrimSpace(text)
	utils.CountLineWordsFromFile(text)
}
