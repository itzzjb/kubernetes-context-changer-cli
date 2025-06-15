# ktx

A simple, production-ready CLI tool to easily switch Kubernetes contexts.

- Fast, interactive context selection
- Works with any kubeconfig
- User-friendly terminal prompts
- Supports Linux, macOS, and Windows

---

## Features

- List and switch between Kubernetes contexts interactively
- Built with Cobra, Viper, and Survey for a great CLI experience
- Colorful terminal output
- Cross-platform support (amd64/arm64 for Linux/macOS, amd64 for Windows)

---

## Installation

### Download a Release

Go to the [Releases](https://github.com/itzzjb/kubernetes-context-changer-cli/releases) page and download the appropriate binary for your OS and architecture:

- `ktx-linux-amd64`, `ktx-linux-arm64`
- `ktx-darwin-amd64`, `ktx-darwin-arm64` (macOS Intel/Apple Silicon)
- `ktx-windows-amd64.exe`

Make it executable (Linux/macOS):
```sh
chmod +x ./ktx-<os>-<arch>
```
Move it to a directory in your `$PATH`, e.g.:
```sh
sudo mv ./ktx-<os>-<arch> /usr/local/bin/ktx
```

### Build from Source

Requires Go 1.18+:
```sh
git clone https://github.com/itzzjb/kubernetes-context-changer-cli.git
cd kubernetes-context-changer-cli
go build -o ktx
```

---

## Usage

```sh
ktx [context]
```

- Run `ktx` to interactively select a context.
- Or specify a context name directly: `ktx my-context`

---

## Supported Platforms
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

---

## macOS: Security Notice

When running `ktx` for the first time on macOS, you may see:
> "Apple cannot verify this app for malicious software"

**Solution:**
1. Right-click the `ktx` binary and choose **Open**. Click **OK** in the dialog. You only need to do this once per version.
2. Or, if you see a quarantine error, run:
   ```sh
   xattr -d com.apple.quarantine ./ktx
   ```

This is standard for open source CLI tools not distributed via the Mac App Store.

---

## Contributing

Pull requests and issues are welcome! Please open an issue to discuss your ideas or report bugs.

---

## License

MIT License. See [LICENSE](LICENSE) for details.
