package expense

import (
	"fmt"

	"github.com/arikchakma/backend-projects/expense-tracker/internal/log"
)

type Budget struct {
	Month  int32   `json:"month"`
	Amount float64 `json:"amount"`
}

func NewBudget(month int32, amount float64) *Budget {
	return &Budget{
		Month:  month,
		Amount: amount,
	}
}

func BudgetMonth(month int32, amount float64) error {
	budgets, err := ReadBudgetsFromFile()
	if err != nil {
		return err
	}

	for i, budget := range budgets {
		if budget.Month == month {
			budgets[i].Amount = amount
			log.Success(fmt.Sprintf("Budget for %d updated to %.2f", month, amount))
			return WriteBudgetsToFile(budgets)
		}
	}

	budgets = append(budgets, Budget{
		Month:  month,
		Amount: amount,
	})
	log.Success(fmt.Sprintf("Budget for %d set to %.2f", month, amount))
	return WriteBudgetsToFile(budgets)
}

func GetMonthlyBudget(month int32) (float64, error) {
	budgets, err := ReadBudgetsFromFile()
	if err != nil {
		return 0, err
	}

	for _, budget := range budgets {
		if budget.Month == month {
			return budget.Amount, nil
		}
	}

	return 0, nil
}
