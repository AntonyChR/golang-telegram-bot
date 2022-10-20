package bot

type Responder struct {
	apiService *ApiClient
}

type Content struct {
	Text string
	Type string
	Path string
}

func (r *Responder) Reply(m Message, c Content) {

	data := TextMessage{
		ChatID:           m.Chat.ID,
		ReplyToMessageID: m.MessageID,
		Text:             c.Text,
	}

	if c.Path == "" {
		r.apiService.SendText(data)
		return
	}

	typeFile := "document"

	if c.Type != "" {
		typeFile = c.Type
	}
	r.apiService.SendFile(typeFile, c.Path, data)
}

func (r *Responder) SendToChat(chatId int, c Content) {
	r.apiService.SendText(TextMessage{ChatID: chatId, Text: c.Text})
}
