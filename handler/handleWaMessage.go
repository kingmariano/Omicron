package handler

import (
	"go.mau.fi/whatsmeow/types"
)

type WaClientMessage struct {
	SenderJID types.JID
	UsernameJID string
	MessageJID []types.MessageID
	ChatJID types.JID
    SenderNumber string
}