package cli

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"github.com/x0y14/pm1/password"
	"io"
)

var (
	titleStyle        = lipgloss.NewStyle().MarginLeft(2)
	itemStyle         = lipgloss.NewStyle().PaddingLeft(4)
	selectedItemStyle = lipgloss.NewStyle().PaddingLeft(2).Foreground(lipgloss.Color("170"))
	paginationStyle   = list.DefaultStyles().PaginationStyle.PaddingLeft(4)
	helpStyle         = list.DefaultStyles().HelpStyle.PaddingLeft(4).PaddingBottom(1)
	quitTextStyle     = lipgloss.NewStyle().Margin(1, 0, 2, 4)
)

type vaultItem password.Vault

func (i vaultItem) FilterValue() string { return "" }

type vaultItemDelegate struct{}

func (d vaultItemDelegate) Height() int                               { return 1 }
func (d vaultItemDelegate) Spacing() int                              { return 0 }
func (d vaultItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d vaultItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(vaultItem)
	if !ok {
		return
	}

	str := fmt.Sprintf("%d. %s", index+1, i.Name)

	fn := itemStyle.Render
	if index == m.Index() {
		fn = func(s string) string {
			return selectedItemStyle.Render("> " + s)
		}
	}

	fmt.Fprint(w, fn(str))
}

func vaultListView(storage *password.Storage) View {
	var items []list.Item
	for _, v := range storage.Vaults {
		items = append(items, vaultItem(*v))
	}
	l := list.New(items, vaultItemDelegate{}, 20, 14)
	l.Title = "Vaults"
	l.SetShowStatusBar(false)
	l.SetFilteringEnabled(false)
	l.Styles.Title = titleStyle
	l.Styles.PaginationStyle = paginationStyle
	l.Styles.HelpStyle = helpStyle

	return View{
		Action: func(m Model) Model {
			i, ok := l.SelectedItem().(vaultItem)
			if ok {
				m.MainView = confidentialListView(i.Name)
			}
			return m
		},
		Render: func(m Model) string {
			return "\n" + l.View()
		},
	}
}
