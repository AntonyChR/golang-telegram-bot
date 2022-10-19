package bot

import (
	"os"
	"os/signal"
	"syscall"
)

type Bot struct {
	server      *server
	apiClient   *ApiClient
	commandDesc map[string]string
	ch          chan Message

	onMessageCb CallBack 
	callBacksWithRegexp []callBackWithRegexp
}

type callBackWithRegexp struct{
	cb CallBack 
	pattern string
}

type CallBack func(Message)

func (b *Bot) OnMessage(cb CallBack) {
	b.onMessageCb= cb
}

func (b *Bot) OnMessageWithRegexp(regExp string, cb CallBack) {
	b.callBacksWithRegexp = append(b.callBacksWithRegexp, callBackWithRegexp{cb, regExp,})
}

func (b *Bot) DescribeCommmands(desc map[string]string) {
	b.commandDesc = desc
}

func (b *Bot) Start() {

	b.apiClient.RegisterWebhook()

	if len(b.commandDesc) > 0 {
		b.apiClient.SetCommandsDescription(b.commandDesc)
	}

	// Handle termination signal (ctrl-c)
	sigTermChan := make(chan os.Signal)
	signal.Notify(sigTermChan, syscall.SIGINT, syscall.SIGTERM)
	go b.ListenCtrlCSignal(sigTermChan)

	go b.ListenIncommingMsg()

	b.server.Start()
}

func (b *Bot) ListenIncommingMsg() {
	for message := range b.ch {
		if b.onMessageCb!= nil {
			b.onMessageCb(message)
		}
	}
}

func (b *Bot) ListenCtrlCSignal(sigTermChan chan os.Signal) {
	<-sigTermChan
	b.apiClient.RemoveWebhook()
	os.Exit(0)
}
