package commands
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	"reflect"
)

type GameMode struct {
	GameMode mode
	// Target   target will be soon
}

func (t GameMode) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	mode := map[mode]gamemode.GameMode{
		"0": gamemode.Survival{}, "s": gamemode.Survival{},
		"1": gamemode.Creative{}, "c": gamemode.Creative{},
		"2": gamemode.Adventure{}, "a": gamemode.Adventure{},
	}[t.GameMode]

	p.SetGameMode(mode)
	output.Printf("Set own game mode to " + map[gamemode.GameMode]string{
		gamemode.Survival{}: "survival", gamemode.Creative{}: "creative", gamemode.Adventure{}: "adventure",
	}[mode] + ".")
}

type mode string

func (mode) Type() string {
	return "mode"
}

func (mode) Options() []string {
	return []string{"0", "1", "2", "s", "c", "a"}
}

func (mode) SetOption(option string, r reflect.Value) {
	r.SetString(option)
}