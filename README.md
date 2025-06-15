# ktx ☸️

A simple, production-ready CLI tool to easily switch Kubernetes contexts.

## Features 
- List and switch between Kubernetes contexts interactively 
- Built with Cobra, Viper, and Survey for a great CLI experience


## Supported Platforms 
- Linux (amd64, arm64)
- macOS (amd64, arm64)
- Windows (amd64)

## Installation 

### Mac / Linux 
1. Download the binary for your OS from the [Releases](https://github.com/itzzjb/kubernetes-context-changer-cli/releases) page.
2. Make it executable and move to your PATH:
   ```sh
   chmod +x ./ktx-<os>-<arch>
   ```
   ```sh
   mv ./ktx-<os>-<arch> ktx
   ```
   ```sh
   sudo mv ./ktx /usr/local/bin
   ```
3. Or build from source (requires Go 1.18+):
   ```sh
   git clone https://github.com/itzzjb/kubernetes-context-changer-cli.git
   ```
   ```sh
   cd kubernetes-context-changer-cli
   ```
   ```sh
   go build -o ktx
   ```
   ```sh
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

#### Notes 
- As long as `ktx.exe` is in your `PATH`, you do not need to type the `.exe` extension—just use `ktx`.
- If you are running from the current directory and it's not in your `PATH`, use `./ktx` or `./ktx.exe`.
- You can also run the tool by double-clicking `ktx.exe`, but it is designed for interactive use in a terminal.
- Make sure your `KUBECONFIG` environment variable is set if your kubeconfig is not in the default location (`%USERPROFILE%\.kube\config`).

---

## macOS: Security Notice

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

---

## Contributing 

Pull requests and issues are welcome! Please open an issue to discuss your ideas or report bugs.

---

## License 

License. See [LICENSE](LICENSE) for details.
