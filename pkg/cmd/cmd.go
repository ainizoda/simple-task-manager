package cmd

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
)

func Prompt(msg string, from, to int) int {
	for {
		input, err := strconv.Atoi(ScanWithErr(msg))

		if input < from || input > to || err != nil {
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

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
	}
}

func ScanWithErr(text string) string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(text)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Println("can't access the buffer")
		os.Exit(2)
	}
	return scanner.Text()
}
