package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	screen  string
	list    list.Model
	Welcome string
}

func InitialModel() Model {
	l := NewMenu()

	return Model{
		screen:  "#menu",
		list:    l,
		Welcome: "STEGO- Golang Based Stegnography Tool.",
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		h, v := docStyle.GetFrameSize()

		welcomeHeight := 2

		m.list.SetSize(
			msg.Width-h,
			msg.Height-v-welcomeHeight,
		)

	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "enter":
			selected := m.list.SelectedItem().(item)
			switch selected.Title() {
			case "#encode":
				m.screen = "#encode"
			case "#decode":
				m.screen = "#decode"
			case "#exit":
				return m, tea.Quit

			}

		}

	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	content := m.Welcome + "\n\n" + m.list.View()
	v := tea.NewView(docStyle.Render(content))
	v.AltScreen = true
	return v
}
