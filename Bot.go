package bot

import (
	"os"
	"os/signal"
	"strings"
	"syscall"
)

type Bot struct {
	server       *server
	apiClient    *ApiClient
	commandDesc  map[string]string
	msgTransport chan Message
	commands     map[string]CallBack
	onMessageCb  CallBack
}

type CallBack func(Message)

func (b *Bot) DescribeCommmands(desc map[string]string) {
	b.commandDesc = desc
}

func (b *Bot) OnMessage(cb CallBack) {
	b.onMessageCb = cb
}

func (b *Bot) OnCommand(cmd string, cb CallBack) {
	b.commands[cmd] = cb
}

// Register webhook and initilize http server to
// listen incomming messages.
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

		if len(b.commands) == 0 {
			continue
		}

		text := strings.Split(message.Text, " ")

		if text[0] == "" {
			continue
		}

		cmd := text[0]

		if _, exists := b.commands[cmd]; exists {
			b.commands[cmd](message)
		}

	}
}

func (b *Bot) listenCtrlCSignal(sigTermChan chan os.Signal) {
	<-sigTermChan
	b.apiClient.RemoveWebhook()
	os.Exit(0)
}
