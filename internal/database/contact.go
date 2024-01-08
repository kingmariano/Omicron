package database

import (
	"go.mau.fi/whatsmeow/types"
	"errors"
)

type Contact struct {
	ID int `json:"id"`
	FullName string `json:"name"`
	JID types.JID `json:"jid"`
}

func (db *ContactDB) CreateContact(contactInfo types.ContactInfo, jid types.JID) (Contact, error){
	dbStructure, err := db.loadDB()
	if err != nil{
		return Contact{}, err
	}
	if contactInfo.FullName == ""{
	     return Contact{}, errors.New("Contact not saved")
	}
	
	id := len(dbStructure.Contacts) + 1
	contact := Contact{
		ID: id,
		FullName: contactInfo.FullName,
		JID:  jid,
	}
	dbStructure.Contacts[id] = contact
	err = db.WriteDB(dbStructure)
	if err != nil {
		return Contact{}, err
	}
	return contact, nil
  
  
}