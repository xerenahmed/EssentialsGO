package commands
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)

type Op struct {
	Target string `optional:""`
}

func (t Op) Run(source cmd.Source, output *cmd.Output) {
	p := source.(*player.Player)
	pt := p
	if t.Target != "" {
		if pt, _ = Server.PlayerByName(t.Target); pt == nil{
			output.Error(t.Target + " can't found.")
			return
		}
	}

	pt.ShowCoordinates()

	output.Printf("Setted op the %s.", pt.Name())
}