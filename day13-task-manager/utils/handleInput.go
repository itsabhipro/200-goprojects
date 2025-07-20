package utils

import (
	"bufio"
	"day13-task-manager/store"
	"fmt"
	"os"
	"reflect"
	"strings"
	"time"
)

func HandleTaskInput(task *store.Task) error {
	reader := bufio.NewReader(os.Stdin)

	taskFields := reflect.TypeOf(*task)
	taskValues := reflect.ValueOf(task).Elem()

	for i := 0; i < taskFields.NumField(); i++ {
		if taskFields.Field(i).Name == "taskId" || taskFields.Field(i).Name == "createdDate" || taskFields.Field(i).Name == "status" {
			continue
		}

		fmt.Printf("Enter %s: ", taskFields.Field(i).Name)
		input, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		input = strings.TrimSpace(input)
		if taskFields.Field(i).Name == "TargetDate" {
			layout := "2006-01-02"
			date, err := time.Parse(layout, input)
			if err != nil {
				return err
			}
			field := taskValues.FieldByName(taskFields.Field(i).Name)
			field.Set(reflect.ValueOf(date))
			continue
		}
		field := taskValues.FieldByName(taskFields.Field(i).Name)
		field.Set(reflect.ValueOf(input))
	}
	return nil
}

func HandleChoice(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	choice, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println(err)
		return ""
	} else {
		choice = strings.TrimSpace(choice)
		return choice
	}
}
