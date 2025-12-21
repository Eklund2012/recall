package tui

import "fmt"

func (m model) View() string {
	if len(m.cards) == 0 {
		return "No cards to review.\n"
	}

	c := m.cards[m.index]
	view := fmt.Sprintf("Q: %s\n", c.Question)
	if m.showAns {
		view += fmt.Sprintf("A: %s\n", c.Answer)
	}

	view += "\n\nj/k: navigate • enter: reveal • q: quit"
	return view
}
