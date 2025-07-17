# kc - Kartones CLI

A simple CLI tool to showcase how easy it is to build it with Go supporting multiple operating systems.

## Commands

- `help` - Show help information (default command)
- `read-config <file>` - Example of parsing JSON and YAML files
- `list-dir <directory>` - Example of using the `os` package (listing files and directories)

## Building

### Prerequisites

- [Bazel](https://bazel.build/) (tested with v8.3) and [Bazelisk](https://github.com/bazelbuild/bazelisk)

### Build for the current platform

```bash
bazelisk build //:kc
```

### Cross-compilation

```bash
# macOS
bazelisk build //:kc_darwin_amd64
bazelisk build //:kc_darwin_arm64

# Linux
bazelisk build //:kc_linux_amd64
bazelisk build //:kc_linux_arm64

# Windows
bazelisk build //:kc_windows_amd64
bazelisk build //:kc_windows_arm64
```

Binaries will be available in `bazel-bin/`.

## Usage Examples

```bash
# needed the first time
bazelisk build //:kc

./bazel-bin/kc_/kc help

./bazel-bin/kc_/kc read-config sample-config.json

./bazel-bin/kc_/kc read-config sample-config.yaml

./bazel-bin/kc_/kc list-dir .
``` 