package cli

import (
	"fmt"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/x0y14/pm1/password"
	"io"
)

type confidentialItem password.Confidential

func (i confidentialItem) FilterValue() string { return "" }

type confidentialItemDelegate struct{}

func (d confidentialItemDelegate) Height() int                               { return 1 }
func (d confidentialItemDelegate) Spacing() int                              { return 0 }
func (d confidentialItemDelegate) Update(msg tea.Msg, m *list.Model) tea.Cmd { return nil }
func (d confidentialItemDelegate) Render(w io.Writer, m list.Model, index int, listItem list.Item) {
	i, ok := listItem.(confidentialItem)
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

func confidentialListView(vaultName string) View {
	return View{
		Action: func(m Model) Model {
			return m
		},
		Render: func(m Model) string {
			var items []list.Item
			vault, err := m.storage.FindVaultByName(vaultName)
			if err != nil {
				m.MainView = ErrorView(err)
				return fmt.Sprintf("%v", err)
			}

			for _, v := range vault.Data {
				items = append(items, confidentialItem(*v))
			}

			l := list.New(items, confidentialItemDelegate{}, 20, 14)
			l.Title = "Passwords"
			l.SetShowStatusBar(false)
			l.SetFilteringEnabled(false)
			l.Styles.Title = titleStyle
			l.Styles.PaginationStyle = paginationStyle
			l.Styles.HelpStyle = helpStyle
			return "\n" + l.View()
		},
	}
}
