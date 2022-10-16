package bot

import (
	"os"
	"os/signal"
	"syscall"
)

type Bot struct {
	cmdCallbacks map[string]CmdCallback
	server       *server
	apiClient    *ApiClient
	commandDesc  map[string]string
	ch           chan Command
}

type CmdCallback func(c Command)

// register command implementation
func (b *Bot) OnCommand(cmd string, callback CmdCallback) {
	b.cmdCallbacks[cmd] = callback
}
func (b *Bot) OnMessage(regExPattern string) {

}

func (b *Bot) DescribeCommmands(desc map[string]string) {
	b.commandDesc = desc
}

func (b *Bot) Start() {

	b.apiClient.RegisterWebhook()
	b.apiClient.SetCommandsDescription(b.commandDesc)

	// Handle termination signal (ctrl-c) 
	sysChan := make(chan os.Signal)	
	signal.Notify(sysChan, syscall.SIGINT, syscall.SIGTERM)

	go b.ListenIncommingData(sysChan)

	b.server.Start()
}

func (b *Bot) ListenIncommingData(sysChan chan os.Signal) {
	for{
		select{
		case <- sysChan:
			b.apiClient.RemoveWebhook()
			os.Exit(0)
		case command := <- b.ch:
			if _, exists := b.cmdCallbacks[command.Name]; exists {
				b.cmdCallbacks[command.Name](command)
			}
		}
	}
}
