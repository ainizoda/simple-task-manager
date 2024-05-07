package menu

import (
	"fmt"
	"os"
	"taskmanager/pkg/cmd"
	"taskmanager/pkg/task"
	"taskmanager/pkg/utils"

	"github.com/rodaine/table"
)

type TaskService interface {
	Create(title string, description string) error
	GetAll() []*task.Task
	Update(updatedTask *task.Task) error
	RemoveById(id string) error
}

type Menu struct {
	taskSvc TaskService
}

func NewMenu(taskSvc TaskService) *Menu {
	return &Menu{taskSvc: taskSvc}
}

func (m *Menu) showTasks() {
	tasks := m.taskSvc.GetAll()

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

func (m *Menu) createTask() {
	err := m.taskSvc.Create(
		cmd.ScanWithErr("Task title: "),
		cmd.ScanWithErr("Task description: "),
	)

	if err != nil {
		fmt.Printf("\n%s\n\n", err)
		return
	}
	m.ExecCmd(First)
}

func (m *Menu) updateTask() {
	m.ExecCmd(First)

	tasks := m.taskSvc.GetAll()
	num := cmd.Prompt("[UPDATE] Type № of task which you want to update: ", 1, len(tasks))

	taskPtr := tasks[num-1] // getting task pointer by index
	taskToUpd := *taskPtr   // copy its value, to prevent direct updating

	fmt.Println()
	fmt.Printf("1. Title: %s\n", taskToUpd.Title)
	fmt.Printf("2. Description: %s\n", taskToUpd.Description)
	fmt.Printf("3. Status: %s\n", taskToUpd.Status)

	opt := cmd.Prompt("\n(What do you want to update?): ", 1, 3)

	fields := []string{"Title", "Description", "Status"}
	propToUpd := fields[opt-1]

	if propToUpd == "Status" {
		m.updateTaskStatus(&taskToUpd)
	} else {
		newVal := cmd.ScanWithErr(fmt.Sprintf("\nNew %s: ", propToUpd))
		utils.SetField(&taskToUpd, propToUpd, newVal)
	}

	err := m.taskSvc.Update(&taskToUpd)

	if err != nil {
		fmt.Println(err)
	}

	m.ExecCmd(First) // return to task list
}

func (m *Menu) removeTask() {
	m.ExecCmd(First)

	tasks := m.taskSvc.GetAll()
	num := cmd.Prompt("[DELETE] Type № of task which you want to remove: ", 1, len(tasks))

	taskToBeRmvd := tasks[num-1]

	m.taskSvc.RemoveById(taskToBeRmvd.ID)
	m.ExecCmd(First)
}

func (m *Menu) updateTaskStatus(taskToUpd *task.Task) {
	for idx, status := range task.Statuses {
		fmt.Printf("\n%d. %s", idx+1, status)
	}

	opt := cmd.Prompt("\n\nSelect new status: ", 1, len(task.Statuses))

	utils.SetField(taskToUpd, "Status", task.Statuses[opt-1])
	m.ExecCmd(First)
}

func (m *Menu) exit() {
	os.Exit(1)
}

func (m *Menu) RenderMain() {
	fmt.Print("\nWelcome to simple task manager!\n")
	fmt.Print("\n1. Show tasks\n2. Create task\n3. Update task\n4. Remove Task\n0. Exit\n\n")
}

func (m *Menu) ExecCmd(opt MenuOption) {
	cmd.ClearWindow()
	m.RenderMain()

	routes := map[MenuOption]func(){}

	routes[First] = m.showTasks
	routes[Second] = m.createTask
	routes[Third] = m.updateTask
	routes[Fourth] = m.removeTask
	routes[Exit] = m.exit

	fn := routes[opt]
	fn()
}
