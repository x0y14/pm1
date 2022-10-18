package main

import (
	"github.com/jessevdk/go-flags"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/x0y14/pm1"
)

func main() {
	opts := pm1.Option{}
	args, err := flags.Parse(&opts)
	if err != nil {
		log.Printf("faile to parse flags: %v", err)
	}

	p := tea.NewProgram(pm1.NewModel(&opts, args))
	if err := p.Start(); err != nil {
		log.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
