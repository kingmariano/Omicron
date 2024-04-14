package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"github.com/charlesozo/whisperbot/internal/database"
	_ "github.com/lib/pq"
)

type waConfig struct {
	DB    *database.Queries
	DBURL string
}

func main() {
	dbURL, err := loadEnv()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open("postgres", dbURL)

	if err != nil {
		log.Fatal("Connect to database", err)
	}
	defer conn.Close()
	queries := database.New(conn)

	wacfg := waConfig{
		DB:    queries,
		DBURL: dbURL,
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	waclient, err := wacfg.waConnect(ctx)
	if err != nil {
		log.Panic(err)
	}
	waclient.WaitForConnection(15 * time.Second)

	fmt.Printf("Whatsapp is connected %v\n", waclient.IsConnected())
	fmt.Printf("User is loggedIn %v\n", waclient.IsLoggedIn())
	err = waclient.SendPresence("available")
	if err != nil {
		log.Print(err)
	}
    err = waclient.SetStatusMessage("new about me")
	if err != nil {
		log.Fatal(err)
	}
	
	// userInput := make(chan string)

    // go func() {
    //     // Simulate user input
    //     time.Sleep(30 * time.Second)
    //    cancel()
    // }()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	select {
		case <-c:
			waclient.Disconnect()
		case <-ctx.Done():
			fmt.Printf("Context done: %v\n", ctx.Err())
			log.Fatal("context cancelled from child")
		
	}
	// <-c

	// waclient.Disconnect()

}
