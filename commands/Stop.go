package commands
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/commands/op"
	"github.com/eren5960/essentialsgo/global"
)

type Stop struct {}

func (t Stop) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	if !op.IsOp(p){
		output.Error("You don't have permission for this command.")
		return
	}
	output.Printf("Stopping server.")
	if err := global.Server.Close(); err != nil {
		output.Printf("error shutting down server: %v", err)
	}
}