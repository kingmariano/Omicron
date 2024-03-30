package main

var NewIntroductoryMessage string = "👋 Hello [User]! Welcome to Omnicron, your friendly and versatile WhatsApp bot! Omnicron can perform a wide range of operations, from image/video generation  to audio/video transcription and a whole lot more. You can start by chatting with omnicron like a friend or interact with it by typing commands starting with */'command_name'* .👉 _To see all available commands, type_  */help*. Let's get started! Feel free to ask Omnicron anything or explore its many features. Have fun chatting! 😊"

var MessageUser string = "hey nice to meet you again [User]!"

var NewUserMessages = []string{"👋 Hello [User]! I am Omnicron, your friendly and versatile WhatsApp bot! I can perform a wide range of operations, which includes image/video generation, audio/video transcription and a whole lot more. You can chat with me like a friend or use the command ex: */command_name* to perform other tasks .", "👉 _To see all available commands, type_  */help*. Let's get started! Feel free to ask Omnicron anything or explore its many features. Have fun chatting! 😊"}

var GenerateImageInstruction string = "Enter text prompt to generate stunning and unique images. _ex: A peaceful sunset over a calm ocean, with vibrant colors reflecting in the water._"

var GenerateVideoInstruction string = "Enter text prompt to generate videos with high quality graphics. _ex: Cinematic view of the nagasaki Hiroshima event._"

var TranscribeAudioInstruction string = "Upload audio file  to transcribe or record using the whatsapp voice recorder. Please note:  _audio music files are not supported and audio should be clear_."

var TextToSpeechInstruction string = "Enter the text you want to convert to audio."

var DownloadVideoURLInstruction string = "Paste or Type video url _eg: video.youtube.com_  you want to download; supports videos from youtube, twitter, instagram, facebook, tiktok, videmeo_."

var VideoToAudioInstruction string = "Upload video file you want to convert to audio/mp3"

var ConvertDocToDocInstruction string = "Upload the document you want to convert to another file format. _ex: img.png_ ."

var DownloadSongInstruction string = "Enter name of  song you want to download to device."

var DownloadMovieInstruction string = "Enter name of  movie you want to download to device."

var DownloadAppInstruction string = "Enter name of APK, desktop app, or any other file format you want to download."

var ShazamInstruction string = "Use the whatsapp voice recording feature to look up songs."

var FindLocationInstruction string = "Enter the phone number you want to lookup its current location. _ex: +1234567894_"

var VerifyInstruction string = "Enter your email address to begin verification process."

var AIMessage string = "I am chatgpt how can i help you."

var HelpInstruction string = "Explore omnicron features by using any of the following commands\n👉*/verify*: Use this command to start the account verification process and gain access to additional features or services.\n👉*/generate-image*: Generate unique and creative images based on your prompt.\n👉*/generate-video*: Generate High Quality video based on your prompt.Transform your ideas to reality.\n👉*/transcribe-audio*: Transcribe your audio files with this command.\n👉*/text2speech*: Convert text into speech. Perfect for creating voice message, audio books or adding narration to your project.\n👉*/doc-interact*: This command helps you interact with your document, such as summarizing a pdf file,generating questons from the document etc.\n👉*/download-video*: Download videos from various sources _(eg. Youtube, Tiktok, Twitter, Instagram)_.\n👉*/video2audio*: Convert various video format to audio/mp3.\n👉*/search-song*: Use this command to search for audio music either by name or sound clip.\n👉*/find-location*:Get the current location of a device in real-time. _This command is for educational purpose and should be used with caution_.\n👉*/news*:Know whats happening around you and in the world with this command\n🔄 Conversational Mode: Chat with Omnicron like a friend! start typing, and Omnicron will respond in a friendly and engaging manner."
