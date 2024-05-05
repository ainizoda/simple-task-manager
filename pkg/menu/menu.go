package menu

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
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
		scan("Task title: "),
		scan("Task description: "),
	)

	if err != nil {
		fmt.Printf("\n%s\n\n", err)
		return
	}
	ExecCmd(First)
}

func RenderMain() {
	fmt.Print("\nWelcome to simple task manager!\n")
	fmt.Print("\n1. Show tasks\n2. Create task\n3. Update task\n4. Remove Task\n0. Exit\n\n")
}

func ExecCmd(opt MenuOption) {
	clearWindow()
	RenderMain()

	switch opt {
	case First:
		showTasks()
	case Second:
		createTask()
	case Exit:
		os.Exit(1)
	}
}

func Prompt() MenuOption {
	var opt MenuOption

	for {
		input, err := strconv.Atoi(scan("$: "))

		if input < 0 || input > 4 || err != nil {
			fmt.Print("invalid option!\n\n")
			continue
		}
		opt = MenuOption(input)
		break
	}
	return opt
}

func clearWindow() {
	var command []string

	switch runtime.GOOS {
	case "windows":
		command = []string{"cmd", "/c", "cls"}
	default:
		command = []string{"clear"}
	}
	cmd := exec.Command(command[0], command[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func scan(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("can't access the buffer")
		os.Exit(2)
	}
	return scanner.Text()
}
