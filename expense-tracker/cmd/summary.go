package cmd

import (
	"github.com/arikchakma/backend-projects/expense-tracker/internal/expense"
	"github.com/spf13/cobra"
)

var SummaryMonth int

func NewSummaryCmd() *cobra.Command {
	summaryCmd := &cobra.Command{
		Use:   "summary",
		Short: "Show a brief summary of expenses",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunSummaryExpenseCmd(args)
		},
	}

	summaryCmd.Flags().IntVarP(&SummaryMonth, "month", "m", 0, "Month for which to show the summary")

	return summaryCmd
}

func RunSummaryExpenseCmd(args []string) error {
	return expense.SummaryExpenses(SummaryMonth)
}
