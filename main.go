package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"github.com/x0y14/pm1/command"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/x0y14/pm1/cli"
)

const (
	exportPath = "secure/enc.json"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	flag.Parse()
	if len(flag.Args()) < 1 {
		fmt.Println(command.Help())
		os.Exit(0)
	}

	var cmd *command.Command
	var err error
	switch flag.Arg(0) {
	case "help":
		fmt.Println(command.Help())
		os.Exit(0)
	case "vault":
		cmd, err = command.Vault()
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	default:
		fmt.Printf("unknown command: %s\n\n", flag.Arg(0))
		fmt.Println(command.Help())
		os.Exit(0)
	}

	model := cli.InitialModel(cmd)
	p := tea.NewProgram(model)

	_, err = p.StartReturningModel()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
