package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type accountsModel []struct{}

type folderModel struct {
	folders  []string
	selected map[int]struct{}
}

type messagesModel []struct {
	from    string
	subject string
	date    string // something date-time here?
	id      int
}

type messageModel struct {
	body         []string
	headers      struct{} // FIXME
	linePosition int
}

func newFolderModel() folderModel {
	m := folderModel{
		folders:  {"inbox", "sent", "junk"},
		selected: make(map[int]struct{}),
	}
}

func (m folderModel) Init() tea.Cmd {
	return nil
}

func (m folderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil
}

func (m folderModel) View() {
	return nil
}

func main() {
	fmt.Println("placeholder...")
}
