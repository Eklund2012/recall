package tui

import (
	"github.com/Eklund2012/recall/internal/cards"
)

type model struct {
	cards   []cards.Card
	index   int
	showAns bool
}

func NewModel() model {
	return model{
		cards:   cards.Cards,
		index:   0,
		showAns: false,
	}
}
