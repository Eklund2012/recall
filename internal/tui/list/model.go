package list

import (
	"github.com/Eklund2012/recall/internal/cards"
	"github.com/Eklund2012/recall/internal/tui/edit"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type viewState int

const (
	listView viewState = iota
	editView
)

type model struct {
	list      list.Model
	editModel edit.Model
	store     *cards.Store
	state     viewState
}

func NewModel(store *cards.Store) model {
	l := list.New(
		itemsFromStore(store),
		list.NewDefaultDelegate(),
		0,
		0,
	)

	l.Title = "Flashcards"

	l.AdditionalShortHelpKeys = func() []key.Binding {
		return []key.Binding{
			keys.Delete,
			keys.Edit,
		}
	}

	return model{
		list:  l,
		store: store,
		state: listView,
	}
}

func itemsFromStore(store *cards.Store) []list.Item {
	deck, err := store.Active()
	if err != nil {
		return []list.Item{}
	}

	items := make([]list.Item, 0, len(deck.Cards))
	for _, c := range deck.Cards {
		items = append(items, item{
			title: c.Question,
			desc:  c.Answer,
			index: c.Position,
		})
	}
	return items
}
