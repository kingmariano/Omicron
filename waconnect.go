package main

import (
	"context"
	"fmt"
	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
	"log"
	"os"
	"sync"
)



var senderJIDChan = make(chan types.JID,10)
var usernameChan = make(chan string,10)
var messageIDChan = make(chan []types.MessageID,10)
var chatJIDChan = make(chan types.JID, 10)
var senderNumberChan = make(chan string, 10)
var messageChan = make(chan *waProto.Message,10)

func eventHandler(evt interface{}) {
	var (
		usersLock sync.Mutex
		wg        sync.WaitGroup // Mutex for concurrent map access
	)
	switch v := evt.(type) {
	case *events.Message:
		if v.Info.Chat.Server == "s.whatsapp.net" {
			usersLock.Lock()
			defer usersLock.Unlock()
			wg.Add(1)
			go func(senderJID types.JID, username string, messageID []types.MessageID, chatJID types.JID, senderNumber string, message *waProto.Message) {
				defer wg.Done()
				senderJIDChan <- senderJID
				usernameChan <- username
				messageIDChan <- messageID
				chatJIDChan <- chatJID
				senderNumberChan <- senderNumber
				messageChan <- message
			}(v.Info.Sender, v.Info.PushName, []types.MessageID{v.Info.ID}, v.Info.Chat, v.Info.Sender.User, v.Message)
			fmt.Println("GetConversation : ", v.Message.GetConversation())
			fmt.Println("Sender : ", v.Info.Sender)
			fmt.Println("Sender Number : ", v.Info.Sender.User)
			fmt.Println("IsGroup : ", v.Info.IsGroup)
			fmt.Println("MessageSource : ", v.Info.MessageSource)
			fmt.Println("ID : ", v.Info.ID)
			fmt.Println("PushName : ", v.Info.PushName)
			fmt.Println("BroadcastListOwner : ", v.Info.BroadcastListOwner)
			fmt.Println("Category : ", v.Info.Category)
			fmt.Println("Chat : ", v.Info.Chat)
			fmt.Println("DeviceSentMeta : ", v.Info.DeviceSentMeta)
			fmt.Println("IsFromMe : ", v.Info.IsFromMe)
			fmt.Println("MediaType : ", v.Info.MediaType)
			fmt.Println("Multicast : ", v.Info.Multicast)
			fmt.Println("Info.Chat.Server : ", v.Info.Chat.Server)
			// senderJIDChan <- v.Info.Sender
			// usernameChan <- v.Info.PushName
			// messageIDChan <- []types.MessageID{v.Info.ID}
			// chatJIDChan <- v.Info.Chat
			// senderNumberChan <- v.Info.Sender.User
			// fmt.Println("channel passed in")
		}

	}
}
func (cfg *waConfig) handleIncomingMessages(client *whatsmeow.Client) {

	for {
        // Receive incoming messages from channels
        // For example:
        senderJID := <-senderJIDChan
        username := <-usernameChan
        chatJID := <-chatJIDChan
        senderNumber := <-senderNumberChan
        messageID := <-messageIDChan
        

        // Create a context for each message
        ctx, cancel := context.WithCancel(context.Background())
        done := make(chan struct{})

        // Start a goroutine to handle users concurrently
        go cfg.HandleUsers(ctx, client, senderJID, username, chatJID, senderNumber, messageID, done)

        // Use a select statement to wait for either the context to be canceled
        // or the message to be processed
        select {
        case <-ctx.Done():
            // Context canceled, stop the user processing
            cancel()
        case <-done:
            // User processed, continue with the next message
        }
    }
}

func (cfg *waConfig) waConnect() (*whatsmeow.Client, error) {
	dbLog := waLog.Stdout("Database", "DEBUG", true)
	clientLog := waLog.Stdout("Client", "DEBUG", true)
	container, err := sqlstore.New("postgres", cfg.DBURL, dbLog)
	if err != nil {
		log.Fatalf("Unable to create a database store %v", err)
	}
	deviceStore, err := container.GetFirstDevice()
	if err != nil {
		log.Fatalf("Unable to create a device store %v", err)
	}

	client := whatsmeow.NewClient(deviceStore, clientLog)
	client.AddEventHandler(eventHandler)
	go cfg.handleIncomingMessages(client)
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
			} else if evt.Event == "authenticated" {
				fmt.Println("User is logged in!")
				os.Exit(0) // Exit the program with a status code of 0
			} else {
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
