package tui

import (
	"fmt"
)

func (m model) View() string {
	if len(m.store.Cards) == 0 {
		return "No cards to review.\n"
	}

	c := m.store.Cards[m.index]
	if c.Question == "" {
		return "Card question is empty.\n"
	}
	view := fmt.Sprintf("Q%d: %s\n", c.Position, c.Question)
	if m.showAns {
		view += fmt.Sprintf("A%d: %s\n", c.Position, c.Answer)
	}

	view += "\n\nj/k: navigate • enter: reveal • q: quit"
	return view
}
