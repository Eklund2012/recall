package review

import (
	"fmt"
)

func (m model) View() string {
	deck, err := m.store.Active()
	if err != nil {
		// No active deck, quit or show a message
		return "No active deck selected.\n"
	}
	if len(deck.Cards) == 0 {
		return "No cards to review.\n"
	}

	c := deck.Cards[m.index]
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
