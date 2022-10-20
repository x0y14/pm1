package pm1

import (
	"github.com/charmbracelet/bubbles/spinner"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Seen int

const (
	Launched Seen = iota
	WaitingForFinishLoadingEnglishDictionary
	WaitingForEnteringMasterPassword
)

var seenTypes = [...]string{
	Launched:                                 "Launched",
	WaitingForFinishLoadingEnglishDictionary: "WaitingForFinishLoadingEnglishDictionary",
	WaitingForEnteringMasterPassword:         "WaitingForEnteringMasterPassword",
}

func (s Seen) String() string {
	return seenTypes[s]
}

type SeenChangeMsg struct {
	NewSeen Seen
}

var (
	spinnerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("63"))
)

type Model struct {
	opt       *Option
	args      []string
	seen      Seen
	spinner   spinner.Model
	isLoading bool
}

func NewModel(opt *Option, args []string) Model {
	s := spinner.New()
	s.Style = spinnerStyle

	return Model{
		opt:       opt,
		args:      args,
		seen:      Launched,
		spinner:   s,
		isLoading: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "ctrl+c":
			return m, tea.Quit
		}
	case SeenChangeMsg:
		m.seen = msg.NewSeen
		switch m.seen {
		case WaitingForFinishLoadingEnglishDictionary, WaitingForEnteringMasterPassword:
			m.isLoading = true
			return m, spinner.Tick
		default:
			m.isLoading = false
			return m, nil
		}
	case spinner.TickMsg:
		var cmd tea.Cmd
		m.spinner, cmd = m.spinner.Update(msg)
		return m, cmd
	}

	return m, nil
}

func (m Model) View() string {
	var s string
	if m.isLoading {
		s += m.spinner.View() + " "
	}
	s += "now: " + m.seen.String()
	return s
}
