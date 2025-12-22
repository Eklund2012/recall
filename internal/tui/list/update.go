package list

import (
	tea "github.com/charmbracelet/bubbletea"
)

func (m model) Init() tea.Cmd {
	return nil
}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "d":
			if it, ok := m.list.SelectedItem().(item); ok {
				if err := m.store.Delete(it.index); err != nil {
					return m, nil
				}

				// Refresh list after delete
				m.list.SetItems(itemsFromStore(m.store))
			}
		}

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}
