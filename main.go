package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	TaskID      int
	ProgressID  int
	Description string
}

/*
Task Progress :
1. Progress Done
2. Progress Not Done
3. Progress In Progress
*/
type TaskProgress struct {
	ProgressID int
	Status     string
}

func RemoveTask(s []Task, index int) []Task {
	var newDataTask []Task
	for _, task := range s {
		if task.TaskID != index {
			// Only append tasks that do not match the taskID
			newDataTask = append(newDataTask, task)
		}
	}
	return newDataTask

}

func UpdateTask(s []Task, index int, value string) []Task {

	for i, task := range s {
		if task.TaskID == index {
			// Update the task's description
			s[i].Description = value
		}
	}
	return s
}

func filterTasksByProgressID(tasks []Task, value int) (filteredTasks []Task) {
	for _, task := range tasks {
		if task.ProgressID == value {
			filteredTasks = append(filteredTasks, task)
		}
	}

	return
}

func main() {
	tasks := make([]Task, 0)

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("...Task Tracker...")
		fmt.Println("1. Add Task")
		fmt.Println("2. Update Task")
		fmt.Println("3. Delete Task")
		fmt.Println("4. List Task")
		fmt.Print("Choose Option : ")

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		switch input {
		case "1":
			fmt.Print("Enter the task description : ")

			description, _ := reader.ReadString('\n')

			description = strings.TrimSpace(description)

			fmt.Println()
			fmt.Println("1. Progress Done")
			fmt.Println("2. Progress Not Done")
			fmt.Println("3. Progress In Progress")
			fmt.Print("Enter the task progress : ")

			progressid, _ := reader.ReadString('\n')

			// fmt.Printf("Tipe data progressid: %T, nilai progressid: %v\n", progressid, progressid)

			progressid_, err := strconv.Atoi(strings.TrimSpace(progressid))

			// fmt.Printf("Tipe data progressid_: %T, nilai progressid_: %v\n", progressid_, progressid_)

			if err != nil {
				fmt.Println("Please enter a valid numeric value-1.")
			} else if progressid_ == 0 {
				fmt.Println("Please enter a valid numeric value-2.")
			} else {
				taskID := len(tasks) + 1

				tasks = append(tasks, Task{
					TaskID:      taskID,
					ProgressID:  progressid_,
					Description: description,
				})

				fmt.Println("Task added successfully.")
			}
		case "2":

			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Description)
			}
			fmt.Print("Choose the task : ")

			input, _ := reader.ReadString('\n')

			input_, _ := strconv.Atoi(strings.TrimSpace(input))

			fmt.Print("Enter new the task description : ")

			newinput, _ := reader.ReadString('\n')

			newinput = strings.TrimSpace(newinput)

			result := UpdateTask(tasks, input_, newinput)
			fmt.Println()
			fmt.Println(result)
			fmt.Println("Task has been updated.")
		case "3":

			for i, task := range tasks {
				fmt.Printf("%d. %s\n", i+1, task.Description)
			}
			fmt.Print("Choose the task : ")

			input, _ := reader.ReadString('\n')

			input_, _ := strconv.Atoi(strings.TrimSpace(input))

			result := RemoveTask(tasks, input_)

			tasks = result

			fmt.Println()
			fmt.Println("Result")
			fmt.Println(result)
			fmt.Println("Task has been removed.")

		case "4":
			// tasks := []Task{
			// 	{1, 2, "Task I"},
			// 	{2, 1, "Task II"},
			// 	{3, 2, "Task III"},
			// }
			if len(tasks) == 0 {
				fmt.Println("No tasks")
			} else {

				fmt.Println("List of task progress:")
				fmt.Println("1. Progress Done")
				fmt.Println("2. Progress Not Done")
				fmt.Println("3. Progress In Progress")
				fmt.Println("4. All Task Progress")
				fmt.Print("Enter the task progress : ")

				input, _ := reader.ReadString('\n')

				input_, _ := strconv.Atoi(strings.TrimSpace(input))

				switch input_ {
				case 4:
					fmt.Println("tasks", tasks)
					for i, task := range tasks {
						fmt.Printf("%d. %s\n", i+1, task.Description)
					}
				default:
					filteredtasks := filterTasksByProgressID(tasks, input_)
					fmt.Println("filteredtasks", filteredtasks)
					for i, task := range filteredtasks {
						fmt.Printf("%d. %s\n", i+1, task.Description)
					}
				}
			}
		}

	}
}
