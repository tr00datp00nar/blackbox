# Blackbox

🌳Stateful Command Tree Monolith with Bonzai for the Silver Spring Black Box Theater.

## Prerequisites

- [Google Cloud](cloud.google.com) Account
- Google cloud OAuth credentials
- A `.env` file or sourced environment variables for the following:
  (See the `.env.example` file)
  - CALENDAR_ID
  - RESULTS_LOC

### Setting up the configuration location

The blackbox tool looks for configuration in the following directories:

- Linux: `$XDG_CONFIG_DIR/blackbox`
- Windows: `%APPDATA%\blackbox`
- macOS: `$HOME/Library/blackbox`

Create the directory if it doesn't exist, using your preferred method.

### Creating Google Cloud Oauth Credentials

1. Log in to the Google Cloud Console
2. Create a new project and navigate to it
3. Go to "APIs & Services"
4. Click "+ Enable APIs & Services"
5. Search for `Google Calendar API` and select and enable it
6. Navigate to "Credentials"
7. Click "+ CREATE CREDENTIALS" > "+ OAuth client ID"
8. Configure the OAuth Consent Screen

   - External is the default if you aren't using a Google Workspaces account.

9. Navigate back to "Credentials"
10. Click "+ CREATE CREDENTIALS" > "+ OAuth client ID"
11. Select Application type: Desktop app
12. Name your application, e.g., "blackbox"
13. Download your credentials JSON as credentials.json

### Creating your .env file

```bash
cp  .env.example /path/to/your/config/directory/.env
```

Edit the `.env` as needed

## Installation

### Local Clone (RECOMMENDED)

1. Clone this repository using `git clone https://github.com/tr00datp00nar/blackbox.git` and from within that directory, run:

```bash
go install .
```

2. Create the necessary configuration directory from above.
3. Make sure that the configuration directory is populated with your `.env` and `credentials.json` files.

### Download one of the [release binaries](https://github.com/tr00datp00nar/blackbox/releases) (NOT RECOMMENDED):

```bash
curl -L https://github.com/tr00datp00nar/blackbox/releases/latest/download/blackbox-linux-amd64 -o ~/.local/bin/tr00datp00nar
curl -L https://github.com/tr00datp00nar/blackbox/releases/latest/download/blackbox-darwin-amd64 -o ~/.local/bin/tr00datp00nar
curl -L https://github.com/tr00datp00nar/blackbox/releases/latest/download/blackbox-darwin-arm64 -o ~/.local/bin/tr00datp00nar
curl -L https://github.com/tr00datp00nar/blackbox/releases/latest/download/blackbox-windows-amd64 -o ~/.local/bin/tr00datp00nar
```

### Install directly with `go` (NOT RECOMMENEDED):

```bash
go install github.com/tr00datp00nar/blackbox@latest
```

## Tab Completion in Bash

To activate bash completion just use the `complete -C` option from your `.bashrc` or command line. There is no messy sourcing required. All the completion is done by the program itself.

```bash
complete -C blackbox blackbox
```

If you don't have bash or tab completion check out the shortcut commands instead.

## Tab Completion in Zsh

Zsh does a good job of learning your commands over time all by itself, but some of the custom completions may not work as well. Personally, I use the Oh-My-Zsh option below, but the creator of Bonzai and the original Z command tree (rwxrob) prefers the default Linux shell (Bash) over the default Mac shell (Zsh). (PRs to rwxrob's repository are welcome to integrate completion into Zsh without dumping a ton of shell code that has to be sourced.)

### Oh-My-Zsh

Oh-My-Zsh has an available plugin called [zsh-bash-completions-fallback](https://github.com/3v1n0/zsh-bash-completions-fallback). This plugin allows zsh to fallback to bash completions when it can't find the appropriate completions itself.

Once installed, you can use the same `complete -C blackbox blackbox` as you normally would in bash.

## Embedded Documentation

All documentation (like manual pages) has been embedded into the source code of the application. See the source or run the program with help to access it.

## Building

Releases are built using the following commands:

```bash
blackbox go build
gh release create
gh release upload TAG build/*
```
