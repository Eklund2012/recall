package tui

import "github.com/Eklund2012/recall/internal/cards"

type model struct {
	store   *cards.Store
	index   int
	showAns bool
}

func NewModel(store *cards.Store) model {
	return model{
		store: store,
	}
}
