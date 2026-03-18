package ui

import (
	"charm.land/bubbles/v2/list"
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

func InitialModel() Model {
	title := list.DefaultStyles(true).Title

	items := []list.Item{
		item{"Encode", "Hide data inside an image"},
		item{"Decode", "Extract hidden data"},
		item{"Exit", "Quit the program"},
	}

	m := Model{
		list:    list.New(items, list.NewDefaultDelegate(), 0, 70),
		Welcome: "STEGO-a stegnography tool in golang",
	}
	m.list.Styles.Title = title

	m.list.Title = "Choose an option from the list."

	return m
}
