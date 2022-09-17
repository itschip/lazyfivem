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
curl -Lo lazyfivem.tar.gz https://github.com/itschip/lazyfivem/releases/download/1.0.1/lazyfivem_1.0.1_Linux_32-bit.tar.gz
```
```
sudo tar xf lazyfivem.tar.gz -C /usr/local/bin lazyfivem
```

### Adding servers
Create a folder called `.lazyfivem` and a file inside called `config.yaml` at the very root of your computer.
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

![image](https://user-images.githubusercontent.com/59088889/189553625-afa0926d-a16d-4be3-8023-fb20f4e8f95c.png)


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
