package db

import (
	"encoding/json"
	"fmt"
	"os"
)

func check(err error, message string) {
	if err != nil {
		fmt.Println(message);
		os.Exit(1)
	}
}

func GetTasks() []string {
	file, err := readJsonFile()
	check(err, "Failed to read json file")
	
	var tasks []string

	err = json.NewDecoder(file).Decode(&tasks)
	check(err, "Failed to decode json file")

	return tasks
}

func AddTask(task string) {
	tasks := GetTasks()
	tasks = append(tasks, task)
	writeTasksToFile(tasks)
}

func contains[T comparable](s []T, e T) bool {
	for _, v := range s {
			if v == e {
					return true
			}
	}
	return false
}


func DeleteTasks(ids []int) {
	tasks := GetTasks()

	var newTasks = []string{}

	for i, task := range tasks {
		if !contains(ids, i) {
			newTasks = append(newTasks, task)
		}
	}

	writeTasksToFile(newTasks)
}

func writeTasksToFile(tasks []string) {
	file, err := os.Create("tasks.json")
	check(err, "Failed to create json file")

	err = json.NewEncoder(file).Encode(tasks)
	check(err, "Failed to encode json file")
}

func readJsonFile() (*os.File, error) {
	file, err := os.Open("tasks.json")

	if err != nil {
		return nil, err
	}

	return file, nil
}

func createInitialJsonFile() {
	err := os.WriteFile("tasks.json", []byte("[]"), 0644)
	check(err, "Failed to create initial json file")
}

func init() {
	_, err := readJsonFile()

	if err != nil {
		createInitialJsonFile()
	}
}