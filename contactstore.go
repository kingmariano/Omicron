package main

import (
	"fmt"
	"go.mau.fi/whatsmeow"
	"os"

)

func (cfg *waConfig) GetAllContacts(client *whatsmeow.Client) error {
contactsInfo, err := client.Store.Contacts.GetAllContacts()
if err != nil {
	fmt.Printf("Unable to get all contacts %v", err)
	return err
}

for jid, contact := range contactsInfo {
	contact, err := cfg.ContactDB.CreateContact(contact, jid)
	if err != nil {
		fmt.Println(err)
		continue
	}
	if contact.ID == 2{
		fmt.Println("Reached max number of test users")
		os.Exit(0)
	}
	msgResponse, err := cfg.SendMessage(client, jid,contact)
	if err != nil{
		fmt.Printf("error sending message %v", err)
		continue
	}

	fmt.Println(msgResponse.Timestamp.UTC())

}
return nil
}
