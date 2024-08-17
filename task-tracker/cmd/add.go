package cmd

import (
	"errors"

	"github.com/arikchakma/backend/task-tracker/pkg/task"
	"github.com/spf13/cobra"
)

func NewAddCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "add",
		Short: "Add a task to the task list",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddTaskCmd(args)
		},
	}
	return cmd
}

func RunAddTaskCmd(args []string) error {
	if len(args) == 0 {
		return errors.New("task description is required")
	}

	description := args[0]
	return task.AddTask(description)
}
