package essentialsgo

import (
	"github.com/df-mc/dragonfly/dragonfly"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/xerenahmed/essentialsgo/commands"
	"github.com/xerenahmed/essentialsgo/commands/gamemode"
	"github.com/xerenahmed/essentialsgo/commands/op"
	"github.com/xerenahmed/essentialsgo/commands/time"
	"github.com/xerenahmed/essentialsgo/console"
	"github.com/xerenahmed/essentialsgo/global"
)

func LoadCommands(server *dragonfly.Server, withOut []string) {
	global.Server = server
	op.LoadOps()

	cs := GetCommands()

	for _, c_ := range withOut {
		delete(cs, c_)
	}

	for _, c := range cs {
		cmd.Register(c)
		global.Commands = append(global.Commands, c)
	}
}

func GetCommands() map[string]cmd.Command {
	return map[string]cmd.Command{
		"gamemode":        cmd.New("gamemode", "Changes the player to a specific game mode.", []string{"gm", "gamemode"}, gamemode.GameMode{}),
		"teleport":        cmd.New("teleport", "Teleport everywhere.", []string{"tp", "teleport"}, commands.TeleportXYZ{}, commands.TeleportPlayer{}),
		"xyz":             cmd.New("xyz", "Show/hide coordinates.", []string{"xyz"}, commands.XYZ{}),
		"setworldspawn":   cmd.New("setworldspawn", "Sets a worlds' spawn point. Your coordinates will be used.", []string{"setworldspawn"}, commands.SetWorldSpawnXYZ{}, commands.SetWorldSpawn{}),
		"defaultgamemode": cmd.New("defaultgamemode", "Set the default gamemode.", []string{"defaultgamemode"}, gamemode.DefaultGameMode{}),
		"stop":            cmd.New("stop", "Stop the server.", []string{"stop"}, commands.Stop{}),
		"op":              cmd.New("op", "Give op permissions to a player.", []string{"op"}, op.Op{}),
		"deop":            cmd.New("deop", "Take op permissions from a player.", []string{"deop"}, op.Deop{}),
		"help":            cmd.New("help", "Show server commands and descriptions.", []string{"help"}, commands.Help{}),
		"time":            cmd.New("time", "Changes or queries the worlds game time.", []string{"time"}, time.Set{}, time.SetTimeSpec{}, time.Query{}, time.Add{}),
	}
}

func LoadConsole() {
	console.StartConsole()
}
