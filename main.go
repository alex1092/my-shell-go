package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"github.com/c-bata/go-prompt"
)

func main() {
	p := prompt.New(
		executor,
		completer,
		prompt.OptionPrefix("$ "),
	)

	p.Run()
}

func executor(in string) {
	clearIn := strings.TrimSpace(in)
	cmds := strings.Split(clearIn, " ")

	switch cmds[0] {
	case "exit":
		os.Exit(0)
	case "pwd":
		dir, err := os.Getwd()
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println(dir)
		}
	case "cd":
		cd(cmds[1:])
	case "nvim":
		nvim(cmds[1:])
	case "echo":
		fmt.Println(strings.Join(cmds[1:], " "))
	case "type":
		handleType(cmds[1:])
	case "cursor":
		cursor(cmds[1:])
	default:
		runCommand(cmds)
	}
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "exit", Description: "Exit the shell"},
		{Text: "pwd", Description: "Print working directory"},
		{Text: "cd", Description: "Change directory"},
		{Text: "nvim", Description: "Open file with nvim"},
		{Text: "echo", Description: "Echo arguments"},
		{Text: "type", Description: "Describe a command"},
		{Text: "cursor", Description: "Run cursor command"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func cursor(args []string) {
	cmd := exec.Command("cursor", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println(err.Error())
	}
}

func nvim(args []string) {
	if len(args) == 0 {
		fmt.Println("code: missing file operand")
		return
	}
	cmd := exec.Command("nvim", args[0])
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("code: %s\n", err.Error())
	}
}

func cd(args []string) {
	var dir string
	if len(args) == 0 || args[0] == "~" {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stdout, "Unable to find home directory\n")
			return
		}
		dir = homeDir
	} else {
		dir = args[0]
	}

	if err := os.Chdir(dir); err != nil {
		fmt.Fprintf(os.Stdout, "%s: No such file or directory\n", dir)
	}
}

func handleType(args []string) {
	if len(args) == 0 {
		fmt.Println("type: missing operand")
		return
	}
	switch args[0] {
	case "exit", "echo", "type", "pwd":
		fmt.Printf("%s is a shell builtin\n", args[0])
	default:
		paths := strings.Split(os.Getenv("PATH"), ":")
		isFound := false
		for _, path := range paths {
			fullPath := filepath.Join(path, args[0])
			if _, err := os.Stat(fullPath); !os.IsNotExist(err) {
				fmt.Printf("%s is %s\n", args[0], fullPath)
				isFound = true
				break
			}
		}
		if !isFound {
			fmt.Printf("%s: not found\n", args[0])
		}
	}
}

func runCommand(cmds []string) {
	command := exec.Command(cmds[0], cmds[1:]...)
	command.Stderr = os.Stderr
	command.Stdout = os.Stdout
	err := command.Run()
	if err != nil {
		fmt.Printf("%s: command not found\n", cmds[0])
	}
}
