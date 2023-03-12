# Golang Telegram API client

## Install 
```bash
go get https://github.com/AntonyChR/golang-telegram-bot
```

## Configuration

The struct gtb.Config has thre fields
- Port: port on which the server listens for requests
- ServerUrl: url of the server where the bot is running, you can use heroku or any host provider, you can also use Ngrok as shown in the example below
- BotToken:  visit this [link](https://core.telegram.org/bots/tutorial#obtain-your-bot-token) to get your api token through [Bot Father](https://t.me/botfather)

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
	
	// listens for the /info command and responds with a text
	myBot.OnCommand("/info", func(m gtb.Message) {
		responder.SendToChat(
			m.Chat.ID,
			gtb.Msg{
				Text: "Hello I'm a bot :),
			})
	})
	
	// lsten to any message
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

// Ban chat member
// until can be specified with the following formats
//
//	minutes: "1m", "2m",..."nm"
//	hours: "1h", "2h",..."nh"
//	until specific date: "dd/mm/yy"
//
// If user is banned for more than 366 days or less than 30 seconds from the
// current time they are considered to be banned forever.
// https://core.telegram.org/bots/api#banchatmember
func (r *Responder) Ban(chaId int, userId int, until string) error

// Example
myBot.OnMessage(func(m gtb.Message) {
	if m.Text == "forbidden phrase" {
		r.Reply(m, gtb.Msg{Text: "your message has prohibited content, you will be banned from the group."})
		r.Ban(m.Chat.ID, m.From.ID, "1m") // banned
	}
})

// Unban chat member
func (r *Responder) Unban(chaId int, userId int) error
```

## Types

[gtl.Message](https://github.com/AntonyChR/golang-telegram-bot/blob/main/entities.go#L10): Incoming message data
```golang
type Message struct {
	MessageID int      `json:"message_id"`
	From      From     `json:"from"`
	Chat      Chat     `json:"chat"`
	Date      int      `json:"date"`
	Text      string   `json:"text"`
	Photo     []Photo  `json:"photo"` //index 2 item is the highest resolution
	Video     Video    `json:"video"`
	Voice     Voice    `json:"voice"`
	Audio     Audio    `json:"audio"`
	Caption   string   `json:"caption"`
	Entities  []Entity `json:"entities"`

	NewChatMember  ChatMember `json:"new_chat_member"`
	LeftChatMember ChatMember `json:"left_chat_member"`
}
```

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
