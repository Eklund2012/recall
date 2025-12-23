package edit

import (
	"github.com/charmbracelet/bubbles/textarea"
	tea "github.com/charmbracelet/bubbletea"
)

// Message types for communicating with parent
type SaveMsg struct {
	Position int
	Question string
	Answer   string
}

type CancelMsg struct{}

func (m Model) Init() tea.Cmd {
	return textarea.Blink
}

func (m Model) Update(msg tea.Msg) (Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "esc":
			// Cancel edit
			return m, func() tea.Msg { return CancelMsg{} }

		case "ctrl+s":
			// Save changes
			return m, func() tea.Msg {
				return SaveMsg{
					Position: m.card.Position,
					Question: m.questionArea.Value(),
					Answer:   m.answerArea.Value(),
				}
			}

		case "tab", "shift+tab":
			// Switch focus between fields
			if m.focused == questionField {
				m.focused = answerField
				m.questionArea.Blur()
				cmds = append(cmds, m.answerArea.Focus())
			} else {
				m.focused = questionField
				m.answerArea.Blur()
				cmds = append(cmds, m.questionArea.Focus())
			}
			return m, tea.Batch(cmds...)
		}
	}

	// Update the focused textarea
	var cmd tea.Cmd
	if m.focused == questionField {
		m.questionArea, cmd = m.questionArea.Update(msg)
	} else {
		m.answerArea, cmd = m.answerArea.Update(msg)
	}
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
