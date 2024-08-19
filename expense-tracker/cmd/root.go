package cmd

import "github.com/spf13/cobra"

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "expense-tracker",
		Short: "Expense Tracker is a CLI tool for managing expenses.",
		Long: `Manage your expenses with ease using Expense Tracker.

Complete code available at https://github.com/arikchakma/backend-projects`,
	}

	cmd.AddCommand(NewAddCmd())
	cmd.AddCommand(NewListCmd())
	cmd.AddCommand(NewSummaryCmd())
	cmd.AddCommand(NewDeleteCmd())
	cmd.AddCommand(NewBudgetCmd())

	return cmd
}
