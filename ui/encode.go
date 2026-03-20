package ui

import tea "charm.land/bubbletea/v2"

func UpdateEncode(m Model, msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch m.step {
	case 0:
		m.CoverPicker, cmd = m.CoverPicker.Update(msg)

		if didselect, path := m.CoverPicker.DidSelectFile(msg); didselect {
			m.CoverImage = path
			m.step = 1

		}
	case 1:
		m.SecretPicker, cmd = m.SecretPicker.Update(msg)

		if didselect, path := m.SecretPicker.DidSelectFile(msg); didselect {
			m.SecretFile = path
			m.step = 2
		}
	}

	return m, cmd
}
