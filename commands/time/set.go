package time

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/xerenahmed/essentialsgo/commands/op"
	"github.com/xerenahmed/essentialsgo/global"
	"reflect"
)

type Set struct {
	Sub    set
	Amount int `name:"amount"`
}

type SetTimeSpec struct {
	Sub  set
	Time spec `name:"time"`
}

func setTime(source cmd.Source, output *cmd.Output, t int) {
	if !op.IsOp(source) {
		output.Error("You don't have permission for this command.")
		return
	}

	w := global.Server.World()
	if p, ok := source.(*player.Player); ok {
		w = p.World()
	}
	w.SetTime(t)
	output.Printf("Set the time to %d", t)
}

func (t Set) Run(source cmd.Source, output *cmd.Output) {
	setTime(source, output, t.Amount)
}

func (t SetTimeSpec) Run(source cmd.Source, output *cmd.Output) {
	setTime(source, output, int(t.Time))
}

type set string

func (set) SubName() string {
	return "set"
}

type spec int

func (spec) Type() string {
	return "TimeSpec"
}

func (spec) Options() []string {
	return []string{"day", "night", "noon", "midnight", "sunrise", "sunset"}
}

func (spec) SetOption(o string, r reflect.Value) {
	r.SetInt(map[string]int64{
		"day": 1000, "night": 13000, "noon": 6000, "midnight": 18000, "sunrise": 23000, "sunset": 12000,
	}[o])
}
