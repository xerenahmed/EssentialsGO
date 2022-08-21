package gamemode

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xerenahmed/essentialsgo/commands/op"
)

type GameMode struct {
	GameMode mode
	Target   []cmd.Target `optional:""`
}

func (t GameMode) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}

	mode := StringToGameMode(string(t.GameMode))
	modeString := GameModeToName(mode)

	if len(t.Target) > 0 {
		if pt, ok := t.Target[0].(*player.Player); ok {
			pt.SetGameMode(mode)
			output.Printf("Set %s game mode to %s.", pt.Name(), modeString)
			pt.Messagef("Your game mode has been changed to %s.", modeString)
		} else {
			output.Errorf("Target is invalid!")
		}

		return
	}

	if p, ok := source.(*player.Player); ok {
		p.SetGameMode(mode)
		output.Printf("Set own game mode to %s.", modeString)
	} else {
		output.Error("Usage: /gamemode <GameMode: mode> <Target: target>")
	}
}
