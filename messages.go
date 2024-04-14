package main



var NewUserMessages = []string{"👋 Hello [User]! I am Omnicron, your friendly and versatile WhatsAppAI Bot. I can perform a wide range of operations, which includes image/video generation, audio/video transcription and more. You can start by chatting with me like a  friend or use the command (ex: */command_name*) to perform other tasks .", "👉To see all available commands, type */help*. Feel free to ask me anything. Have fun chatting! 😊"}

var GenerateImageInstruction string = "Enter text prompt to generate stunning and unique images. (ex: *A peaceful sunset over a calm ocean, with vibrant colors reflecting in the water.*)"

var GenerateVideoInstruction string = "Enter text prompt to generate videos with high quality graphics. ex: (*Cinematic view of the nagasaki Hiroshima event.*)"

var TranscribeAudioInstruction string = "Upload audio file to transcribe or record voice using the whatsapp voice recorder. Please note:  *audio music files are not supported and audio should be clear*."

var TextToSpeechInstruction string = "Enter the text you want to convert to audio."

var DownloadVideoURLInstruction string = "Paste or Type video url _eg: video.youtube.com_  you want to download; supports videos from youtube, twitter, instagram, facebook, tiktok, videmeo_."

var VideoToAudioInstruction string = "Upload video file you want to convert to audio/mp3"

var TranslateAudioInstruction string = "Upload audio file you want to convert to translate (_whatsapp voice note is also accepted_)"

var ConvertDocToDocInstruction string = "Upload the document you want to convert to another file format. _ex: img.png_ ."

var CompressFileInstruction string = "Upload the file you want to compress. _ex: img.png_."

var DocInteractInstruction string = "Upload the document you want to interact with"

var YoutubeSummarizeInstruction string = "Paste the youtube video URL you want to summarize"

var DownloadSongInstruction string = "Enter name of song you want to download to device."

var DownloadMovieInstruction string = "Enter name of movie you want to download to device."

var DownloadAppInstruction string = "Enter name of APK, desktop app, or any other file format you want to download."

var SearchSongInstruction string = "Enter name or record the song you want to search for using the whatsapp voice recorder (_audio must be clear and above 10 seconds_)"

var ShazamInstruction string = "Use the whatsapp voice recording feature to look up songs."

var FindLocationInstruction string = "Enter the phone number you want to lookup its current location. (ex: *+1234567894*)"

var VerifyInstruction string = "Enter your email address to begin verification process."

var AIMessage string = "I am chatgpt how can i help you."

var HelpInstruction string = "Explore omnicron features by using any of the following commands\n👉 */verify*: Use this command to start the account verification process and gain access to additional features or services.\n👉 */generate-image*: Generate unique and creative images based on your prompt.\n👉 */generate-video*: Generate High Quality video based on your prompt.Transform your ideas to reality.\n👉 */transcribe-audio*: Transcribe your audio files with this command.\n👉 */text2speech*: Convert text into speech. Perfect for creating voice message, audio books or adding narration to your project.\n👉 */compress-file*: Compress file to desired output.\n👉 */youtube-summarize*: Summarize Youtube video with AI.\n👉 */doc-interact*: This command helps you interact with your document, such as summarizing a pdf file,generating questons from the document etc.\n👉 */download-song*: Use this command to download audio music file to device.\n👉 */download-movie*: Use this command to generate download link for movie.\n👉 */download-apk*: Use this command to generate download link for apk's and desktop apps.\n👉 */compress-file*: Compress your document/file with this command.\n👉 */convert-file*: Convert to your document/file to any other format using this command.\n👉  */translate-audio*: Translate your audio to desired language using this command.\n👉 */download-video_url*: Download videos from various sources _(eg. Youtube, Tiktok, Twitter, Instagram)_.\n👉 */video2audio*: Convert various video format to audio/mp3.\n👉 */shazam*: Use this command to search for audio music either by name or sound clip.\n👉 */find-location*: Get the current location of a device in real-time. _This command is for educational purpose and should be used with caution_.\n👉 */news*: Know whats happening around you and in the world with this command\n🔄 *Conversational Mode*: Chat with Omnicron like a friend! start typing, and Omnicron will respond in a friendly and engaging manner."
