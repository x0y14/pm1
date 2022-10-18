package pm1

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

type Model struct {
	storage   *Storage
	generator *PasswordGenerator
}

func NewModel(opts *Option, args []string) *Model {
	log.Printf("options: %v\n", opts)
	log.Printf("args: %v\n", args)
	return &Model{
		storage:   nil,
		generator: nil,
	}
}

func (m Model) Init() tea.Cmd {
	const encJsonPath = "secure/enc.json"

	// setup storage
	// find exported file
	if IsExistFile(encJsonPath) {
		//encrypted, iv, err := Load(encJsonPath)
		//if err != nil {
		//	log.Fatalf("failed to load: %s: %v", encJsonPath, err)
		//	return nil
		//}

	} else {
		// ファイルがなかったので新しく作る.
		// vault
		personalVault := NewVault("personal")
		m.storage = NewStorage()
		m.storage.Register(personalVault)
		// password generator
		m.generator = NewPasswordGenerator()
		err := m.generator.Init()
		if err != nil {
			log.Fatalf("failed to initialize password generator: %v", err)
		}
	}

	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m Model) View() string {
	return ""
}
