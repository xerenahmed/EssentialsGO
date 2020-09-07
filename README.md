# Essentials GO
Essential commands for [Dragonfly](github.com/df-mc/dragonfly).

## Ready for:
- [x] /gamemode - Changes the player to a specific game mode.
- [x] /teleport - Teleport everywhere.
- [x] /setworldspawn - Show/hide coordinates.
- [x] /defaultgamemode - Sets a worlds's spawn point.
- [x] /xyz - Set the default gamemode.
- [x] /stop - Stop the server from in game.
- [x] /op - Give op permissions to player.
- [x] /deop - Take op permissions from player.

## Usage
### Get the package
`go get -u github.com/eren5960/essentialsgo`
### Import package
```go 
import "github.com/eren5960/essentialsgo"
```
### Load commands
```go
essentialsgo.LoadCommands(server, nil) // the server is *dragonfly.Server{}, nil for load all commands
```
#### Load commands without some commands
```go
essentialsgo.LoadCommands(server, []string{"stop", "defaultgamemode"}) // All will be loaded except "stop" and "defaultgamemode" command
```
#### Finis
If you want to support the project, you can give a star, report errors and bugs as issues.