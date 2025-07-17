package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"sort"
)

type HelpCommand struct {
	lister CommandLister
	writer io.Writer
}

func NewHelpCommand(lister CommandLister) *HelpCommand {
	return &HelpCommand{
		lister: lister,
		writer: os.Stdout,
	}
}

func (h *HelpCommand) Name() string {
	return "help"
}

func (h *HelpCommand) Description() string {
	return "Show help information about commands"
}

func (h *HelpCommand) Usage() string {
	return "kc help [command]"
}

func (h *HelpCommand) Execute(ctx context.Context, args []string) error {
	if len(args) > 0 {
		return h.showCommandHelp(args[0])
	}

	return h.showGeneralHelp()
}

func (h *HelpCommand) showGeneralHelp() error {
	fmt.Fprintf(h.writer, "kc - Kartones CLI\n\n")
	fmt.Fprintf(h.writer, "Usage:\n")
	fmt.Fprintf(h.writer, "kc <command> [arguments]\n\n")
	fmt.Fprintf(h.writer, "Available commands:\n")

	commands := h.lister.All()
	var names []string
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		cmd := commands[name]
		fmt.Fprintf(h.writer, "  %-15s %s\n", name, cmd.Description())
	}

	fmt.Fprintf(h.writer, "\nUse 'kc help <command>' for more information about a command.\n")

	return nil
}

func (h *HelpCommand) showCommandHelp(cmdName string) error {
	cmd, exists := h.lister.Get(cmdName)
	if !exists {
		return fmt.Errorf("unknown command: %s", cmdName)
	}

	fmt.Fprintf(h.writer, "Usage: %s\n\n", cmd.Usage())
	fmt.Fprintf(h.writer, "%s\n", cmd.Description())

	return nil
}
