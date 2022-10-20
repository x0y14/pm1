package cli

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"time"
)

type Model struct {
	MainView  View
	textInput textinput.Model
}

func InitialModel(opt *Option, args []string) Model {
	ti := textinput.New()
	ti.Prompt = "> "
	return Model{
		MainView:  WaitingForToFinishLoadingStorage,
		textInput: ti,
	}
}

func (m Model) Init() tea.Cmd {
	// call action of WaitingForToFinishLoadingStorage
	return func() tea.Msg {
		time.Sleep(time.Second * 2)
		return EventMsg{EventType: EncJsonNotFound}
	}
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			m = m.MainView.Action(m)
		}
	case EventMsg:
		switch msg.EventType {
		case EncJsonNotFound:
			m.textInput.Focus()
			m.MainView = WaitingForToFinishEnteringMasterPassword
			time.Sleep(time.Second * 2)
			m.MainView = Success
		case EncJsonFound:
			m.MainView = WaitingForToFinishLoadingStorage
		}
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}

func (m Model) View() string {
	var s string
	s += m.MainView.Render(m)
	s += "\n"
	return s
}
