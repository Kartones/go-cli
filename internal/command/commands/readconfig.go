package commands

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

type ReadConfigCommand struct {
	writer io.Writer
}

func NewReadConfigCommand() *ReadConfigCommand {
	return &ReadConfigCommand{
		writer: os.Stdout,
	}
}

func (r *ReadConfigCommand) Name() string {
	return "read-config"
}

func (r *ReadConfigCommand) Description() string {
	return "Read JSON or YAML files"
}

func (r *ReadConfigCommand) Usage() string {
	return "kc read-config <file>"
}

func (r *ReadConfigCommand) Execute(ctx context.Context, args []string) error {
	if len(args) == 0 {
		return fmt.Errorf("no file specified\nUsage: %s", r.Usage())
	}

	filename := args[0]
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return fmt.Errorf("file not found: %s", filename)
	}

	content, err := os.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading file: %v", err)
	}

	ext := strings.ToLower(filepath.Ext(filename))

	var data map[string]interface{}

	switch ext {
	case ".json":
		err = json.Unmarshal(content, &data)
		if err != nil {
			return fmt.Errorf("error parsing JSON: %v", err)
		}
	case ".yaml", ".yml":
		err = yaml.Unmarshal(content, &data)
		if err != nil {
			return fmt.Errorf("error parsing YAML: %v", err)
		}
	default:
		return fmt.Errorf("only JSON and YAML files are supported")
	}

	return r.displayConfig(filename, data)
}

func (r *ReadConfigCommand) displayConfig(filename string, data map[string]interface{}) error {
	fmt.Fprintf(r.writer, "Configuration file: %s\n", filename)
	fmt.Fprintf(r.writer, "First-level keys:\n")

	if len(data) == 0 {
		fmt.Fprintf(r.writer, "  (no keys found)\n")
		return nil
	}

	keys := make([]string, 0, len(data))
	for key := range data {
		keys = append(keys, key)
	}
	// sort keys for consistent output (specially yaml files)
	sort.Strings(keys)

	for _, key := range keys {
		value := data[key]
		valueType, err := getValueType(value)
		if err != nil {
			return fmt.Errorf("error getting type for key %s: %v", key, err)
		}
		fmt.Fprintf(r.writer, "  %-20s %s\n", key, valueType)
	}

	return nil
}

func getValueType(value interface{}) (string, error) {
	switch v := value.(type) {
	case nil:
		return "null", nil
	case bool:
		return fmt.Sprintf("bool: %t", v), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("number: %v", v), nil
	case float32, float64:
		return fmt.Sprintf("float: %v", v), nil
	case string:
		return fmt.Sprintf("string: %s", v), nil
	case []interface{}:
		return fmt.Sprintf("array with %d items", len(v)), nil
	case map[string]interface{}:
		return fmt.Sprintf("object with %d keys", len(v)), nil
	default:
		return "", fmt.Errorf("unknown type: %T", v)
	}
}
