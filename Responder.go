package gtb

type Responder struct {
	apiService *ApiClient
}

type Msg struct {
	Text string // (*optional)
	Type string // (*optional) "document" | "photo" | "audio" | "video"
	Path string // (*optional) Relative file path
}

// Reply by tagging the sender of the message
func (r *Responder) Reply(m Message, c Msg) {
	r.sendMessage(TextMessage{
		ChatID:           m.Chat.ID,
		ReplyToMessageID: m.MessageID,
		Text:             c.Text,
	}, c)
}

// Send message to the chat with the specified id (chatId)
func (r *Responder) SendToChat(chatId int, c Msg) {
	r.sendMessage(TextMessage{ChatID: chatId, Text: c.Text}, c)
}

func (r *Responder) sendMessage(textFields TextMessage, c Msg) {

	if c.Path == "" {
		r.apiService.SendText(textFields)
		return
	}

	typeFile := "document"

	if c.Type != "" {
		typeFile = c.Type
	}
	r.apiService.SendFile(typeFile, c.Path, textFields)

}
