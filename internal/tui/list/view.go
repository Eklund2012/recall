package list

func (m model) View() string {
	switch m.state {
	case listView:
		return m.list.View()
	case editView:
		return m.editModel.View()
	default:
		return m.list.View()
	}
}
