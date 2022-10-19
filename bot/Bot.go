package bot

import (
	"os"
	"os/signal"
	"syscall"
)

type Bot struct {
	cmdCallbacks map[string]CmdCallback
	msgCallbacks map[string]CmdCallback
	server       *server
	apiClient    *ApiClient
	commandDesc  map[string]string
	ch           chan Message
}

type CmdCallback func(m Message)

// register command implementation
func (b *Bot) OnCommand(cmd string, callback CmdCallback) {
	b.cmdCallbacks[cmd] = callback
}
func (b *Bot) OnAnyTextMessage(callback func(Message)) {

}

func (b *Bot) OnMessageWithPattern(regExPattern string, callback func(Message)) {

}

func (b *Bot) DescribeCommmands(desc map[string]string) {
	b.commandDesc = desc
}

func (b *Bot) Start() {

	b.apiClient.RegisterWebhook()
	b.apiClient.SetCommandsDescription(b.commandDesc)

	// Handle termination signal (ctrl-c)
	sigTermChan := make(chan os.Signal)
	signal.Notify(sigTermChan, syscall.SIGINT, syscall.SIGTERM)

	go b.ListenEvents(sigTermChan)

	b.server.Start()
}

func (b *Bot) ListenEvents(sigTermChan chan os.Signal) {
	for {
		select {
		case <-sigTermChan:
			b.apiClient.RemoveWebhook()
			os.Exit(0)
		case message := <-b.ch:

			for k, b := range b.msgCallbacks {

			}
		}
	}
}
