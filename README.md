# Essentials GO
Essential commands for [Dragonfly](https://github.com/df-mc/dragonfly).

## Ready for:
- [x] /help - Show server commands and descriptions.
- [x] /gamemode - Changes the player to a specific game mode.
- [x] /teleport - Teleport everywhere.
- [x] /defaultgamemode - Set the default gamemode.
- [x] /setworldspawn - Sets a worlds' spawn point.
- [x] /xyz - Show/hide coordinates.
- [x] /op - Give op permissions to player.
- [x] /deop - Take op permissions from a player.
- [x] /stop - Stop the server from in-game.
- [x] /time - Changes or queries the worlds game time.

## Usage
### Get the package
`go get -u github.com/eren5960/essentialsgo`
### Import package
```go 
import "github.com/eren5960/essentialsgo"
```
### Load commands
```go
essentialsgo.LoadCommands(server, nil) // the server is *dragonfly.Server{}, nil to load all commands
```
### Exclude commands from loading
```go
essentialsgo.LoadCommands(server, []string{"stop", "defaultgamemode"}) // All commands will be loaded, except "stop" and "defaultgamemode"
```

### Lastly
If you want to support the project, you can give this project a star.<br>
Please report issues or errors in the Issues tab.

### Simple Console Command Sender
```go
essentialsgo.LoadConsole()
```
![](https://user-images.githubusercontent.com/35738714/92475213-b255e580-f1e5-11ea-9e15-c5cbfc71e98e.gif)
