package commands
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

type Stop struct {}

func (t Stop) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	if !IsOp(p){
		output.Error("You don't have permission for this command.")
		return
	}
	output.Printf("Stopping server.")
	if err := Server.Close(); err != nil {
		output.Printf("error shutting down server: %v", err)
	}
}