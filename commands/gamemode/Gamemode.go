package gamemode
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/commands/op"
	"reflect"
)

type GameMode struct {
	GameMode mode
	// Target   target will be soon
}

func (t GameMode) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	if !op.IsOp(p){
		output.Error("You don't have permission for this command.")
		return
	}
	mode := StringToGameMode(string(t.GameMode))

	p.SetGameMode(mode)
	output.Printf("Set own game mode to %s.", GameModeToName(mode))
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