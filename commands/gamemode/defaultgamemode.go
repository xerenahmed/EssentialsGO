package gamemode

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/xerenahmed/essentialsgo/commands/op"
	"github.com/xerenahmed/essentialsgo/global"
)

type DefaultGameMode struct {
	GameMode mode // look at gamemode.go for "mode"
}

func (t DefaultGameMode) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}
	mode := StringToGameMode(string(t.GameMode))
	global.Server.World().SetDefaultGameMode(mode)
	output.Printf("Default world game mode is %s now.", GameModeToName(mode))
}
