package gtb

// Types of incoming telegram request
type Body struct {
	UpdateID int     `json:"update_id"`
	Message  Message `json:"message"`
}

// Message data
type Message struct {
	MessageID int      `json:"message_id"`
	From      From     `json:"from"`
	Chat      Chat     `json:"chat"`
	Date      int      `json:"date"`
	Text      string   `json:"text"`
	Photo     []Photo  `json:"photo"`
	Caption   string   `json:"caption"`
	Entities  []Entity `json:"entities"`

	HasImages bool
}

type Chat struct {
	ID        int    `json:"id"`
	FirstName string `json:"first_name"`
	Type      string `json:"type"`
}

type From struct {
	ID           int    `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LanguageCode string `json:"language_code"`
}

type Photo struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int64  `json:"file_size"`
	Width        int64  `json:"width"`
	Height       int64  `json:"height"`
}

type Entity struct {
	Offset int64  `json:"offset"`
	Length int64  `json:"length"`
	Type   string `json:"type"`
}
