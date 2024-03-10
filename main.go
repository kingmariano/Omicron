package main

import (
	// "context"
	"database/sql"
	"fmt"
	"github.com/charlesozo/whisperbot/internal/database"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
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

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	waclient.Disconnect()

}
