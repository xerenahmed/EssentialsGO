package gamemode
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/commands/op"
)

type DefaultGameMode struct {
	GameMode mode // @see Gamemode.go for "mode"
	// Target   target will be soon
}

func (t DefaultGameMode) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	if !op.IsOp(p){
		output.Error("You don't have permission for this command.")
		return
	}
	mode := StringToGameMode(string(t.GameMode))

	p.World().SetDefaultGameMode(mode)
	output.Printf("The world's default game mode is now %s.", GameModeToName(mode))
}