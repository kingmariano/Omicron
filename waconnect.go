package main

import (
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
	"go.mau.fi/whatsmeow/types/events"
	"log"
	"context"
	"fmt"
	"os"
)

func eventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		fmt.Println("Received a message!", v.Message.GetConversation())

	}
}

func (cfg *waConfig) waConnect() (*whatsmeow.Client, error){
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	container, err := sqlstore.New("postgres", cfg.Port, dbLog)
	if err != nil {
		log.Fatalf("Unable to create a database store %v", err)
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		log.Fatalf("Unable to create a device store %v", err)
	} 
	
	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)
	
		if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(context.Background())
		err = client.Connect()
		if err != nil {
			return nil, err
		}
		for evt := range qrChan {
			if evt.Event == "code" {
				qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
			} else if evt.Event == "authenticated"{
				fmt.Println("User is logged in!")
				os.Exit(0) // Exit the program with a status code of 0
			}else {
				fmt.Println("Login event:", evt.Event)
			}
		}
	} else {
		err := client.Connect()
		if err != nil {
			return nil, err
		}
	
	}
	return client, nil

}