package main

import "taskmanager/pkg/menu"

func main() {
	menu.RenderMain()
	for {
		menu.ExecCmd(menu.Prompt())
	}
}
