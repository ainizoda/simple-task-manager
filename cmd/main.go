package main

import (
	"taskmanager/pkg/cmd"
	"taskmanager/pkg/menu"
)

func main() {
	for {
		menu.ExecCmd(menu.MenuOption(cmd.Prompt()))
	}
}
