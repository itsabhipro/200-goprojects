package main

import (
	"bufio"
	"day6-safe-file-delete/utils"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter directory path: ")
	reader := bufio.NewReader(os.Stdin)
	fileDirectory, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fileDirectory = strings.TrimSpace(fileDirectory)
	fmt.Print("Enter your file name: ")
	fileName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	fileName = strings.TrimSpace(fileName)
	path, err := utils.GetFilePath(fileDirectory, fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	file, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = utils.MoveToTrash(file)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("file has been deleted")
	defer file.Close()
}
