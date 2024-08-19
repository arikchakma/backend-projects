package cmd

import (
	"fmt"

	"github.com/arikchakma/backend-projects/expense-tracker/pkg/expense"
	"github.com/spf13/cobra"
)

var Description string
var Amount float64
var Category string

func NewAddCmd() *cobra.Command {
	addCmd := &cobra.Command{
		Use:   "add",
		Short: "Insert a new expense",
		RunE: func(cmd *cobra.Command, args []string) error {
			return RunAddExpenseCmd(args)
		},
	}

	addCmd.Flags().StringVarP(&Description, "description", "d", "", "Description of the expense (required)")
	addCmd.MarkFlagRequired("description")
	addCmd.Flags().Float64VarP(&Amount, "amount", "a", 0, "Amount of the expense (required)")
	addCmd.MarkFlagRequired("amount")
	addCmd.Flags().StringVarP(&Category, "category", "c", "general", "Category of the expense")

	return addCmd
}

func RunAddExpenseCmd(args []string) error {
	if Amount < 0 {
		return fmt.Errorf("amount cannot be negative")
	}

	return expense.AddExpense(Description, Amount, Category)
}
