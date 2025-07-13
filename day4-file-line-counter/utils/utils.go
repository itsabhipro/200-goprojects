package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func CountLineWordsFromFile(path string) {
	f, err := os.Open(path)
	if os.IsNotExist(err) {
		fmt.Println("File doesn't exists.")
		return
	}
	defer f.Close()
	count := 0
	numberOfWords := 0
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		if reader.Err() == io.EOF {
			fmt.Println("heats ending of line")
			return
		}
		if reader.Err() != nil {
			fmt.Println(reader.Err())
			return
		}
		text := strings.TrimSpace(reader.Text())
		numberOfWords = numberOfWords + len(strings.Split(text, " "))
		count = count + 1
	}
	fmt.Println("---------- Result ------------")
	fmt.Println()
	fmt.Println("Number of lines in files is ", count, " lines")
	fmt.Println("Total numbers of words in file is ", numberOfWords, " words.")
	fmt.Println()
	fmt.Println("------------------------------")

}
