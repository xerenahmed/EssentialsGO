package essentialsgo

import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/eren5960/essentialsgo/commands"
)

func LoadCommands(server *dragonfly.Server) {
	commands.Server = server
	cmd.Register(cmd.New("gamemode", "Change gamemode of player.", []string{"gm", "gamemode"}, commands.GameMode{}))
}
