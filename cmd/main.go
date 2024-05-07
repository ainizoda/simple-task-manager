package main

import (
	"taskmanager/pkg/cmd"
	"taskmanager/pkg/menu"
	"taskmanager/pkg/task"
)

func main() {
	taskSvc := task.Service{}
	m := menu.NewMenu(&taskSvc)
	m.RenderMain()

	for {
		m.ExecCmd(menu.MenuOption(cmd.Prompt("$: ", 0, 4)))
	}
}
