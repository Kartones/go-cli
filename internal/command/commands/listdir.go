package commands

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type ListDirCommand struct {
	writer io.Writer
}

func NewListDirCommand() *ListDirCommand {
	return &ListDirCommand{
		writer: os.Stdout,
	}
}

func (l *ListDirCommand) Name() string {
	return "list-dir"
}

func (l *ListDirCommand) Description() string {
	return "List directory contents"
}

func (l *ListDirCommand) Usage() string {
	return "kc list-dir <directory>"
}

func (l *ListDirCommand) Execute(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no directory specified")
	}

	dirPath := args[0]
	dirPath = filepath.Clean(dirPath)

	info, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return fmt.Errorf("directory not found: %s", dirPath)
	}
	if err != nil {
		return fmt.Errorf("error accessing directory: %v", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("path is not a directory: %s", dirPath)
	}

	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return fmt.Errorf("error reading directory: %v", err)
	}

	return l.displayEntries(dirPath, entries)
}

func (l *ListDirCommand) displayEntries(dirPath string, entries []os.DirEntry) error {
	fmt.Fprintf(l.writer, "Directory: %s\n", dirPath)

	if len(entries) == 0 {
		fmt.Fprintf(l.writer, "[EMPTY]\n")
		return nil
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}

		entryType := "FIL"
		if entry.IsDir() {
			entryType = "DIR"
		}

		size := fmt.Sprintf("%d", info.Size())

		fmt.Fprintf(l.writer, "%s %-25s %-15s\n", entryType, entry.Name(), size)
	}

	return nil
}
