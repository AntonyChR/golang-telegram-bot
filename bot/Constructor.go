package bot

import (
	"fmt"
	"regexp"
)

type Config struct {
	Port      string
	ServerUrl string
	BotToken  string
}

func NewTelegramBot(c Config) (*Bot, error) {

	if err := validateConfig(c); err != nil {
		return &Bot{}, err
	}

	commandTransport := make(chan Command)

	httpServer := &server{
		Port:            c.Port,
		IncomingCommand: commandTransport,
	}

	apiClient := &TelegramApi{
		BaseUrl:   "https://api.telegram.org/bot" + c.BotToken,
		EndPoints: map[string]string{},
		ServerUrl: c.ServerUrl,
	}

	bot := Bot{
		ch:          commandTransport,
		server:      httpServer,
		telegramApi: apiClient,
		callBacks:   make(map[string]Callback),
	}
	return &bot, nil
}

func validateConfig(c Config) error {
	isValidToken, _ := regexp.Match("[0-9]{8,10}:[a-zA-Z0-9_-]{35}", []byte(c.BotToken))
	fmt.Println(isValidToken)
	return nil
}
