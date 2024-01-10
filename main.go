package main

import (
	// "context"
	"context"
	"database/sql"
	"fmt"

	"github.com/charlesozo/whisperbot/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"google.golang.org/protobuf/proto"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type waConfig struct {
	Port      string
	ContactDB *database.ContactDB
}

func main() {
	godotenv.Load(".env")
	checkAndDeleteFile()
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("database environment variable is not set")
	}
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Connect to database", err)
	}
	defer conn.Close()
	contactdb, err := database.NewDB("database.json")

	if err != nil {
		log.Fatal(err)
	}
	wacfg := waConfig{
		Port:      dbURL,
		ContactDB: contactdb,
	}
	waclient, err := wacfg.waConnect()
	if err != nil {
		log.Panic(err)
	}
	waclient.WaitForConnection(10 * time.Second)

	fmt.Printf("Whatsapp is connected %v\n", waclient.IsConnected())
	fmt.Printf("User is loggedIn %v\n", waclient.IsLoggedIn())

	err = waclient.SendPresence("available")
	if err != nil {
		fmt.Printf("presence error: %v", err)
	}
	//    wacfg.GetAllContacts(waclient)
	myJID, err := wacfg.CreateWaUserId(waclient, []string{"+2347036179840"})
	if err != nil {
		fmt.Println("Cant create JID")
		os.Exit(0)
	}
	imageBytes, err := os.ReadFile("./images/2024.jpg")
	if err != nil {
		log.Fatalf("Error reading image file %v", err)
		return
	}
	resp, err := waclient.Upload(context.Background(), imageBytes, whatsmeow.MediaImage)
	// handle error
	if err != nil {
		log.Fatalf("Error uploading image %v", err)
		
	}

	imageMsg := &waProto.ImageMessage{
		Caption:  proto.String("Hello, world!"),
		Mimetype: proto.String("image/png"), // replace this with the actual mime type
		// you can also optionally add other fields like ContextInfo and JpegThumbnail here

		Url:           &resp.URL,
		DirectPath:    &resp.DirectPath,
		MediaKey:      resp.MediaKey,
		FileEncSha256: resp.FileEncSHA256,
		FileSha256:    resp.FileSHA256,
		FileLength:    &resp.FileLength,
	}

	response, err := waclient.SendMessage(context.Background(), myJID[0], &waProto.Message{
		// Conversation: proto.String("Hello charles"),
		ImageMessage: imageMsg,
	})
	if err != nil {
		fmt.Println("Error Sending Message")
		os.Exit(0)
	}
	fmt.Println(response.Timestamp)
	// os.Exit(1)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	waclient.Disconnect()

}
