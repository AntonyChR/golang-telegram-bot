package bot

type Bot struct {
	callBacks   map[string]Callback
	server      *Server
	telegramApi *TelegramApi
	commandDesc map[string]string
	ch          chan Command
}

type Callback func(c Command)

// register command implementation
func (b *Bot) On(cmd string, callback Callback) {
	b.callBacks[cmd] = callback
}

func (b *Bot) DescribeCommmands(desc map[string]string) {
	b.commandDesc = desc
}

func (b *Bot) Start() {

	b.telegramApi.registerWebhook()
	b.telegramApi.SetCommandsDescription(b.commandDesc)
	//b.telegramApi.removeWebhook()

	go b.ListenCommands()

	b.server.Start()
}

func (b *Bot) ListenCommands() {
	for command := range b.ch {
		if _, exists := b.callBacks[command.Name]; exists {
			b.callBacks[command.Name](command)
		}
	}
}
