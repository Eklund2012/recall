package list

import "github.com/charmbracelet/bubbles/key"

type keyMap struct {
	Delete key.Binding
}

var keys = keyMap{
	Delete: key.NewBinding(
		key.WithKeys("d"),
		key.WithHelp("d", "delete"),
	),
}
