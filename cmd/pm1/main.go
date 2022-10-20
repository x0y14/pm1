package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
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

	opt := cli.Option{}
	args, err := flags.Parse(&opt)
	if err != nil {
		log.Printf("faile to parse flags: %v", err)
	}

	model := cli.InitialModel(&opt, args)
	p := tea.NewProgram(model)

	_, err = p.StartReturningModel()
	if err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
