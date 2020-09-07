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
	p := source.(*player.Player)
	if !op.IsOp(p){
		output.Error("You don't have permission for this command.")
		return
	}

	mode := StringToGameMode(string(t.GameMode))
	modeString :=  GameModeToName(mode)

	if !utils.SubEmpty(t.Target) {
		if pt, _ := utils.PlayerByName(t.Target); pt == nil{
			output.Error(t.Target + " can't found.")
		} else {
			pt.SetGameMode(mode)
			pt.Message(fmt.Printf("Your game mode has been changed to %s.", modeString))
		}

		return
	}

	p.SetGameMode(mode)
	output.Printf("Set own game mode to %s.", modeString)
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

func (GameMode) Cmd() string{
	return "/gamemode"
}