package main

import (
	"context"
	"fmt"
	"strings"
    "google.golang.org/protobuf/proto"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"github.com/charlesozo/whisperbot/internal/database"
	"go.mau.fi/whatsmeow/types"
)

func (cfg *waConfig) SendMessage(client *whatsmeow.Client, jid types.JID, contact database.Contact ) (whatsmeow.SendResponse, error){

	
	fineName := fineTuneName(contact.FullName)
	if fineName == ""{
		fineName = "Dear"
	}
	waMessage := fmt.Sprintf("Happy Christmas %v wishing you the best", fineName)
	sendResponse , err := client.SendMessage(context.Background(), jid, &waProto.Message{
		Conversation: proto.String(waMessage),
		
	})
	if err != nil {
		fmt.Printf("Unable to send Message %v", err)
		return whatsmeow.SendResponse{}, err
    }
	return sendResponse, err
      
}

func fineTuneName(name string) string {
	lowerName := strings.ToLower(name)

	if strings.Contains(lowerName, "dad") {
		return "Mr Charles"
	} else if strings.Contains(lowerName, "mom") {
		return "Mrs Ogoamaka"
	} else if strings.Contains(lowerName, "mr") {
		words := strings.Fields(name)
		if len(words) > 1 {
			return "Mr " + words[1]
		}
	} else if strings.Contains(lowerName, "aunty") {
		words := strings.Fields(name)
		if len(words) > 1 {
			return "Mrs " + words[1]
		}
	} else if strings.Contains(lowerName, "uncle") {
		words := strings.Fields(name)
		if len(words) > 1 {
			return "Mr " + words[1]
		}
	}

	// For any other format, separate the string and return the first word
	words := strings.Fields(name)
	if len(words) > 0 {
		return words[0]
	}

	return ""
}