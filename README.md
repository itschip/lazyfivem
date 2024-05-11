# lazyfivem
A simple CUI for running any FiveM server and executing commands.

## Install

### Windows (scoop)
```
# Add the extras bucket
scoop bucket add itschip https://github.com/itschip/scoop-bucket.git

# Install lazyfivem
scoop install itschip/lazyfivem
```

### Linux
```
curl -Lo lazyfivem.tar.gz https://github.com/itschip/lazyfivem/releases/download/1.0.2/lazyfivem_1.0.2_Linux_32-bit.tar.gz
```
```
sudo tar xf lazyfivem.tar.gz -C /usr/local/bin lazyfivem
```

### Adding servers
Create a `config.yaml` to:
- Linux: ~/.lazyfivem/
- Windows: $HOME:
Create a folder called `.lazyfivem` and a file inside called `config.yaml`.
Example:
```
mkdir $HOME/.lazyfivem
# Then create config.yaml inside
```

You can then add servers to `config.yaml` like this (example)
```yml
Windows Server: "X:/ServerFX/starter.bat"
Linux Server: "cd ~/FXServer/server-data && ~/FXServer/server/run.sh +exec server.cfg"
```

Then, just run `lazyfivem` in any terminal.


## Mappings
* **TAB** - Moves cursor between the sidebar and command exec view
* **ENTER** - Starts the selected when you're in the sidebar view, and executes the command when you're in the command exec view.

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
