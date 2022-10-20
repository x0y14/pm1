package cli

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"time"
)

var (
	errStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000"))
)

func (m Model) View() string {
	var s string
	if m.err != nil {
		s += errStyle.Render(fmt.Sprintf("%v", m.err))
		s += "\n"
	}
	s += m.MainView.Render(m)
	s += "\n"
	return s
}

type View struct {
	Action func(m Model) Model
	Render func(m Model) string
}

var WaitingForToFinishLoadingStorage = View{
	Action: func(m Model) Model {
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("loading storage...\n")
	},
}

var WaitingForToFinishEnteringMasterPassword = View{
	Action: func(m Model) Model {
		enteredMasterPassword := m.textInput.Value()
		if enteredMasterPassword != "pass" {
			time.Sleep(time.Second * 1)
			m.err = fmt.Errorf("decryption failed")
			return m
		}
		m.err = nil
		m.textInput.Blur()
		m.MainView = Success
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("please entering master password (length: 4 < n)\n%s\n", m.textInput.View())
	},
}

var Success = View{
	Action: func(m Model) Model {
		return m
	},
	Render: func(m Model) string {
		return fmt.Sprintf("successful loading")
	},
}
