package expense

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

func expensesFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}

	return path.Join(cwd, "expenses.json")
}

func budgetsFilePath() string {
	cwd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return ""
	}

	return path.Join(cwd, "budgets.json")
}

func ReadExpensesFromFile() ([]Expense, error) {
	filePath := expensesFilePath()
	// Check if file exists
	// if doesn't exist, return empty list and create the file
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file...")
		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			return nil, err
		}

		defer file.Close()

		return []Expense{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	defer file.Close()

	tasks := []Expense{}
	err = json.NewDecoder(file).Decode(&tasks)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func WriteExpensesToFile(tasks []Expense) error {
	filePath := expensesFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}

func ReadBudgetsFromFile() ([]Budget, error) {
	filePath := budgetsFilePath()
	// Check if file exists
	// if doesn't exist, return empty list and create the file
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		fmt.Println("File does not exist. Creating file...")
		file, err := os.Create(filePath)
		os.WriteFile(filePath, []byte("[]"), os.ModeAppend.Perm())

		if err != nil {
			fmt.Println("Error creating file:", err)
			return nil, err
		}

		defer file.Close()

		return []Budget{}, nil
	}

	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return nil, err
	}

	defer file.Close()

	budgets := []Budget{}
	err = json.NewDecoder(file).Decode(&budgets)
	if err != nil {
		fmt.Println("Error decoding file:", err)
		return nil, err
	}

	return budgets, nil
}

func WriteBudgetsToFile(budgets []Budget) error {
	filePath := budgetsFilePath()
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	err = json.NewEncoder(file).Encode(budgets)
	if err != nil {
		return err
	}

	return nil
}
