package bot

type Responder struct {
	apiService *ApiClient
}

type Content struct {
	Text string
	Type string
	Data []byte
}

func (r *Responder) Reply(m Message, c Content) {
	if c.Type == "" || c.Type == "text" {
		data := TextMessage{
			ChatID:           m.Chat.ID,
			ReplyToMessageID: m.MessageID,
			Text:             c.Text,
		}
		r.apiService.SendText(data)
	}
}

func (r *Responder) SendToChat(chatId int, c Content) {
	r.apiService.SendText(TextMessage{ChatID: chatId, Text: c.Text})
}
