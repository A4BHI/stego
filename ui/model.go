package ui

import (
	"os"

	"charm.land/bubbles/v2/filepicker"
	"charm.land/bubbles/v2/list"
	"charm.land/bubbles/v2/textinput"
	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

var (
	bdr = lipgloss.NewStyle().BorderStyle(myCuteBorder).BorderForeground(lipgloss.Color("63")).Bold(true).Foreground(lipgloss.Color("#84cb94"))

	myCuteBorder = lipgloss.Border{
		Top:         "._.:*:",
		Bottom:      "._.:*:",
		Left:        "|*",
		Right:       "|*",
		TopLeft:     "*",
		TopRight:    "*",
		BottomLeft:  "*",
		BottomRight: "*",
	}

	textstyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA"))

	appStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("226")).
			Background(lipgloss.Color("63")).
			PaddingTop(2).
			PaddingLeft(4)
)

type Model struct {
	screen  string
	list    list.Model
	Welcome string

	CoverPicker  filepicker.Model
	SecretPicker filepicker.Model
	TextInput1   textinput.Model
	TextInput2   textinput.Model
	width        int
	height       int
	FocusIndex   int
	step         int
	CoverImage   string
	SecretFile   string
	OutputImage  string
}

func InitialModel() Model {
	l := NewMenu()
	cover := filepicker.New()

	cover.AllowedTypes = []string{".png"}
	cover.CurrentDirectory, _ = os.UserHomeDir()

	secret := filepicker.New()

	secret.AllowedTypes = []string{}

	secret.CurrentDirectory, _ = os.UserHomeDir()

	t1 := textinput.New()
	t1.Placeholder = "Output image name "
	t1.CharLimit = 50
	t1.SetVirtualCursor(false)
	t1.Focus()
	t1.SetWidth(20)

	t2 := textinput.New()
	t2.Placeholder = "Encryption Password"
	t2.CharLimit = 50
	t2.SetVirtualCursor(false)
	// t2.Focus()
	t2.SetWidth(20)

	return Model{
		screen:       "#menu",
		list:         l,
		Welcome:      "STEGO- Golang Based Stegnography Tool.",
		CoverPicker:  cover,
		SecretPicker: secret,
		TextInput1:   t1,
		TextInput2:   t2,
		step:         -1,
	}
}

func (m Model) Init() tea.Cmd {
	return tea.Batch(
		m.CoverPicker.Init(),
		m.SecretPicker.Init(),
		textinput.Blink,
	)

}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	if msg, ok := msg.(tea.WindowSizeMsg); ok {
		m.width = msg.Width
		m.height = msg.Height
		h, v := docStyle.GetFrameSize()
		m.list.SetSize(msg.Width-h, msg.Height-v-2)

		m.CoverPicker, _ = m.CoverPicker.Update(msg)
		m.SecretPicker, _ = m.SecretPicker.Update(msg)
		m.TextInput1, _ = m.TextInput1.Update(msg)
		m.TextInput2, _ = m.TextInput2.Update(msg)
	} //Adjust the window size

	switch m.screen {
	case "#menu":
		return UpdateMenu(m, msg)
	case "#encode":
		return UpdateEncode(m, msg)
	case "#finalencodescreen":
		return UpdateEnccodeScreen(m, msg)

	}
	return m, nil

}

func (m Model) View() tea.View {
	switch m.screen {
	case "#menu":
		content := m.Welcome + "\n\n" + m.list.View()
		v := tea.NewView(docStyle.Render(content))
		v.AltScreen = true
		return v
	case "#encode":
		switch m.step {

		case 0:
			v := tea.NewView(textstyle.Render("Select Cover Image:") + "\n" + m.CoverPicker.View())

			v.AltScreen = true
			return v
		case 1:
			v := tea.NewView(textstyle.Render("Select Secret File:") + "\n" + m.SecretPicker.View())
			v.AltScreen = true

			return v

		}
	case "#finalencodescreen":

		header := "Start Encoding\n\n" +
			"Cover Image: " + m.CoverImage + "\n" +
			"Secret File: " + m.SecretFile + "\n\n" +
			m.headerView()

		str := header + "\n" + m.TextInput1.View() + "\nEnter Encryption Password\n" + m.TextInput2.View() + m.footerView()
		str2 := lipgloss.Place(m.width, m.height, lipgloss.Center, lipgloss.Center, bdr.Render(str))
		v := tea.NewView("TESTING HEADING " + str2)
		if m.FocusIndex == 0 {

			v.Cursor = m.TextInput1.Cursor()
			v.Cursor.Y = lipgloss.Height(header)
			v.Cursor.X = m.TextInput1.Cursor().X

		} else {
			v.Cursor = m.TextInput2.Cursor()
			v.Cursor.Y = lipgloss.Height(header) + 2
			v.Cursor.X = m.TextInput2.Cursor().X
		}

		// if m.step == 2 {
		// 	v := tea.NewView("\nEnter Encryption Password\n" + m.TextInput2.View())
		// 	if !m.TextInput2.VirtualCursor() {

		// 	}
		// }

		v.AltScreen = true
		v.BackgroundColor = lipgloss.Red
		return v
	}

	return tea.NewView("Unknown Screen")
}
func (m Model) headerView() string {
	return `Output Name:
Enter destination name (e.g. encoded_image) [Extension .png auto-appended]`
}
func (m Model) footerView() string { return "\n(ctrl+c to quit)\n(ctrl+b to go back)" }
