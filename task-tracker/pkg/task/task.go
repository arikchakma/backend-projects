package task

import (
	"fmt"
	"time"

	"github.com/charmbracelet/lipgloss"
)

type TaskStatus string

const (
	TASK_STATUS_TODO        TaskStatus = "todo"
	TASK_STATUS_IN_PROGRESS TaskStatus = "in-progress"
	TASK_STATUS_DONE        TaskStatus = "done"
)

type Task struct {
	ID          int64      `json:"id"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

func NewTask(id int64, description string) *Task {
	return &Task{
		ID:          id,
		Description: description,
		Status:      TASK_STATUS_TODO,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func ListTasks(status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var filteredTasks []Task
	switch status {
	case "all":
		filteredTasks = tasks
	case TASK_STATUS_TODO:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_TODO {
				filteredTasks = append(filteredTasks, task)
			}
		}

	case TASK_STATUS_IN_PROGRESS:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_IN_PROGRESS {
				filteredTasks = append(filteredTasks, task)
			}
		}

	case TASK_STATUS_DONE:
		for _, task := range tasks {
			if task.Status == TASK_STATUS_DONE {
				filteredTasks = append(filteredTasks, task)
			}
		}
	}

	fmt.Println()
	fmt.Println(
		lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66")).
			MarginBottom(1).
			Render(fmt.Sprintf("Tasks (%s)", status)))
	for _, task := range filteredTasks {
		formattedId := lipgloss.NewStyle().
			Bold(true).
			Width(5).
			Render(fmt.Sprintf("ID:%d", task.ID))
		formattedStatus := lipgloss.NewStyle().
			Bold(true).
			Width(12).
			Foreground(lipgloss.Color(statusColor(task.Status))).
			Render(string(task.Status))

		relativeUpdatedTime := task.UpdatedAt.Format("2006-01-02 15:04:05")

		taskStyle := lipgloss.NewStyle().
			Border(lipgloss.NormalBorder(), false, false, true, false).
			BorderForeground(lipgloss.Color("#3C3C3C")).
			Render(fmt.Sprintf("%s %s %s (%s)", formattedId, formattedStatus, task.Description, relativeUpdatedTime))
		fmt.Println(taskStyle)
	}
	fmt.Println()

	return nil
}

func statusColor(status TaskStatus) string {
	switch status {
	case TASK_STATUS_TODO:
		return "#3C3C3C"
	case TASK_STATUS_IN_PROGRESS:
		return "202"
	case TASK_STATUS_DONE:
		return "#04B575"
	default:
		return "#3C3C3C"
	}
}

func TaskStatusFromString(status string) TaskStatus {
	switch status {
	case "todo":
		return TASK_STATUS_TODO
	case "in-progress":
		return TASK_STATUS_IN_PROGRESS
	case "done":
		return TASK_STATUS_DONE
	default:
		return "all"
	}
}

func AddTask(description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var newTaskId int64
	if len(tasks) > 0 {
		lastTask := tasks[len(tasks)-1]
		newTaskId = lastTask.ID + 1
	} else {
		newTaskId = 1
	}

	task := NewTask(newTaskId, description)
	tasks = append(tasks, *task)

	style := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66"))

	formattedId := style.Render(fmt.Sprintf("(ID: %d)", task.ID))
	fmt.Printf("\nTask added successfully: %s\n\n", formattedId)
	return WriteTasksToFile(tasks)
}

func DeleteTask(id int64) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID != id {
			updatedTasks = append(updatedTasks, task)
		}
	}

	if len(updatedTasks) == len(tasks) {
		notFoundStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66")).
			Padding(1, 0).
			Render(fmt.Sprintf("Task not found (ID: %d)", id))
		fmt.Printf("%s", notFoundStyle)
		return nil
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask deleted successfully: %s\n\n", formattedId)
	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskStatus(id int64, status TaskStatus) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			switch status {
			case TASK_STATUS_TODO:
				task.Status = TASK_STATUS_TODO
			case TASK_STATUS_IN_PROGRESS:
				task.Status = TASK_STATUS_IN_PROGRESS
			case TASK_STATUS_DONE:
				task.Status = TASK_STATUS_DONE
			}
			task.UpdatedAt = time.Now()
		}

		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		notFoundStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66")).
			Padding(1, 0).
			Render(fmt.Sprintf("Task not found (ID: %d)", id))
		fmt.Printf("%s", notFoundStyle)
		return nil
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask updated successfully: %s\n\n", formattedId)
	return WriteTasksToFile(updatedTasks)
}

func UpdateTaskDescription(id int64, description string) error {
	tasks, err := ReadTasksFromFile()
	if err != nil {
		return err
	}

	var taskExists bool = false
	var updatedTasks []Task
	for _, task := range tasks {
		if task.ID == id {
			taskExists = true
			task.Description = description
			task.UpdatedAt = time.Now()
		}
		updatedTasks = append(updatedTasks, task)
	}

	if !taskExists {
		notFoundStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FFCC66")).
			Padding(1, 0).
			Render(fmt.Sprintf("Task not found (ID: %d)", id))
		fmt.Printf("%s", notFoundStyle)
		return nil
	}

	formattedId := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FFCC66")).
		Render(fmt.Sprintf("(ID: %d)", id))
	fmt.Printf("\nTask updated successfully: %s\n\n", formattedId)
	return WriteTasksToFile(updatedTasks)
}
