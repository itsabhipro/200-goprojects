package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/yuin/goldmark"
)

func main() {
	mdFile, err := os.ReadFile("example.md")
	if err != nil {
		fmt.Println(err)
		return
	}
	outputFile, err := os.Create("outputHtml.html")
	if err != nil {
		log.Println(err)
		return
	}
	outputHtml := bufio.NewWriter(outputFile)
	md := goldmark.New()
	err = md.Convert(mdFile, outputHtml)
	if err != nil {
		log.Println(err)
		return
	}

}
