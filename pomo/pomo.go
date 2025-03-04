/*
Copyright © 2025 CARLOS MIRANDA carlosermiranda.dev@gmail.com
*/
package pomo

import (
	"encoding/json"
	"os"
)

type TaskPomo struct {
	Id          string
	Description string
	Pomodoro    int
	TimeSpent   int
	Done        bool
	Priority    int
}

func SaveTasks(filename string, tasks []TaskPomo) error {
	tasksJson, err := json.Marshal(tasks)

	if err != nil {
		return err
	}

	err = os.WriteFile(filename, tasksJson, 0644)

	if err != nil {
		return err
	}

	return nil
}

func ReadTasks(filename string) ([]TaskPomo, error) {
	tasksPomoFile, err := os.ReadFile(filename)

	if err != nil {
		return []TaskPomo{}, err
	}

	var tasksPomo []TaskPomo

	if err := json.Unmarshal(tasksPomoFile, &tasksPomo); err != nil {
		return []TaskPomo{}, err
	}

	return tasksPomo, nil
}
