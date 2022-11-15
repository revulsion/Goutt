package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

//TODO
type accountsModel []struct{}

type config struct {
	Server   string
	Username string
	Password string
}

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
		folders:  []string{"inbox", "sent", "junk"},
		selected: make(map[int]struct{}),
	}
	return m
}

//TODO
func (m folderModel) Init() tea.Cmd {
	return nil
}

//TODO
func (m folderModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	return nil, nil
}

//TODO
func (m folderModel) View() {
	//return nil
}

func parseConfig() config {
	// this should be $HOME/.config.json  (or even better XDG_HOME)
	content, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}

	// marshall me some json
	var conf config
	err = json.Unmarshal(content, &conf)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}

	return conf
}

//TODO
var iconn imapConnection

func main() {
	config := parseConfig()
	fmt.Println(config.Username)
	iconn.newImapClient(config)
	iconn.getMailboxes()
	iconn.logout()
}
