package bot

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
	//b.telegramApi.removeWebhook()

	go b.ListenCommands()

	b.server.Start()
}

func (b *Bot) ListenCommands() {
	for command := range b.ch {
		if _, exists := b.cmdCallbacks[command.Name]; exists {
			b.cmdCallbacks[command.Name](command)
		}
	}
}
