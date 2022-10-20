package bot

import (
	"os"
	"os/signal"
	"regexp"
	"syscall"
)

type Bot struct {
	server       *server
	apiClient    *ApiClient
	commandDesc  map[string]string
	msgTransport chan Message

	onMessageCb         CallBack
	callBacksWithRegexp []callBackWithRegexp
}

type callBackWithRegexp struct {
	cb      CallBack
	pattern string
}

type CallBack func(Message)

func (b *Bot) OnMessage(cb CallBack) {
	b.onMessageCb = cb
}

func (b *Bot) OnMessageWithRegexp(regExp string, cb CallBack) {
	b.callBacksWithRegexp = append(b.callBacksWithRegexp, callBackWithRegexp{cb, regExp})
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
	go b.listenCtrlCSignal(sigTermChan)

	go b.listenIncommingMsg()

	b.server.Start()
}

func (b *Bot) listenIncommingMsg() {
	for message := range b.msgTransport {

		if b.onMessageCb != nil {
			b.onMessageCb(message)
		}

		if len(b.callBacksWithRegexp) > 0 {
			for _, cb := range b.callBacksWithRegexp {
				isMatch, _ := regexp.Match(cb.pattern, []byte(message.Text))
				if isMatch {
					continue
				}
				cb.cb(message)
			}
		}
	}
}

func (b *Bot) listenCtrlCSignal(sigTermChan chan os.Signal) {
	<-sigTermChan
	b.apiClient.RemoveWebhook()
	os.Exit(0)
}
