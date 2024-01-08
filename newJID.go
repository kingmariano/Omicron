package main

import (
	"log"

	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)
func  (cfg *waConfig) CreateWaUserId(client *whatsmeow.Client, phone []string) ([]types.JID, error){
	responses, err := client.IsOnWhatsApp(phone)
	var waJID []types.JID
		if err != nil {
		log.Fatal(err)
	}
	for _, r := range responses {
		  if r.IsIn {
			waJID = append(waJID, r.JID)
		  }
	}
	return waJID, nil
   
}