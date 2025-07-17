package interfaces

import "context"

type Command interface {
	Name() string
	Description() string
	Usage() string
	Execute(ctx context.Context, args []string) error
}

// To enable the help command to list all commands, without explicitly accepting a CommandRegistry
type CommandLister interface {
	All() map[string]Command
	Get(name string) (Command, bool)
}
