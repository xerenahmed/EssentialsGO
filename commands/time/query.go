package time

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"reflect"
)

type Query struct {
	Sub  query
	Time timeQuery `name:"time"`
}

func (t Query) Run(source cmd.Source, output *cmd.Output) {
	// TODO
}

type query string

func (query) SubName() string {
	return "query"
}

type timeQuery string

func (timeQuery) Type() string {
	return "TimeQuery"
}

func (timeQuery) Options() []string {
	return []string{"day", "daytime", "gametime"}
}

func (timeQuery) SetOption(o string, r reflect.Value) {}
