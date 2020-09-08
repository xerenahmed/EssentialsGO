package time

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/eren5960/essentialsgo/global"
	"reflect"
)

type Query struct {
	Sub query
}

func (t Query) Run(source cmd.Source, output *cmd.Output) {
	global.Server.World().Time()
}

type query string

func (query) Type() string {
	return "query"
}

func (query) Options() []string {
	return []string{"query"}
}

func (query) SetOption(string, reflect.Value) {}