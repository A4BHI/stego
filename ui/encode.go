package ui

import (
	tea "charm.land/bubbletea/v2"
)

func UpdateEncode(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+b":
			if m.screen == "#encode" && m.step == 1 {
				m.step = 0
				m.CoverImage = ""
				m.SecretFile = ""
			} else {
				m.screen = "#menu"
				m.step = -1
				m.CoverImage = ""
				m.SecretFile = ""
			}

			return m, nil
		}
	}

	var cmd tea.Cmd
	switch m.step {
	case 0:
		m.CoverPicker, cmd = m.CoverPicker.Update(msg)

		if didselect, path := m.CoverPicker.DidSelectFile(msg); didselect {
			m.CoverImage = path
			m.step = 1
			return m, m.SecretPicker.Init() //initialize the filepicker immediately

		}
	case 1:
		m.SecretPicker, cmd = m.SecretPicker.Update(msg)

		if didselect, path := m.SecretPicker.DidSelectFile(msg); didselect {
			m.SecretFile = path
			m.step = 2
			m.screen = "#finalencodescreen"
		}
	}

	return m, cmd
}

func UpdateEnccodeScreen(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyPressMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "ctrl+b":
			m.screen = "#encode"
			m.step = 1

			m.SecretFile = ""
			return m, nil
		case "enter":
			if m.FocusIndex == 0 {
				m.FocusIndex = 1
				m.TextInput1.Blur()
				return m, m.TextInput2.Focus()
			}

			m.TextInput2, cmd = m.TextInput2.Update(msg)
			return m, cmd

			// if m.TextInput1.Value() != " " {

			// }
			// m.OutputImage = m.TextInput1.Value()
			// cfg := config.Config{
			// 	InputImage:  m.CoverImage,
			// 	SecretFile:  m.SecretFile,
			// 	OutputImage: m.TextInput.Value(),
			// }

		}
	}

	if m.FocusIndex == 0 {
		m.TextInput1, cmd = m.TextInput1.Update(msg)

		return m, cmd
	} else {
		m.TextInput2, cmd = m.TextInput2.Update(msg)

		return m, cmd
	}

}
