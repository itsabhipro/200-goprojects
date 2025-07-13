package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Give your path: ")
	root, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return
	}
	root = strings.TrimSpace(root)
	walkDir(root, "")

}
func walkDir(path string, indent string) error {
	dirs, err := os.ReadDir(path)
	if err != nil {
		return err
	}

	for index, dir := range dirs {
		symbol := "├── "
		nextPrefix := indent + "│   "

		if index == len(dirs)-1 {
			symbol = "└── "
			nextPrefix = "     "
		}
		icon := "📁"
		if !dir.IsDir() {
			icon = "📄"
		}
		info, err := dir.Info()
		if err != nil {
			return err
		}
		fmt.Println(indent + symbol + icon + " " + dir.Name() + "→• Modified Date: " + info.ModTime().String() + "→• Size: " + fmt.Sprint(info.Size()))
		fmt.Println()
		if dir.IsDir() {
			err = walkDir(filepath.Join(path, dir.Name()), nextPrefix)
			if err != nil {
				return err
			}
		}

	}
	return nil
}
