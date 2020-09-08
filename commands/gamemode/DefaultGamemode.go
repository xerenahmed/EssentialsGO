package gamemode
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/eren5960/essentialsgo/commands/op"
	"github.com/eren5960/essentialsgo/global"
)

type DefaultGameMode struct {
	GameMode mode // @see Gamemode.go for "mode"
}

func (t DefaultGameMode) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source){
		output.Error("You don't have permission for this command.")
		return
	}
	mode := StringToGameMode(string(t.GameMode))
	global.Server.World().SetDefaultGameMode(mode)
	output.Printf("The world's default game mode is now %s.", GameModeToName(mode))
}