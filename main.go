package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	for {
		fmt.Fprint(os.Stdout, "$ ")
		in, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			fmt.Println(err.Error())
			continue
		}
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
		case "echo":
			fmt.Println(strings.Join(cmds[1:], " "))
		case "type":
			handleType(cmds[1:])
		default:
			runCommand(cmds)
		}
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
