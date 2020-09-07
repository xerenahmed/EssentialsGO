package op
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/global"
)

type Deop struct {
	Target string `optional:""`
}

func (t Deop) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	if !IsOp(p){
		output.Error("You don't have permission for this command.")
		return
	}
	pt := p
	if t.Target != "" {
		if pt, _ = global.Server.PlayerByName(t.Target); pt == nil{
			output.Error(t.Target + " can't found.")
			return
		}
	}

	DelOp(pt.Name())
	output.Printf("Has been taken op permissions from %s.", pt.Name())
}