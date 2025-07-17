package registry

import "github.com/kartones/kc/internal/interfaces"

type Command = interfaces.Command
type CommandLister = interfaces.CommandLister

type CommandRegistry struct {
	commands map[string]Command
}

func NewRegistry() *CommandRegistry {
	return &CommandRegistry{
		commands: make(map[string]Command),
	}
}

func (r *CommandRegistry) Register(cmd Command) {
	r.commands[cmd.Name()] = cmd
}

func (r *CommandRegistry) Get(name string) (Command, bool) {
	cmd, exists := r.commands[name]
	return cmd, exists
}

func (r *CommandRegistry) All() map[string]Command {
	return r.commands
}
