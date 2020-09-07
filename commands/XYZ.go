package commands

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
)
// Eren5960 <ahmederen123@gmail.com>
type XYZ struct {}

var opened = map[string]bool{}

func (t XYZ) Run(source cmd.Source, output *cmd.Output) {
	var msg string
	p, _ := source.(*player.Player)
	id := p.Name()

	if opened[id] {
		p.HideCoordinates()
		opened[id] = false
		msg = "hidden"
		return
	}
	p.ShowCoordinates()
	opened[id] = true
	msg = "shown"

	output.Print("Coordinates " + msg + ".")
}