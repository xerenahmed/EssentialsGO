package commands

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/console"
)

type XYZ struct{}

var opened = map[string]bool{}

func (t XYZ) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*console.Console); ok {
		output.Error("Use in game!")
		return
	}
	var msg string
	p, _ := source.(*player.Player)
	id := p.Name()

	opened[id] = !opened[id]

	if opened[id] {
		p.ShowCoordinates()
		msg = "shown"
	} else {
		p.HideCoordinates()
		msg = "hidden"
	}

	output.Print("Coordinates " + msg + ".")
}
