package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xerenahmed/essentialsgo/commands/op"
	"github.com/xerenahmed/essentialsgo/global"
)

type Add struct {
	Sub    add
	Amount int `name:"amount"`
}

func (t Add) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
		return
	}

	w := global.Server.World()
	if p, ok := source.(*player.Player); ok {
		w = p.World()
	}
	w.SetTime(w.Time() + t.Amount)
	output.Printf("Added %d to the time", t.Amount)
}

type add string

func (add) SubName() string {
	return "add"
}
