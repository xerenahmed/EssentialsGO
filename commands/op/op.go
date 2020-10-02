package op
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
)

type Op struct {
	Player string
}

func (t Op) Run(source cmd.Source, output *cmd.Output) {
	if !IsOp(source){
		output.Error("You don't have permission for this command.")
		return
	}

	if t.Player != "" {
		AddOp(t.Player)
		output.Printf("Has been granted op permissions to %s.", t.Player)
	} else {
		output.Error("Usage: /op <Player: string>")
	}
}

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