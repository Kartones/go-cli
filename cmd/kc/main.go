package main

import (
	"context"
	"fmt"
	"os"

	"github.com/kartones/kc/internal/command/commands"
	"github.com/kartones/kc/internal/command/registry"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func run() error {
	reg := registry.NewRegistry()

	reg.Register(commands.NewReadConfigCommand())
	reg.Register(commands.NewListDirCommand())

	helpCmd := commands.NewHelpCommand(reg)
	reg.Register(helpCmd)

	args := os.Args[1:]
	if len(args) == 0 {
		return helpCmd.Execute(context.Background(), nil)
	}

	cmdName := args[0]
	cmdArgs := args[1:]

	cmd, exists := reg.Get(cmdName)
	if !exists {
		fmt.Fprintf(os.Stderr, "Unknown command: %s\n\n", cmdName)
		return helpCmd.Execute(context.Background(), nil)
	}

	return cmd.Execute(context.Background(), cmdArgs)
}
