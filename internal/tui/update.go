package tui

import tea "github.com/charmbracelet/bubbletea"

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		case "enter", " ":
			m.showAns = !m.showAns
		case "j", "down":
			if m.index < len(m.cards)-1 {
				m.index++
				m.showAns = false
			}
		case "k", "up":
			if m.index > 0 {
				m.index--
				m.showAns = false
			}
		}
	}
	return m, nil
}
