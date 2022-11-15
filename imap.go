package main

import (
	"log"

	"github.com/emersion/go-imap"
	"github.com/emersion/go-imap/client"
)

type imapConnection struct {
	client         *client.Client
	selectedFolder string
	err            error
}

func (ic *imapConnection) newImapClient(c config) error {
	ic.client, ic.err = client.DialTLS(c.Server, nil)

	if ic.err != nil {
		log.Fatal("whoops") // FIXME I cant log stdout with bubbletea
	}
	//defer c.Logout()

	if ic.err = ic.client.Login(c.Username, c.Password); ic.err != nil {
		log.Fatalf("whoops 2 %s", ic.err)
	}

	return ic.err
}

func (ic *imapConnection) getMailboxes() {

	if ic.client != nil {
		log.Println("not null2")
	}

	mailboxes := make(chan *imap.MailboxInfo, 10)
	done := make(chan error, 1)
	go func() {
		done <- ic.client.List("", "*", mailboxes)
	}()

	log.Println("Mailboxes:")
	for m := range mailboxes {
		log.Println("* " + m.Name)
	}
}

func (ic *imapConnection) getMailbox()         {}
func (ic *imapConnection) getMessage()         {}
func (ic *imapConnection) moveMessage()        {}
func (ic *imapConnection) deleteMessage()      {}
func (ic *imapConnection) updateMessageFlags() {}
func (ic *imapConnection) logout() {
	ic.client.Logout()
}
