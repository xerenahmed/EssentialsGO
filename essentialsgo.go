package essentialsgo
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/eren5960/essentialsgo/commands"
)

func LoadCommands(server *dragonfly.Server) {
	commands.Server = server
	cmd.Register(cmd.New("gamemode", "Changes the player to a specific game mode.", []string{"gm", "gamemode"}, commands.GameMode{}))
	cmd.Register(cmd.New("teleport", "Teleport everywhere.", []string{"tp", "teleport"}, commands.Teleport{}))
	cmd.Register(cmd.New("xyz", "Show coordinates.", []string{"xyz"}, commands.XYZ{}))
	cmd.Register(cmd.New("setworldspawn", "Sets a worlds's spawn point. Your coordinates will be used.", []string{"setworldspawn"}, commands.SetWorldSpawn{}))
}
