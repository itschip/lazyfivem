# lazyfivem
A simple TUI for running any FiveM server and executing commands.

## Installation

### Windows
- via Scoop (Recommended)
```
# Add the extras bucket
scoop bucket add itschip https://github.com/itschip/scoop-bucket.git

# Install lazyfivem
scoop install itschip/lazyfivem
```

### Linux
```bash
curl -Lo lazyfivem.tar.gz https://github.com/itschip/lazyfivem/releases/download/1.0.2/lazyfivem_1.0.2_Linux_32-bit.tar.gz

sudo tar xf lazyfivem.tar.gz -C /usr/local/bin lazyfivem
```

### Via GitHub Releases
Go to [the latest release page](https://github.com/itschip/lazyfivem/releases/latest) and download the corresponding version for your OS.

## Configuration
Create a `config.yaml` to:

- Unix-based OS:
    - if `$XDG_CONFIG_HOME` is specified, then: `$XDG_CONFIG_HOME/lazyfivem/`
    - otherwise `$HOME/.config/lazyfivem/`

- Darwin: `$HOME/Library/Application Support/lazyfivem/`

- Plan 9: `$home/lib`

- Windows: `C:\Users\USERNAME\AppData\Roaming\lazyfivem\` aka. `%APPDATA%\lazyfivem\`

- [Read more](https://pkg.go.dev/os#UserConfigDir)

Example:
- Windows:
  ```pwsh
  mkdir $APPDATA\lazyfivem
  
  # Then create config.yaml inside
  ```
- Linux:
  ```bash
  mkdir ~/.config/lazyfivem
  
  cd ~/.config/lazyfivem
  
  touch config.yaml
  ```

### Custom Config Directory

You can change the config directory by exporting the `LAZYFIVEM_CONFIG_HOME` environment variable. You can then place the `config.yaml` file there.

### Adding Servers

Create a new entry in the `config.yaml` file as follows:

`Server Name: '<command to start the server>'`

#### Windows:

```yaml
Server Name: 'PATH\TO\SERVER\start.bat'
```

Example:
```yaml
Yet Another RP Server: 'D:\FxServer\start.bat'
```

#### Linux:

```yaml
Server Name: 'PATH/TO/SERVER/start.sh'
```

Example:
```yaml
Yet Another RP Server: 'PATH/TO/SERVER/start.sh'
```

## Usage
Run `lazyfivem` in any terminal.

### Navigation
* **TAB** - Switches focus between the sidebar and command line
* **ENTER** - Starts the selected server in the sidebar, and executes a command in the command line.
* **UP/DOWN ARROW** - Scrolls up/down in the sidebar
* **LEFT/RIGHT ARROW** - Moves cursor left/right in the command exec lin
* **Ctrl-C** - Quit program

### Example of Prisma in FiveM failing.
![vmplayer_lqEQwqQhEj](https://user-images.githubusercontent.com/59088889/190914038-df755ce0-7485-4c7b-95fb-7af33a2cb686.png)


## Toggleterm
```lua
local lazyfivem = Terminal:new({
  cmd = "lazyfivem",
  direction = "float",
  count = 15,
  float = {
    border = "double",
  }
})

function _lazyfivem_toggle()
  lazyfivem:toggle()
end
```
