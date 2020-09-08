package time
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/global"
	"reflect"
	"strconv"
)

type Set struct {
	Sub set
	Amount int `name:"amount"`
}

type SetTimeSpec struct {
	Sub set
	TimeSpec spec
}

func setTime(source cmd.Source, output *cmd.Output, t int){
	w := global.Server.World()
	if p, ok := source.(*player.Player); ok {
		w = p.World()
	}
	w.SetTime(t)
	output.Printf("Set the time to %d", t)
}

func (t Set) Run(source cmd.Source, output *cmd.Output) {
	setTime(source, output, t.Amount)
	output.Printf("")
}

func (t SetTimeSpec) Run(source cmd.Source, output *cmd.Output) {
	tm, _ := strconv.Atoi(string(t.TimeSpec))
	setTime(source, output, tm)
}

type set string

func (set) Type() string {
	return "set"
}

func (set) Options() []string {
	return []string{"set"}
}

func (set) SetOption(string, reflect.Value) {}

type spec string

func (spec) Type() string {
	return "spec"
}

func (spec) Options() []string {
	return []string{"day", "night", "noon", "midnight", "sunrise", "sunset"}
}

func (spec) SetOption(o string, r reflect.Value) {
	r.SetString(map[string]string{
		"day": "1000", "night": "13000", "noon": "6000", "midnight": "18000", "sunrise": "23000", "sunset": "12000",
	}[o])
}