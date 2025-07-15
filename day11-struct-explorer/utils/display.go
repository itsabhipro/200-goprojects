package utils

import "fmt"

func DisplayBanner() {
	fmt.Printf("Welcome to Golang Struct explorer.\n\n")
	fmt.Printf("1.\t Explore struct\n")
	fmt.Printf("2.\t Write struct into json\n")
	fmt.Printf("3.\t Load struct from file\n")
	fmt.Printf("4.\t Exit Program\n\n")
	fmt.Printf("Choose any option: ")
}
