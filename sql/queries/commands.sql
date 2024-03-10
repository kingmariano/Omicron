-- name: ExecuteCommand :exec
INSERT INTO COMMANDS(command_name,description) VALUES
('/generate-image','Enter text prompt to generate stunning and unique images. _ex:A peaceful sunset over a calm ocean, with vibrant colors reflecting in the water._'),
('/transcribe-audio', 'Upload the audio file you would like to transcribe or record sound using the whatsapp voice recorder. please note _audio music files are not supported. audio should be clear_'),
('/text2speech', 'Enter the text you want to convert to audio.'),
('/doc-interact', 'Upload document file ex _pdf, doc format_'),
('/download-video','Type video url _eg: video.youtube.com_ that you wish to download; supports videos from youtube, twitter, instagram, facebook, tiktok, videmeo_'),
('/video2audio','Upload video file you wish to convert to audio/mp3'),
('/search-song','Enter name of song or use the whatsapp voice recording feature to look up songs'),
('/find-location', 'Enter the phone number you want to lookup its current location. _eg: +1234567894_'),
('/verify', 'Enter your email address to begin verification process'),
('/news', 'Select the category of news update'),
('/help', 'Explore omnicron features by typing any of these commands\n👉*/verify*: Use this command to start the account verification process and gain access to additional features or services.\n👉*/generate-image*: Generate unique and creative images based on your prompt.\n👉*/generate-video*: Generate High Quality video based on your prompt.Transform your ideas to reality.\n👉*/transcribe-audio*: Transcribe your audio files with this command.\n👉*/text2speech*: Convert text into speech. Perfect for creating voice message, audio books or adding narration to your project.\n👉*/doc-interact*: This command helps you interact with your document, such as summarizing a pdf file,generating questons from the document etc.\n👉*/download-video*: Download videos from various sources _(eg. Youtube, Tiktok, Twitter, Instagram)_.\n👉*/video2audio*: Convert various video format to audio/mp3.\n👉*/search-song*: Use this command to search for audio music either by name or sound clip.\n👉*/find-location*:Get the current location of a device in real-time. _This command is for educational purpose and should be used with caution_.\n👉*/news*:Know whats happening around you and in the world with this command\n🔄 Conversational Mode: Chat with Omnicron like a friend! start typing, and Omnicron will respond in a friendly and engaging manner.');

