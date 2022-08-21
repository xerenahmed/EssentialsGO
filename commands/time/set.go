package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/player"
	"github.com/xerenahmed/essentialsgo/commands/op"
	"github.com/xerenahmed/essentialsgo/global"
)

type Set struct {
	Sub    cmd.SubCommand `cmd:"set"`
	Amount int            `name:"amount"`
}

type SetTimeSpec struct {
	Sub  cmd.SubCommand `cmd:"set"`
	Time spec           `name:"time"`
}

func setTime(source cmd.Source, output *cmd.Output, t int) {
	if !op.IsOp(source) {
		output.Error("You don't have permission to run this command.")
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
	tf := map[spec]int64{
		"day": 1000, "night": 13000, "noon": 6000, "midnight": 18000, "sunrise": 23000, "sunset": 12000,
	}[t.Time]
	setTime(source, output, int(tf))
}

type spec string

func (spec) Type() string {
	return "TimeSpec"
}

func (spec) Options(source cmd.Source) []string {
	return []string{"day", "night", "noon", "midnight", "sunrise", "sunset"}
}
