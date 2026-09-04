# Password Generator

A fast and customizable command-line tool written in Go to generate random passwords concurrently.

## Features

- **Customizable Length**: Generate passwords of any desired length (default: 256 characters).
- **Character Control**: Toggle numbers and special characters on or off.
- **Concurrent Generation**: Utilizes Go goroutines for fast generation.
- **Clean CLI**: Built with Cobra for an intuitive command-line interface.

## Prerequisites

- [Go](https://go.dev/dl/) 1.19 or higher

## Installation & Build

Clone the repository and build the binary:

```bash
git clone https://github.com/juancarbajal/password-generator.git
cd password-generator
go build -o password-generator
```

Alternatively, you can run it directly without compiling a binary:

```bash
go run main.go
```

## Usage

```bash
./password-generator [flags]
```

### Flags

| Flag | Shorthand | Description | Default |
|------|-----------|-------------|---------|
| `--size` | `-s` | Length of the generated password | `256` |
| `--no-numbers` | `-n` | Do not include numbers (`0-9`) | `false` |
| `--no-symbols` | `-x` | Do not include special characters | `false` |
| `--help` | `-h` | Display help information | |

### Examples

**Generate a default 256-character password:**
```bash
./password-generator
```

**Generate a 16-character password:**
```bash
./password-generator -s 16
```

**Generate a 32-character password without special characters:**
```bash
./password-generator -s 32 -x
```

**Generate a password without numbers:**
```bash
./password-generator -s 20 -n
```

**Generate an alphabet-only password (no numbers or symbols):**
```bash
./password-generator -s 16 -n -x
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
