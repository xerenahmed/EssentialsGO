package time

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/global"
	"reflect"
)

type Add struct {
	Sub add
	Amount int `name:"amount"`
}

func (t Add) Run(source cmd.Source, output *cmd.Output) {
	w := global.Server.World()
	if p, ok := source.(*player.Player); ok {
		w = p.World()
	}
	w.SetTime(w.Time() + t.Amount)
	output.Printf("Added %s to the time", t.Amount)
}

type add string

func (add) Type() string {
	return "add"
}

func (add) Options() []string {
	return []string{"add"}
}

func (add) SetOption(string, reflect.Value) {}