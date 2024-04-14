package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/store/sqlstore"
	"go.mau.fi/whatsmeow/types"
	"go.mau.fi/whatsmeow/types/events"
	waLog "go.mau.fi/whatsmeow/util/log"
)

var (
	userChannels     map[string]userChannel
	userGoroutineMap sync.Map
)

type keystore string

const CTXErrMessage keystore = "errorMessage"

type userChannel struct {
	senderJID    types.JID
	username     string
	messageID    []types.MessageID
	chatJID      types.JID
	senderNumber string
	messageChan  chan *waProto.Message
}

func init() {
	userChannels = make(map[string]userChannel)
}

func GetEventHandler(client *whatsmeow.Client, cfg *waConfig, ctx context.Context) func(interface{}) {
	return func(evt interface{}) {
		switch v := evt.(type) {
		case *events.Message:
			if v.Info.Chat.Server == "s.whatsapp.net" {
				// mark read for any message received.
				go MarkMessageRead(client, v.Info.Chat, v.Info.Sender, []types.MessageID{v.Info.ID})
				userChan, ok := userChannels[v.Info.Sender.User]
				if !ok {
					// this is a new User, hasnt yet been registered in the database.
					userChan = userChannel{
						senderJID:    v.Info.Sender,
						username:     v.Info.PushName,
						messageID:    []types.MessageID{v.Info.ID},
						chatJID:      v.Info.Chat,
						senderNumber: v.Info.Sender.User,
						messageChan:  make(chan *waProto.Message, 10),
					}
					cfg.HandleNewUser(ctx, client, userChan.chatJID, userChan.senderJID, userChan.username, userChan.senderNumber)
					fmt.Println("GetConversation : ", v.Message.GetConversation())
					fmt.Println("Sender : ", userChan.senderJID)
					fmt.Println("Sender Number : ", userChan.senderNumber)
					// add the user to the map.
					userChannels[v.Info.Sender.User] = userChan

				} else {
					// This not first user, send the message into the client channel struct.
					userChan.messageChan <- v.Message
					_, alreadyRunning := userGoroutineMap.Load(userChan.senderNumber)
					if !alreadyRunning {
						// create a new goroutine
						go func() {
							// Mark that a goroutine is now running for this senderNumber
							userGoroutineMap.Store(userChan.senderNumber, true)
							defer userGoroutineMap.Delete(userChan.senderNumber) // Remove the entry when the goroutine exits
							handleClientIncomingMessage(client, ctx, userChan)
						}()
						log.Printf("created a new goroutine for %v", userChan.senderNumber)
					} else {
						log.Printf("goroutine is already running and doing work for %v", userChan.senderNumber)
					}

					}
			}

		}
	}
}

func handleClientIncomingMessage(client *whatsmeow.Client, ctx context.Context, user userChannel) {
   log.Printf("I am the new goroutine for %v", user.senderNumber)
	var msgLock sync.Mutex
	msgLock.Lock()
	defer msgLock.Unlock()
	for {
		select {
		case message := <-user.messageChan:
			msg := message.GetConversation()
			fmt.Println(msg)
			handleuserCommand(ctx, client, user.senderJID, msg)
		case <-ctx.Done():
			fmt.Println(ctx.Value(CTXErrMessage))
			return

		}
	}
}

func (cfg *waConfig) waConnect(ctx context.Context) (*whatsmeow.Client, error) {
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
	client.AddEventHandler(GetEventHandler(client, cfg, ctx))
	if client.Store.ID == nil {
		// No ID stored, new login
		qrChan, _ := client.GetQRChannel(ctx)
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
