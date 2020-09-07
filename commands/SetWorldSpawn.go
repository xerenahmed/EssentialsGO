package commands
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world"
	"github.com/eren5960/essentialsgo/commands/op"
	"github.com/eren5960/essentialsgo/global"
)

type SetWorldSpawn struct {}

func (t SetWorldSpawn) Run(source cmd.Source, output *cmd.Output) {
	p, _ := source.(*player.Player)
	if !op.IsOp(p){
		output.Error("You don't have permission for this command.")
		return
	}
	pos := p.Position()
	bp := world.BlockPos{int(pos.X()), int(pos.Y()), int(pos.Z())}

	global.Server.World().SetSpawn(bp)
	output.Printf("Set the world spawn point to (%d, %d, %d)", bp.X(), bp.Y(), bp.Z())
}