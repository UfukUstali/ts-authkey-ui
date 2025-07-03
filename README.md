# Tailscale Auth Key UI

A simple web-based GUI for authenticating Tailscale using an auth key. This application provides a user-friendly interface to input your Tailscale auth key and automatically start the Tailscale service.

## Features

- Clean, responsive web interface
- Cross-platform support (Windows, macOS, Linux)
- Automatic browser opening
- Input validation and error handling
- Automatic application shutdown after successful authentication
- German localization

## Prerequisites

- [Tailscale](https://tailscale.com/) must be installed and available in your system PATH
- Go 1.24.3 or later (for building from source)

## Usage

1. Clone the repo and build or run using `go` or run the `./build.sh` and start the correct bin from `./dist`
2. Your default browser will open automatically to `http://localhost:8924`
3. Enter your Tailscale auth key in the form
4. Click "Tailscale starten" to authenticate and start Tailscale
5. The application will automatically close after successful authentication

## Building from Source

### Quick Build
```bash
go build -o ts-authkey-ui
```

### Cross-Platform Build
Use the included build script to create binaries for all supported platforms:

```bash
chmod +x build.sh
./build.sh
```

This will create binaries in the `dist/` directory for:
- Windows (amd64)
- macOS (amd64 and arm64)
- Linux (amd64 and arm64)

## Configuration

The application runs on port `8924` by default. This is currently hardcoded but can be modified in the source code if needed.

## Security Notes

- The application only runs locally and does not expose any external network interfaces
- Auth keys are not stored or logged
- The application terminates immediately after successful authentication

---

🤖 Created with [opencode](https://opencode.ai)

### The Prompt:

this project should be compiled to a single binary that should be able to run on any desktop os.
it should serve a html file over http over an uncommon port, open that port with localhost in the default browser.
the html page should show a single input box and a submit button.
when clicked it will make a post request to the running http server and the http server will call `tailscale up --auth-key <key>` and exit if succesfull or return and display an error.
- check if tailscale exists if not show an error page
- the ui should be in german
