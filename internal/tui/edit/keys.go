package edit

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Save   key.Binding
	Cancel key.Binding
	Switch key.Binding
}

var Keys = keyMap{
	Save: key.NewBinding(
		key.WithKeys("ctrl+s"),
		key.WithHelp("ctrl+s", "save"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "cancel"),
	),
	Switch: key.NewBinding(
		key.WithKeys("tab"),
		key.WithHelp("tab", "switch field"),
	),
}
