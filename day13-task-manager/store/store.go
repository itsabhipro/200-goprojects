package store

import (
	"fmt"
	"time"
)

type Task struct {
	taskId      int64
	Title       string
	Description string
	TargetDate  time.Time
	createdDate time.Time
	status      string
}

type tasks struct {
	count int64
	Tasks []Task
}

var store = tasks{
	count: 0,
}

func (t *Task) Save() *tasks {
	t.taskId = store.count + 1
	t.createdDate = time.Now()
	t.status = "Pending"
	store.Tasks = append(store.Tasks, *t)
	return &store
}

func (t *Task) SetStatus(status string) error {
	if status == "Pending" || status == "Completed" {
		t.status = status
		return nil
	}
	return fmt.Errorf("status should only be Pending or Completed")
}

func (t *Task) DisplayTask() {
	fmt.Printf("%-10s %-20s %-30s %-20s %-20s %-20s\n", fmt.Sprint(t.taskId), t.Title, t.Description, t.TargetDate.Format("2006-01-02"), t.createdDate.Format("2006-01-02"), t.status)
}
func DisplayAllTasks() {
	fmt.Printf("%-10s %-20s %-30s %-20s %-20s %-20s\n", "Task id", "Title", "Description", "Taget date", "Created date", "Status")
	fmt.Print("____________________________________________________________________________________\n")
	if len(store.Tasks) == 0 {
		fmt.Println("                       No task                                         ")
	} else {
		for _, task := range store.Tasks {
			task.DisplayTask()
		}
	}
}
func FindTaskById(id int64) *Task {
	for _, task := range store.Tasks {
		if task.taskId == id {
			return &task
		}
	}
	return nil
}

func FindTaskbyTask(title string) *Task {
	for _, task := range store.Tasks {
		if task.Title == title {
			return &task
		}
	}
	return nil
}

func UpdateTask(task Task) bool {
	t := FindTaskById(task.taskId)
	if task.Title != "" {
		t.Title = task.Title
	}
	if task.Description != "" {
		t.Description = task.Description
	}
	if task.TargetDate != t.TargetDate {
		t.TargetDate = task.TargetDate
	}
	t.Save()
	return true
}

func DeleteTask(id int64) bool {
	task := FindTaskById(id)
	var temp tasks = tasks{}
	for _, t := range store.Tasks {
		if t.taskId != task.taskId {
			temp.Tasks = append(temp.Tasks, t)
		}
	}
	store.Tasks = temp.Tasks
	return true
}
