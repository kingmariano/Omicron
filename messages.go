package main

import (
	"strings"
)

var NewIntroductoryMessage string = "👋 Hello [User]! Welcome to Omnicron, your friendly and versatile WhatsApp bot! Omnicron can perform a wide range of operations, from image/video generation  to audio/video transcription and a whole lot more. You can start by chatting with omnicron like a friend or interact with it by typing commands starting with */'command_name'* .👉 _To see all available commands, type_  */help*. Let's get started! Feel free to ask Omnicron anything or explore its many features. Have fun chatting! 😊"
var MessageUser string = "hey nice to meet you again [User]!"

func FormatMessage(messge string, username string) string {
	if username == "" {
		username = "dear"
	}
	formattedMessage := strings.Replace(messge, "[User]", username, -1)
	return formattedMessage
}