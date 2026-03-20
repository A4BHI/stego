package ui

import tea "charm.land/bubbletea/v2"

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

			m.screen = "#finalencodescreen"
		}
	}

	return m, cmd
}

func UpdateEnccodeScreen(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
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
		}
	}
	return m, nil
}
