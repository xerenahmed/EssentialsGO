package time

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/df-mc/dragonfly/server/cmd"
)

type Query struct {
	Sub  cmd.SubCommand `cmd:"query"`
	Time timeQuery      `name:"time"`
}

func (t Query) Run(source cmd.Source, output *cmd.Output) {
	spew.Dump(t, source)
}

type timeQuery string

func (timeQuery) Type() string {
	return "TimeQuery"
}

func (timeQuery) Options(source cmd.Source) []string {
	return []string{"day", "daytime", "gametime"}
}
