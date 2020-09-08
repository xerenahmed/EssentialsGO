package op
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
)

type Deop struct {
	Player string
}

func (t Deop) Run(source cmd.Source, output *cmd.Output) {
	if !IsOp(source){
		output.Error("You don't have permission for this command.")
		return
	}

	if t.Player != "" {
		DelOp(t.Player)
		output.Printf("Has been taken op permissions from %s.", t.Player)
	} else {
		output.Error("Usage: /deop <Player: string>")
	}
}