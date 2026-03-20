package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var docStyle = lipgloss.NewStyle().Margin(1, 2)

type item struct {
	title       string
	description string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.description }
func (i item) FilterValue() string { return i.title }

func NewMenu() list.Model {

	title := list.DefaultStyles(true).Title

	items := []list.Item{
		item{"Encode", "Hide data inside an image"},
		item{"Decode", "Extract hidden data"},
		item{"Exit", "Quit the program"},
	}

	l := list.New(items, list.NewDefaultDelegate(), 0, 70)

	l.Styles.Title = title

	l.Title = "Choose an option from the list."

	return l
}

func UpdateMenu(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {

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
				m.screen = "#encode"

				// cover := filepicker.New()
				// cover.AllowedTypes = []string{}
				// cover.CurrentDirectory, _ = os.UserHomeDir()

				// m.CoverPicker = cover
				return m, m.CoverPicker.Init()
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
