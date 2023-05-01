package gtb

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

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
func (r *Responder) Ban(chatId int, userId int, until string) error {

	minutesRegex := regexp.MustCompile(`[1-9]{1,}m`)
	hoursRegex := regexp.MustCompile(`[1-9]{1,}h`)
	dayRegex := regexp.MustCompile(`2[0-9]{3}-(0[1-9]|1[1-2])-(0[1-9]|[1-2][0-9]|3[0-1])`)

	var unixDate int64

	switch {
	case minutesRegex.MatchString(until):
		timeStr := strings.ReplaceAll("m", until, "")
		timeInt, _ := strconv.Atoi(timeStr)
		unixDate = time.Now().Add(time.Duration(timeInt) * time.Minute).Unix()

	case hoursRegex.MatchString(until):
		timeStr := strings.ReplaceAll("h", until, "")
		timeInt, _ := strconv.Atoi(timeStr)
		unixDate = time.Now().Add(time.Duration(timeInt) * time.Hour).Unix()

	case dayRegex.MatchString(until):
		t, _ := time.Parse("02/01/2006", until)
		unixDate = t.Unix()

	default:
		return errors.New("invalid time format")
	}

	userData := map[string]any{
		"chat_id":    chatId,
		"user_id":    userId,
		"until_date": unixDate,
	}
	err := r.apiService.banChatMember(userData)

	return err
}

func (r *Responder) Unban(chatId int, userId int) error {
	userData := map[string]int{
		"chat_id": chatId,
		"user_id": userId,
	}
	err := r.apiService.unbanChatMember(userData)
	return err
}

// Download and save file in dir with format:
//
// /{file_type}-{file_name}-{file_unique_id}.{ext}
func (r *Responder) DownloadFile(fileId string, dir string) error {
	file := r.apiService.GetFileInfo(fileId)

	pathWithoutSlash := strings.ReplaceAll(file.Result.FilePath, "/", "-")
	fileName := strings.ReplaceAll(pathWithoutSlash, ".", "-"+file.Result.FileUniqueID+".")

	if !file.Ok {
		return errors.New("is not posible get info about file")
	}

	err := r.apiService.downloadFile(file.Result.FilePath, dir+fileName)

	return err

}
