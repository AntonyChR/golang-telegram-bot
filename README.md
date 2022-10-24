# Golang Telegram API client

## Install 
```bash
go get https://github.com/AntonyChR/golang-telegram-bot
```

## Example

```golang
package main

import (
	"fmt"

	gtb "github.com/AntonyChR/golang-telegram-bot"
)

func main() {

	config := gtb.Config{
		Port:      "3000",
		ServerUrl: "https://09c2-38-25-23-224.sa.ngrok.io",
		BotToken:  "12345678:AAHrCtRRJbHML0iaO0rFjZTQN3kdSRM56qd9",
	}

	myBot, responder, _ := gtb.New(config)

	myBot.DescribeCommmands(map[string]string{
		"info": "Get info about bot",
	})

	myBot.OnCommand("/info", func(m gtb.Message) {
		responder.SendToChat(
			m.Chat.ID,
			gtb.Msg{
				Text: "Name: Bot\nVersion: 1.2.1\nActive: True",
			})
	})

	myBot.OnMessage(func(m gtb.Message) {
		fmt.Println(m.Text)
		responder.Reply(m, gtb.Msg{Text: "This is a text message"})
	})

	myBot.Start()

}

```

## gtb.Bot methods


```golang
// Set Commands description
func (b *Bot) DescribeCommands(map[string]string)

//Listen any message
func (b *Bot) OnMessage(func(m gtb.Message))

//Listen command
func(b *Bot) OnCommand(cmd string, func(m gtb.Message))

func (b *Bot) OnLeftMemberChat(func(m gtb.Message)

// Example
	myBot.OnLeftMemberChat(func(m gtb.Message) {
		fmt.Println(m.LeftChatMember.FirstName, " has left the chat")

	})

func (b *Bot) OnNewMemberChat(func(m gtb.Message)

// Example
	myBot.OnNewMemberChat(func(m gtb.Message) {
		fmt.Println(m.NewChatMember.FirstName, " has joined the chat.")
	})
```
## gtb.Responder methods


```golang

// Reply message
func (r *Responder) Reply(m bot.Message, c bot.Msg)

//Send message to specific chat
func (r *Responder) SendToChat(chatId int, c bot.Msg)

// Download and save file in:
// dir/{file_type}-{file_name}-{file_unique_id}.{ext}
func (r *Responder) DownloadFile(fileId string, dir string) error

// Example
	myBot.OnMessage(func(m gtb.Message) {
		err := r.DownloadFile(m.Photo[2].FileID, "./images/")
		if err != nil {
			fmt.Println("Error downloading file")
		}
	})

```

## Types

[gtl.Message](https://github.com/AntonyChR/golang-telegram-bot/blob/main/entities.go#L10): Incoming message data

[gtl.Msg](https://github.com/AntonyChR/golang-telegram-bot/blob/main/Responder.go#L7): Data to send message
```golang

type Msg struct {
	Text string // (*optional)
	Type string // (*optional) "document" | "photo" | "audio" | "video"
	Path string // (*optional) Relative file path: "./file.ext"
}
// If the type of file is not specified, it will be sent as a document

// example: send image
responder.SendToChat(
    m.Chat.ID,
    gtb.Msg{
        Text: "cat",
        Type: "photo",
        Path: "../images/cat.jpg"

    })

```
