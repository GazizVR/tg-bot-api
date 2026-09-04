# A lightweight, idiomatic Go wrapper for the **Telegram Bot API**.

[![Go Reference](https://pkg.go.dev/badge/github.com/gazizvr/tg-bot-api.svg)](https://pkg.go.dev/github.com/gazizvr/tg-bot-api)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](LICENSE)

---

## 🚀 Features

- **Long Polling Update Listener:** Concurrent update handling via Goroutines.
- **Message Operations:** Send, edit text, delete messages, and send documents/files.
- **Interactive Inline Keyboards:** Built-in support for creating inline markups, buttons, and answering callback queries.
- **Media & File Handling:** Direct file downloads by `file_id` and multipart document/media uploads.
- **Clean Architecture:** Simple client initialization with custom base API URLs (supports self-hosted Telegram Bot API instances).

---

## 📦 Installation

To install `tg-bot-api` in your Go project:

```bash
go get [github.com/gazizvr/tg-bot-api](https://github.com/gazizvr/tg-bot-api)

```

---

## 🛠️ Quick Start

Here is a quick example of setting up a Telegram bot, listening to updates, and sending responses with an inline keyboard.

```go
package main

import (
	"fmt"
	"log"

	tgbot "[github.com/gazizvr/tg-bot-api/pkg](https://github.com/gazizvr/tg-bot-api/pkg)"
)

const (
	botToken = "YOUR_TELEGRAM_BOT_TOKEN"
	apiURL   = "[https://api.telegram.org](https://api.telegram.org)"
)

func main() {
	// Initialize client
	client := tgbot.NewClient(botToken, apiURL)

	log.Println("Bot is running...")

	// Listen for incoming updates
	err := client.ListenUpdates(func(upd tgbot.Update) {
		// Handle incoming text message
		if upd.Message != nil {
			chatId := upd.Message.Chat.Id
			text := upd.Message.Text

			fmt.Printf("[%d] Message received: %s\\n", chatId, text)

			// Create Inline Keyboard Markup
			markup := tgbot.NewInlineMarkup(
				[]tgbot.InlineButton{
					{Text: "Option 1", Data: "opt_1"},
					{Text: "Option 2", Data: "opt_2"},
				},
			)

			// Send Message with Inline Markup
			_, err := client.SendMessage(chatId, "Hello! Choose an option:", markup, nil)
			if err != nil {
				log.Println("Error sending message:", err)
			}
		}

		// Handle callback queries from inline buttons
		if upd.Callback != nil {
			queryId := upd.Callback.Id
			chatId := upd.Callback.Message.Chat.Id
			msgId := upd.Callback.Message.Id

			// Acknowledge callback query
			_, _ = client.AnswerCallbackQuery(queryId)

			// Edit message text upon button click
			newText := fmt.Sprintf("You selected: %s", upd.Callback.Data)
			_, err := client.EditMessageText(chatId, msgId, newText, nil)
			if err != nil {
				log.Println("Error editing message:", err)
			}
		}
	}, []string{"message", "callback_query"})

	if err != nil {
		log.Fatalf("Fatal bot error: %v", err)
	}
}

```

---

## 📖 Key Usage Examples

### 1. Sending Messages & Replies

```go
// Simple text message
_, err := client.SendMessage(chatId, "Hello World!", nil, nil)

// Reply to a specific message
var replyToMsgId int64 = 1234
_, err = client.SendMessage(chatId, "This is a reply!", nil, &replyToMsgId)

```

### 2. File & Document Upload

```go
file, err := os.Open("report.pdf")
if err != nil {
    log.Fatal(err)
}
defer file.Close()

// Send document to chat
resp, err := client.SendDocument(chatId, *file, nil)

```

### 3. Downloading Files from Telegram

```go
// Download a file by Telegram file_id to a target directory
savedPath, err := client.DownloadFile("FILE_ID_HERE", "./downloads")
if err != nil {
    log.Println("Download failed:", err)
} else {
    fmt.Println("File downloaded to:", *savedPath)
}

```

### 4. Editing Messages & Keyboards

```go
// Edit message text
_, err := client.EditMessageText(chatId, messageId, "Updated text content", nil)

// Edit reply markup only
newMarkup := tgbot.NewInlineMarkup(
    []tgbot.InlineButton{{Text: "Refreshed Button", Data: "refresh"}},
)
_, err = client.EditMessageReplyMarkup(chatId, messageId, *newMarkup)

```

---

## 📚 API Reference

### `Client` Methods

| Method | Description |
| --- | --- |
| `NewClient(token, baseURL)` | Constructs a new Telegram API client instance. |
| `ListenUpdates(handler, allowed)` | Starts long polling and processes updates asynchronously. |
| `SendMessage(chatId, text, markup, replyTo)` | Sends a text message with optional inline keyboard and reply parameters. |
| `DeleteMessage(chatId, messageId)` | Deletes a message from a chat. |
| `SendDocument(chatId, file, replyTo)` | Uploads and sends a document/file to a chat. |
| `EditMessageText(chatId, messageId, text, markup)` | Edits text and inline keyboard of an existing message. |
| `EditMessageReplyMarkup(chatId, messageId, markup)` | Updates inline buttons on an existing message. |
| `EditMessageMedia(chatId, messageId, mediaType, file, markup)` | Replaces existing message media with a new file. |
| `AnswerCallbackQuery(queryId)` | Acknowledges a callback query sent from an inline keyboard button. |
| `DownloadFile(fileId, dirPath)` | Retrieves file metadata and downloads the file to local storage. |

---

## 📁 Repository Structure

```text
.
├── go.mod                  # Go module definition
├── LICENSE                 # MIT License
├── internal/
│   └── http.go             # HTTP helper functions (GET / Multipart POST)
└── pkg/
    ├── client.go           # Client struct & Update listening loop
    ├── dto.go              # API response structures & error checking
    ├── editMsgMethods.go   # Message editing API methods
    ├── errors.go           # Package errors
    ├── fileMethods.go     # File retrieval & download methods
    ├── methods.go          # Core updates & callback methods
    ├── msgMethods.go       # Message sending & deletion methods
    └── objects.go          # Telegram API data types (Update, Message, Keyboard, etc.)

```

---

## 📄 License

This project is licensed under the MIT License — see the [LICENSE](https://www.google.com/search?q=LICENSE) file for details.

**Author:** Gaziz (`@gazizvr`)
