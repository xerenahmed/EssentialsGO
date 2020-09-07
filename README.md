# Essentials GO
Essential commands for Dragonfly.

## Ready for:
- [x] /gamemode - Changes the player to a specific game mode.
- [x] /teleport - Teleport everywhere.
- [x] /setworldspawn - Show/hide coordinates.
- [x] /defaultgamemode - Sets a worlds's spawn point.
- [x] /xyz - Set the default gamemode.
- [x] /stop - Stop the server from in game.

## Usage
### Get the package
`go get -u github.com/eren5960/essentialsgo`
### Import package
```go 
import "github.com/eren5960/essentialsgo"
```
### Load command
```go
essentialsgo.LoadCommands(server) // the server is *dragonfly.Server{}
```
#### Finis
If you want to support the project, you can give a star, report errors and bugs as issues.