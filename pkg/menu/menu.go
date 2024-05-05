package menu

import (
	"errors"
	"fmt"
	"os"
	"taskmanager/pkg/cmd"
	"taskmanager/pkg/task"

	"github.com/rodaine/table"
)

var (
	ErrInvalidOption = errors.New("error: invalid option")
	tasksSvc         = task.Service{}
)

func showTasks() {
	tasks := tasksSvc.GetAll()
	if len(tasks) == 0 {
		fmt.Printf("- You have no tasks yet, type %d to create one\n\n", Second)
		return
	}
	fmt.Print("Current tasks: \n")
	tbl := table.New("№", "Title", "Description", "Status", "Created At")

	for idx, t := range tasks {
		tbl.AddRow(
			idx+1,
			t.Title,
			t.Description,
			t.Status,
			t.CreatedAt,
		)
	}
	tbl.Print()
	fmt.Println()
}

func createTask() {
	err := tasksSvc.Create(
		cmd.ScanWithErr("Task title: "),
		cmd.ScanWithErr("Task description: "),
	)

	if err != nil {
		fmt.Printf("\n%s\n\n", err)
		return
	}
	ExecCmd(First)
}

func renderMain() {
	fmt.Print("\nWelcome to simple task manager!\n")
	fmt.Print("\n1. Show tasks\n2. Create task\n3. Update task\n4. Remove Task\n0. Exit\n\n")
}

func ExecCmd(opt MenuOption) {
	cmd.ClearWindow()
	renderMain()

	switch opt {
	case First:
		showTasks()
	case Second:
		createTask()
	case Exit:
		os.Exit(1)
	}
}
