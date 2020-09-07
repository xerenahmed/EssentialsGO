package essentialsgo
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/eren5960/essentialsgo/commands"
	"github.com/eren5960/essentialsgo/commands/gamemode"
	"github.com/eren5960/essentialsgo/commands/op"
	"github.com/eren5960/essentialsgo/global"
)

func LoadCommands(server *dragonfly.Server) {
	global.Server = server
	op.LoadOps()
	cmd.Register(cmd.New("gamemode", "Changes the player to a specific game mode.", []string{"gm", "gamemode"}, gamemode.GameMode{}))
	cmd.Register(cmd.New("teleport", "Teleport everywhere.", []string{"tp", "teleport"}, commands.Teleport{}))
	cmd.Register(cmd.New("xyz", "Show/hide coordinates.", []string{"xyz"}, commands.XYZ{}))
	cmd.Register(cmd.New("setworldspawn", "Sets a worlds's spawn point. Your coordinates will be used.", []string{"setworldspawn"}, commands.SetWorldSpawn{}))
	cmd.Register(cmd.New("defaultgamemode", "Set the default gamemode.", []string{"defaultgamemode"}, gamemode.DefaultGameMode{}))
	cmd.Register(cmd.New("stop", "Stop the server.", []string{"stop"}, commands.Stop{}))
	cmd.Register(cmd.New("op", "Give op permissions to player.", []string{"op"}, op.Op{}))
	cmd.Register(cmd.New("deop", "Take op permissions from player.", []string{"deop"}, op.Deop{}))
}