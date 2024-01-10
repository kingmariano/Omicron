package main

import (
	"fmt"
	"os"
)

func checkAndDeleteFile() {
	if _, err := os.Stat("database.json"); err == nil {
		err := os.Remove("database.json")
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		fmt.Println("database.json deleted successfully")
	} else if os.IsNotExist(err) {
		fmt.Println("database.json does not exist")
	} else {
		fmt.Println("Error:", err)
	}
}
