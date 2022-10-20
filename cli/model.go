package cli

import (
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/x0y14/pm1/password"
)

const (
	// exportPath
	// 暗号化されたストレージバイナリデータと、IVを記したjsonファイルの設置場所
	exportPath = "secure/enc.json"
)

type Model struct {
	err      error
	MainView View

	textInput           textinput.Model
	masterPasswordInput textinput.Model

	storage *password.Storage
}

func InitialModel(opt *Option, args []string) Model {
	ti := textinput.New()
	ti.Prompt = "> "

	mi := textinput.New()
	mi.Prompt = "> "
	mi.EchoMode = textinput.EchoPassword
	mi.Placeholder = "master password"

	return Model{
		MainView:            FindingEncJson,
		textInput:           ti,
		masterPasswordInput: mi,
	}
}

func (m Model) Init() tea.Cmd {
	return func() tea.Msg {
		if password.IsExistFile(exportPath) {
			return EventMsg{EventType: EncJsonFound}
		}
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
			m.MainView = CreatingNewStorageAndVault1
			m.masterPasswordInput.Focus()
		case EncJsonFound:
			m.MainView = CheckEncJson
			m = m.MainView.Action(m)
		}
	}

	if m.textInput.Focused() {
		m.textInput, cmd = m.textInput.Update(msg)
		return m, cmd
	}
	if m.masterPasswordInput.Focused() {
		m.masterPasswordInput, cmd = m.masterPasswordInput.Update(msg)
		return m, cmd
	}

	return m, cmd
}
