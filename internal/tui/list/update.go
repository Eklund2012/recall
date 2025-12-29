package list

import (
	"github.com/Eklund2012/recall/internal/tui/edit"
	tea "github.com/charmbracelet/bubbletea"
)

// Custom message types for state transitions
type editSavedMsg struct {
	position int
	question string
	answer   string
}

type editCancelledMsg struct{}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// Handle state-specific routing
	switch m.state {
	case listView:
		return m.updateList(msg)
	case editView:
		return m.updateEdit(msg)
	}
	return m, nil
}

func (m model) updateList(msg tea.Msg) (tea.Model, tea.Cmd) {
	deck, err := m.store.Active()
	if err != nil {
		m.list.NewStatusMessage("⚠ No active deck")
		return m, nil
	}

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		// Edit key
		case "e":
			if it, ok := m.list.SelectedItem().(item); ok {
				card := deck.Cards[it.index-1]
				m.editModel = edit.NewModel(card)
				m.state = editView
				return m, m.editModel.Init()
			}

		// Delete key
		case "d":
			if it, ok := m.list.SelectedItem().(item); ok {
				if err := m.store.DeleteCard(m.store.ActiveDeck, it.index); err != nil {
					m.list.NewStatusMessage("⚠ Failed to delete")
					return m, nil
				}

				// Refresh list
				m.list.SetItems(itemsFromStore(m.store))
				m.list.NewStatusMessage("✓ Card deleted")
			}
		}

	case editSavedMsg:
		// Handle save from edit mode
		err := m.store.Update(msg.position, msg.question, msg.answer)
		m.state = listView
		m.list.SetItems(itemsFromStore(m.store))
		if err != nil {
			m.list.NewStatusMessage("⚠ Failed to save")
		} else {
			m.list.NewStatusMessage("✓ Card updated")
		}
		return m, nil

	case editCancelledMsg:
		// Return to list without saving
		m.state = listView
		m.list.NewStatusMessage("✗ Edit cancelled")
		return m, nil

	case tea.WindowSizeMsg:
		m.list.SetSize(msg.Width, msg.Height-3)
	}

	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) updateEdit(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	// Check for save/cancel messages from edit model
	switch msg := msg.(type) {
	case edit.SaveMsg:
		// Convert to internal message
		return m.updateList(editSavedMsg{
			position: msg.Position,
			question: msg.Question,
			answer:   msg.Answer,
		})

	case edit.CancelMsg:
		// Convert to internal message
		return m.updateList(editCancelledMsg{})
	}

	// Let edit model handle its own updates
	m.editModel, cmd = m.editModel.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}
