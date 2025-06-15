# ktx ☸️

[![GitHub release](https://img.shields.io/github/v/release/itzzjb/kubernetes-context-changer-cli)](https://github.com/itzzjb/kubernetes-context-changer-cli/releases)
[![GitHub issues](https://img.shields.io/github/issues/itzzjb/kubernetes-context-changer-cli)](https://github.com/itzzjb/kubernetes-context-changer-cli/issues)
[![GitHub license](https://img.shields.io/github/license/itzzjb/kubernetes-context-changer-cli)](LICENSE)

A simple, production-ready CLI tool to easily switch Kubernetes contexts.

[![asciicast](https://asciinema.org/a/AXH4Oy2RoQKhzeN3oV9dFrZ4Y.svg)](https://asciinema.org/a/AXH4Oy2RoQKhzeN3oV9dFrZ4Y?t=1)

---

## Table of Contents
- [Features](#features)
- [Quick Start](#quick-start)
- [Supported Platforms](#supported-platforms)
- [Installation](#installation)
- [Usage](#usage)
- [Configuration](#configuration)
- [Troubleshooting](#troubleshooting)
- [FAQ](#faq)
- [Contributing](#contributing)
- [License](#license)

---

## Features
- 🚀 **Interactive context switching**: Quickly list and switch between Kubernetes contexts in your kubeconfig.
- 🛠️ **Cross-platform**: Supports Linux, macOS, and Windows.
- ⚡ **Fast and lightweight**: Built with Go, Cobra, Viper, and Survey for a seamless CLI experience.
- 🔒 **No telemetry**: Your data and kubeconfig never leave your machine.

---

## Quick Start

```sh
# Download and install (Linux/macOS)
wget https://github.com/itzzjb/kubernetes-context-changer-cli/releases/latest/download/ktx-<os>-<arch>
chmod +x ./ktx-<os>-<arch>
sudo mv ./ktx-<os>-<arch> /usr/local/bin/ktx
ktx
```

---

## Supported Platforms
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

---

## Installation

### Mac / Linux
1. Download the binary for your OS from the [Releases](https://github.com/itzzjb/kubernetes-context-changer-cli/releases) page.
2. Make it executable and move to your PATH:
   ```sh
   chmod +x ./ktx-<os>-<arch>
   mv ./ktx-<os>-<arch> ktx
   sudo mv ./ktx /usr/local/bin
   ```
3. Or build from source (requires Go 1.18+):
   ```sh
   git clone https://github.com/itzzjb/kubernetes-context-changer-cli.git
   cd kubernetes-context-changer-cli
   go build -o ktx
   sudo mv ./ktx /usr/local/bin
   ```
4. Run `ktx` to verify installation:
   ```sh
   ktx
   ```

### Windows
1. Download `ktx-windows-amd64.exe` from [Releases](https://github.com/itzzjb/kubernetes-context-changer-cli/releases).
2. Rename to `ktx.exe`.
3. Add the folder containing `ktx.exe` to your system `PATH`.
4. Run from Command Prompt or PowerShell:
   ```sh
   ktx
   ```
> - You can also run the tool by double-clicking `ktx.exe`, but it is designed for interactive use in a terminal.
> - Make sure your `KUBECONFIG` environment variable is set if your kubeconfig is not in the default location (`%USERPROFILE%\.kube\config`).

---

## Usage

Switch context interactively:
```sh
ktx
```

List all available contexts:
```sh
ktx list
```

Switch to a specific context:
```sh
ktx <context-name>
```

---

## Configuration

- By default, `ktx` uses the kubeconfig at `$KUBECONFIG` or `~/.kube/config`.
- To use a different kubeconfig:
  ```sh
  ktx --kubeconfig <path-to-kubeconfig>
  ```

---

## Troubleshooting

- **macOS Security Warning:**
  When running `ktx` for the first time on macOS, you may see a security warning.
  > [!CAUTION]
  > **Apple cannot verify this app for malicious software**
  >
  > This is standard for open source CLI tools not distributed via the Mac App Store.
  >
  > **Solution:**
  > 1. Right-click the `ktx` binary and choose **Open**. Click **OK** in the dialog. You only need to do this once per version.
  > 2. Or, if you see a quarantine error, run:
  >    ```sh
  >    xattr -d com.apple.quarantine ./ktx
  >    ```

- **Command not found:**
  Ensure the binary is in your `PATH` and is executable.

- **Permission denied:**
  Use `chmod +x ./ktx` to make it executable.

---

## FAQ

**Q: Does `ktx` modify my kubeconfig?**
A: `ktx` only switches the current context; it does not modify or delete clusters, users, or contexts.

**Q: Can I use `ktx` with multiple kubeconfig files?**
A: Yes, set the `KUBECONFIG` environment variable before running `ktx`.

**Q: Is telemetry or analytics collected?**
A: No, `ktx` does not collect or send any telemetry data.

---

## Contributing

Pull requests and issues are welcome! Please open an issue to discuss your ideas or report bugs. For major changes, please open an issue first to discuss what you would like to change.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/my-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/my-feature`)
5. Open a pull request

---

## License

This project is licensed under the MIT License. See [LICENSE](LICENSE) for details.
