package main

import (
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/x0y14/pm1"
)

func main() {
	p := tea.NewProgram(pm1.NewModel())
	if err := p.Start(); err != nil {
		log.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
