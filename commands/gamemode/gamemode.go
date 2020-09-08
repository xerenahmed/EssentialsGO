package gamemode
// Eren5960 <ahmederen123@gmail.com>
import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/commands/op"
	"github.com/eren5960/essentialsgo/commands/utils"
	"reflect"
)

type GameMode struct {
	GameMode mode
	Target   string `optional:""`
}

func (t GameMode) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source){
		output.Error("You don't have permission for this command.")
		return
	}

	mode := StringToGameMode(string(t.GameMode))
	modeString :=  GameModeToName(mode)

	if t.Target != "" {
		if pt, _ := utils.PlayerByName(t.Target); pt == nil{
			output.Error(t.Target + " not found.")
		} else {
			pt.SetGameMode(mode)
			pt.Message(fmt.Printf("Your game mode has been changed to %s.", modeString))
			output.Printf("Set %s game mode to %s.", t.Target, modeString)
		}

		return
	}

	if p, ok := source.(*player.Player); ok {
		p.SetGameMode(mode)
		output.Printf("Set own game mode to %s.", modeString)
	} else {
		output.Error("Usage: /gamemode <GameMode: mode> <Target: string>")
	}
}

type mode string

func (mode) Type() string {
	return "mode"
}

func (mode) Options() []string {
	return []string{"0", "1", "2", "s", "c", "a", "survival", "creative", "adventure"}
}

func (mode) SetOption(option string, r reflect.Value) {
	r.SetString(option)
}