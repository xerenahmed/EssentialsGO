package commands

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/go-gl/mathgl/mgl64"
	"strconv"
)

type Teleport struct {
	X, Y, Z  float64
}

func (t Teleport) Run(source cmd.Source, output *cmd.Output) {
	p, _ := source.(*player.Player)
	p.Teleport(mgl64.Vec3{t.X, t.Y, t.Z})
	output.Print("Teleported to X: " + strconv.Itoa(int(t.X)) + " Y: " + strconv.Itoa(int(t.Y)) + " Z: " + strconv.Itoa(int(t.Z)))
}