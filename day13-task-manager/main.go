package main

import (
	"day13-task-manager/store"
	"day13-task-manager/utils"
	"fmt"
)

func main() {
	//utils.DisplayOption()
	var task = store.Task{}
	isRunning := true

	for isRunning {
		utils.DisplayOption()
		choice := utils.HandleChoice("Enter your choice: ")
		if len(choice) > 1 {
			fmt.Println("Please enter only 1 charactor")
			continue
		}
		switch choice {
		case "1":
			err := utils.HandleTaskInput(&task)
			if err != nil {
				fmt.Println("Error: ", err.Error())
				continue
			}
			task.Save()
			fmt.Println("Task has been created.")
			continue
		case "2":
			store.DisplayAllTasks()
			continue
		case "5":
			isRunning = false
			return
		}

	}
}
