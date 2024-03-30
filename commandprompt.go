package main

import (
	"go.mau.fi/whatsmeow"
	// waProto "go.mau.fi/whatsmeow/binary/proto"
	"context"
	"go.mau.fi/whatsmeow/types"
)

func handleuserCommand(ctx context.Context, client *whatsmeow.Client, senderJID types.JID, command string) {
	// handle each command
	switch command {
	case "/generate-image":
		SendCommandInstruction(ctx, client, senderJID, GenerateImageInstruction)
		return

	case "/generate-video":
		SendCommandInstruction(ctx, client, senderJID, GenerateVideoInstruction)
		return

	case "/transcribe-audio":
		SendCommandInstruction(ctx, client, senderJID, TranscribeAudioInstruction)
		return

	case "/text2speech":
		SendCommandInstruction(ctx, client, senderJID, TextToSpeechInstruction)
		return

	case "/download-video_url":
		SendCommandInstruction(ctx, client, senderJID, DownloadVideoURLInstruction)
		return

	case "/video2audio":
		SendCommandInstruction(ctx, client, senderJID, VideoToAudioInstruction)
		return

	case "/download-song":
		SendCommandInstruction(ctx, client, senderJID, DownloadSongInstruction)
		return

	case "/download-movie":
		SendCommandInstruction(ctx, client, senderJID, DownloadMovieInstruction)
		return

	case "/dowload-apk":
		SendCommandInstruction(ctx, client, senderJID, DownloadAppInstruction)
		return

	case "/shazam":
		SendCommandInstruction(ctx, client, senderJID, ShazamInstruction)
		return

	case "/find-location":
		SendCommandInstruction(ctx, client, senderJID, FindLocationInstruction)
		return

	case "/verify":
		SendCommandInstruction(ctx, client, senderJID, VerifyInstruction)
		return

	case "/help":
		SendCommandInstruction(ctx, client, senderJID, HelpInstruction)
		return

	default:
		SendCommandInstruction(ctx, client, senderJID, AIMessage)

	}
	
}
