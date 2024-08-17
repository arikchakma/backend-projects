package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "task-tracker",
		Short: "Task Tracker is a CLI tool for managing tasks",
		Long: `Task Tracker is a CLI tool for managing tasks. It allows you to create, list, and delete tasks.
    
You can also mark tasks as completed and update their status.
Complete code available at https://github.com/arikchakma/backend-projects`,
	}

	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewUpdateCmd())
	cmd.AddCommand(NewStatusDoneCmd())
	cmd.AddCommand(NewStatusInProgressCmd())
	cmd.AddCommand(NewStatusTodoCmd())

	return cmd
}
