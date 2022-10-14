package bot

type Command struct {
	Name             string
	Params           []string
	ChatId           int
	ReplyToMessageId int
}
