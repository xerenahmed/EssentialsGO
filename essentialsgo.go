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

func LoadCommands(server *dragonfly.Server, withOut []string) {
	global.Server = server
	op.LoadOps()

	for name, c := range GetCommands() {
		if withOut != nil && len(withOut) > 0 {
			skip := false
			for _, c_ := range withOut{
				if name == c_ {
					skip = true
					break
				}
			}
			if skip {
				continue
			}
		}
		cmd.Register(c)
	}
}

func GetCommands() map[string]cmd.Command{
	return map[string]cmd.Command{
		"gamemode": 	   cmd.New("gamemode", "Changes the player to a specific game mode.", []string{"gm", "gamemode"}, gamemode.GameMode{}),
		"teleport":        cmd.New("teleport", "Teleport everywhere.", []string{"tp", "teleport"}, commands.Teleport{}),
		"xyz": 		       cmd.New("xyz", "Show/hide coordinates.", []string{"xyz"}, commands.XYZ{}),
		"setworldspawn":   cmd.New("setworldspawn", "Sets a worlds's spawn point. Your coordinates will be used.", []string{"setworldspawn"}, commands.SetWorldSpawn{}),
		"defaultgamemode": cmd.New("defaultgamemode", "Set the default gamemode.", []string{"defaultgamemode"}, gamemode.DefaultGameMode{}),
		"stop": 		   cmd.New("stop", "Stop the server.", []string{"stop"}, commands.Stop{}),
		"op":		       cmd.New("op", "Give op permissions to player.", []string{"op"}, op.Op{}),
		"deop":			   cmd.New("deop", "Take op permissions from player.", []string{"deop"}, op.Deop{}),
	}
}