# Techtonic Core

Techtonic Core is a fundamental component of TDK (Techtonic Development Kit), designed to provide essential utilities and libraries for building Linux applications in Go. It serves as the foundation for the Techtonic ecosystem, offering various modules for common development needs.

## Components

### Currently Implemented

#### 🌋 Seismic (Logging)
A flexible logging system built on top of `uber-zap`, providing structured logging capabilities with support for different output formats.

```go
logger := seismic.NewLogger("console")
logger.Info("Hello, World!", "key", "value")
```

### Planned Components

#### 🌊 Subduction (Error Handling)
A comprehensive error handling system (Coming soon)

#### 🗺️ Strata (Configuration)
YAML configuration management system (Coming soon)

#### ⚡ Tecto (Configuration)
Custom configuration format handler (Coming soon)

#### 🛠️ Platform Utilities
- Filesystem operations
- Network utilities
- And more...

## Installation

```bash
go get gitlab.com/techtonic-team/tdk/techtonic-core
```

## Quick Start

### Logging with Seismic

```go
package main

import "gitlab.com/techtonic-team/tdk/techtonic-core/pkg/logging/seismic"

func main() {
    logger := seismic.NewLogger("console")
    defer logger.Sync()

    logger.Info("Hello, World!", "greeting", "Hello", "target", "World")
}
```

## Project Structure

```
techtonic-core/
├── pkg/
│   ├── logging/
│   │   └── seismic/      # Logging system
│   ├── errors/           # Future error handling
│   └── ...
└── examples/
    └── hello/           # Example applications
```

## Requirements

- Go 1.22 or higher
- Linux operating system

## License

TDK is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.