package ui

import (
	"github.com/charmbracelet/bubbles/list"
)

type item string

func (i item) FilterValue() string { return string(i) }

type Model struct {
	list list.Model
}

func InitialModel() Model {
	title := list.DefaultStyles().Title
	items := []list.Item{
		item("Encode"),
		item("Decode"),
		item("Exit"),
	}

	m := Model{
		list: list.New(items, list.DefaultDelegate{}, 0, 0),
	}

	m.list.Title = "Choose an option from the list."

}
