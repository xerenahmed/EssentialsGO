package op
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

type Deop struct {
	Target string `optional:""`
}

func (t Deop) Run(source cmd.Source, output *cmd.Output) {
	if !IsOp(source){
		output.Error("You don't have permission for this command.")
		return
	}
	var pt string

	if t.Target == "" {
		if p, ok := source.(*player.Player); ok {
			pt = p.Name()
		}
	}

	if pt != "" {
		DelOp(pt)
		output.Printf("Has been taken op permissions from %s.", pt)
	} else {
		output.Error("Usage: /deop <Player: string>")
	}
}