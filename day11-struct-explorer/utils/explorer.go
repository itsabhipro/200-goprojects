package utils

import (
	"bufio"
	"fmt"
	"os"
)

func WriteContentToFile(content string) error {
	cFile, err := os.Create("store.txt")
	if err != nil {
		return err
	}
	defer cFile.Close()
	writer := bufio.NewWriter(cFile)
	writtenSize, err := writer.WriteString(content)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	fmt.Println("File has been written of size: ", writtenSize)
	return nil
}
