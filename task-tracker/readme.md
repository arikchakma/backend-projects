# Task Tracker

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## How to run

Clone the repository and run the following command:

```bash
git clone https://github.com/arikchakma/backend-projects.git
cd backend-projects/task-tracker
```

Run the following command to build and run the project:

```bash
go build -o task-tracker
./task-tracker --help # To see the list of available commands

# To add a task
./task-tracker add "Buy groceries"

# To update a task
./task-tracker update 1 "Buy groceries and cook dinner"

# To delete a task
./task-tracker delete 1

# To mark a task as in progress/done/todo
./task-tracker mark-in-progress 1
./task-tracker mark-done 1
./task-tracker mark-todo 1

# To list all tasks
./task-tracker list
./task-tracker list done
./task-tracker list todo
./task-tracker list in-progress
```
