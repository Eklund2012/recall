package edit

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

var (
	focusedStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62"))

	blurredStyle = lipgloss.NewStyle().
			BorderStyle(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("240"))

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241"))
)

func (m Model) View() string {
	var b strings.Builder

	b.WriteString("Edit Flashcard\n\n")

	// Question field
	questionLabel := "Question:"
	if m.focused == questionField {
		questionLabel = focusedStyle.Render(questionLabel)
	} else {
		questionLabel = blurredStyle.Render(questionLabel)
	}
	b.WriteString(questionLabel + "\n")
	b.WriteString(m.questionArea.View() + "\n\n")

	// Answer field
	answerLabel := "Answer:"
	if m.focused == answerField {
		answerLabel = focusedStyle.Render(answerLabel)
	} else {
		answerLabel = blurredStyle.Render(answerLabel)
	}
	b.WriteString(answerLabel + "\n")
	b.WriteString(m.answerArea.View() + "\n\n")

	// Help text
	help := helpStyle.Render(
		"tab: switch fields • ctrl+s: save • esc: cancel",
	)
	b.WriteString(help)

	if m.err != nil {
		b.WriteString("\n\n" + fmt.Sprintf("Error: %v", m.err))
	}

	return b.String()
}
