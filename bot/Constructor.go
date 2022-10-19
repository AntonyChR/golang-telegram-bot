package bot

import (
	"errors"
	"regexp"
)

type Config struct {
	Port      string
	ServerUrl string
	BotToken  string
}

func New(c Config) (*Bot,*Responder, error) {

	if err := validateConfig(c); err != nil {
		return &Bot{}, &Responder{},nil
	}

	commandTransport := make(chan Message)

	httpServer := &server{
		Port:            c.Port,
		IncomingCommand: commandTransport,
	}

	apiClient := &ApiClient{
		BaseUrl:   "https://api.telegram.org/bot" + c.BotToken,
		EndPoints: endpoints(),
		ServerUrl: c.ServerUrl,
	}

	responder := &Responder{
		apiService: apiClient,
	}

	bot := &Bot{
		ch:           commandTransport,
		server:       httpServer,
		apiClient:    apiClient,
		cmdCallbacks: make(map[string]CmdCallback),
	}
	return bot,responder,nil 
}

func validateConfig(c Config) error {
	isValidToken, _ := regexp.Match("[0-9]{8,10}:[a-zA-Z0-9_-]{35}", []byte(c.BotToken))
	if isValidToken {
		return nil
	}
	return errors.New("[x] Invalid token")
}
