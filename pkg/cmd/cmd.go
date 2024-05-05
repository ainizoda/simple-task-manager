package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func Prompt() int {
	for {
		input, err := strconv.Atoi(Scan("$: "))

		if input < 0 || input > 4 || err != nil {
			fmt.Print("invalid option!\n\n")
			continue
		}
		return input
	}
}

func ClearWindow() {
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

func Scan(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("can't access the buffer")
		os.Exit(2)
	}
	return scanner.Text()
}
