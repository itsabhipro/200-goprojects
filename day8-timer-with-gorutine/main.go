package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"golang.org/x/term"
)

func clearScreen() {
	cmd := exec.Command("cmd", "/C", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()

}

func startTime(isRunning chan bool, action chan rune) {
	clearScreen()
	fmt.Println("Welcome to my timer app")
	fmt.Println("Press s for Stop/Exit and p for pause:")
	running := true
	isPaused := false
	hour := 0
	minute := 0
	second := 0
	for running {
		select {
		case key := <-action:
			if key == 's' {
				running = false
				fmt.Println()
				fmt.Println("Bye bye:)")
				isRunning <- false
				return
			}
		case key := <-action:
			if key == 'p' {
				isPaused = !isPaused
			}

		default:
			//clearScreen()
			if second == 60 {
				minute = minute + 1
				second = 0
			} else if minute == 60 {
				hour = hour + 1
				minute = 0
			} else {
				if !isPaused {
					second = second + 1
				}

			}
			fmt.Printf("\r%02d:%02d:%02d", hour, minute, second)
			time.Sleep(500 * time.Millisecond)
		}

	}
	//isRunning <- false
}

func main() {
	ch := make(chan bool)
	inputText := make(chan rune, 1)
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		log.Println(err)
		return
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)
	go func() {
		fmt.Println()
		for {
			var b = make([]byte, 1)
			os.Stdin.Read(b)
			inputText <- rune(b[0])
			//reader.Reset(reader)
		}
	}()
	go startTime(ch, inputText)
	<-ch
}
