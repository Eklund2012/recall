package list

import (
	"github.com/Eklund2012/recall/internal/cards"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
)

type model struct {
	list  list.Model
	store *cards.Store
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
		}
	}

	return model{
		list:  l,
		store: store,
	}
}

func itemsFromStore(store *cards.Store) []list.Item {
	items := make([]list.Item, 0, len(store.Cards))

	for _, c := range store.Cards {
		items = append(items, item{
			title: c.Question,
			desc:  c.Answer,
			index: c.Position,
		})
	}

	return items
}
