package main

import (
	"taskmanager/pkg/cmd"
	"taskmanager/pkg/menu"
)

func main() {
	menu.RenderMain()
	for {
		menu.ExecCmd(menu.MenuOption(cmd.Prompt()))
	}
}
