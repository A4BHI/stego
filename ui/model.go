package ui

import (
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	screen string
}

// func InitialModel() Model {
// 	title := list.DefaultStyles(true).Title

// 	items := []list.Item{
// 		item{"Encode", "Hide data inside an image"},
// 		item{"Decode", "Extract hidden data"},
// 		item{"Exit", "Quit the program"},
// 	}

// 	m := Model{
// 		list:    list.New(items, list.NewDefaultDelegate(), 0, 70),
// 		Welcome: "STEGO-a stegnography tool in golang",
// 	}
// 	m.list.Styles.Title = title

// 	m.list.Title = "Choose an option from the list."

// 	return m
// }

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
