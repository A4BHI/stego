package ui

import (
	"charm.land/bubbles/v2/list"
	tea "charm.land/bubbletea/v2"
)

type item string

func (i item) FilterValue() string { return string(i) }

type Model struct {
	Welcome string
	list    list.Model
}

func InitialModel() Model {
	title := list.DefaultStyles(true).Title

	items := []list.Item{
		item("Encode"),
		item("Decode"),
		item("Exit"),
	}

	m := Model{
		list:    list.New(items, list.DefaultDelegate{}, 0, 0),
		Welcome: "STEGO-a stegnography tool in golang",
	}
	m.list.Styles.Title = title

	m.list.Title = "Choose an option from the list."

	return m
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
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
	return tea.NewView(m.list.View())
}
