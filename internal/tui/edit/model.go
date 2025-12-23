package edit

import (
	"github.com/Eklund2012/recall/internal/cards"
	"github.com/charmbracelet/bubbles/textarea"
)

type focusField int

const (
	questionField focusField = iota
	answerField
)

type Model struct {
	card         cards.Card
	questionArea textarea.Model
	answerArea   textarea.Model
	focused      focusField
	err          error
}

func NewModel(card cards.Card) Model {
	qa := textarea.New()
	qa.Placeholder = "Enter question..."
	qa.SetValue(card.Question)
	qa.Focus()

	aa := textarea.New()
	aa.Placeholder = "Enter answer..."
	aa.SetValue(card.Answer)

	return Model{
		card:         card,
		questionArea: qa,
		answerArea:   aa,
		focused:      questionField,
	}
}
