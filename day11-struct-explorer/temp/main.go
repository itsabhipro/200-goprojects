package temp

import "fmt"

func main() {
	word := "level"
	var temp string
	for i := len(word) - 1; i >= 0; i-- {
		temp = temp + string(word[i])
	}
	fmt.Println("word: ", word)
	fmt.Println("Reverse: ", temp)
}
