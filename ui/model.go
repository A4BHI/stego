package ui

import (
	"os"

	"charm.land/bubbles/v2/filepicker"
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type Model struct {
	list    list.Model
	Welcome string

	CoverPicker  filepicker.Model
	SecretPicker filepicker.Model

	step        int
	CoverImage  string
	SecretFile  string
	OutputImage string
}

func InitialModel() Model {
	l := NewMenu()
	cover := filepicker.New()
	cover.AllowedTypes = []string{".png"}
	cover.CurrentDirectory, _ = os.UserHomeDir()

	secret := filepicker.New()
	secret.AllowedTypes = []string{}
	secret.CurrentDirectory, _ = os.UserHomeDir()

	return Model{

		list:         l,
		Welcome:      "STEGO- Golang Based Stegnography Tool.",
		CoverPicker:  cover,
		SecretPicker: secret,
		step:         -1,
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
			case "Encode":
				m.step = 0
				NewM, cmd := UpdateEncode(m, msg)
				return NewM, cmd
			case "Decode":

			case "Exit":
				return m, tea.Quit

			}

		}

	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m Model) View() tea.View {
	switch m.step {
	case -1:
		content := m.Welcome + "\n\n" + m.list.View()
		v := tea.NewView(docStyle.Render(content))
		v.AltScreen = true
		return v
	case 0:
		v := tea.NewView("Select Cover Image:\n\n" + m.CoverPicker.View())
		v.AltScreen = true
		return v
	case 1:
		v := tea.NewView("Select Secret File:\n\n" + m.SecretPicker.View())
		v.AltScreen = true
		return v
	}
	return tea.NewView("Unknown Screen")
}
