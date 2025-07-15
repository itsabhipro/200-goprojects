package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ScanfString(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	if prompt != "" {
		fmt.Print(prompt, ": ")
	}
	str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error from ScanfString :", err)
		return ""
	}
	str = strings.TrimSpace(str)
	return str
}
