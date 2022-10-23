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
	Photo     []Photo  `json:"photo"` //index 2 item is the highest resolution
	Video     Video    `json:"video"`
	Voice     Voice    `json:"voice"`
	Audio     Audio    `json:"audio"`
	Caption   string   `json:"caption"`
	Entities  []Entity `json:"entities"`
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
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type Video struct {
	Duration     int    `json:"duration"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
	MIMEType     string `json:"mime_type"`
	Thumb        Thumb  `json:"thumb"`
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
}

type Voice struct {
	Duration     int    `json:"duration"`
	MIMEType     string `json:"mime_type"`
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
}

type Audio struct {
	Duration     int    `json:"duration"`
	FileName     string `json:"file_name"`
	MIMEType     string `json:"mime_type"`
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
}

type Document struct {
	FileName     string `json:"file_name"`
	MIMEType     string `json:"mime_type"`
	FileID       string `json:"file_id"`
	Thumb        Thumb  `json:"thumb"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
}

type Thumb struct {
	FileID       string `json:"file_id"`
	FileUniqueID string `json:"file_unique_id"`
	FileSize     int    `json:"file_size"`
	Width        int    `json:"width"`
	Height       int    `json:"height"`
}

type Entity struct {
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	Type   string `json:"type"`
}

// Check if the message contains:
//
//	typeFile = photo | video | audio | voice | document
//
// if file type is not specified hasFile(""), all types are searched
func (m *Message) HasFile(typeFile string) bool {
	switch typeFile {
	case "photo":
		return len(m.Photo) > 0
	case "video":
		return m.Video.FileID != ""
	case "audio":
		return m.Audio.FileID != ""
	case "voice":
		return m.Voice.FileID != ""
	default:
		return false
	}
}
