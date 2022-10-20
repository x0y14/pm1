package main

import (
	"fmt"
	"github.com/jessevdk/go-flags"
	"log"
	"math/rand"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/x0y14/pm1"
)

func main() {
	opt := pm1.Option{}
	args, err := flags.Parse(&opt)
	if err != nil {
		log.Printf("faile to parse flags: %v", err)
	}

	rand.Seed(time.Now().UTC().UnixNano())

	p := tea.NewProgram(pm1.NewModel(&opt, args))

	// Simulate activity
	go func() {
		for {
			time.Sleep(2 * time.Second)
			p.Send(pm1.SeenChangeMsg{NewSeen: pm1.WaitingForFinishLoadingEnglishDictionary})
			time.Sleep(2 * time.Second)
			p.Send(pm1.SeenChangeMsg{NewSeen: pm1.WaitingForEnteringMasterPassword})
			time.Sleep(2 * time.Second)
			p.Send(pm1.SeenChangeMsg{NewSeen: pm1.Launched})
		}
	}()

	if _, err := p.StartReturningModel(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
