# Installation

## Prerequisites
To build the Base Infrastructure platform from source, you must have the following installed:
- [Go](https://golang.org/dl/) 1.21 or higher
- Git
- Make (optional, but recommended)

## Building from Source

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Hardik-Sankhla/base-infrastructure.git
   cd base-infrastructure
   ```

2. **Build the binary:**
   Using Make:
   ```bash
   make
   ```
   Or using Go directly:
   ```bash
   go build -o platform ./cmd/platform
   ```

3. **Verify the installation:**
   ```bash
   ./platform --help
   ```

## Termux Setup (Android)
If you are running the platform natively on Android via Termux, prepare your Go environment:

```bash
pkg update && pkg upgrade -y
pkg install golang git make clang -y
```

Then follow the standard clone and build steps above.

> [!NOTE]
> Pre-compiled binaries for major operating systems (Linux, Windows, macOS) will be provided via GitHub Releases in future updates.
