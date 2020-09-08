package commands

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/commands/op"
	"github.com/eren5960/essentialsgo/commands/utils"
	"github.com/eren5960/essentialsgo/console"
	"github.com/go-gl/mathgl/mgl64"
)

type TeleportXYZ struct {
	X, Y, Z  float64
}

func (t TeleportXYZ) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*console.Console); ok{
		output.Error("Use in game!")
		return
	}
	p, _ := source.(*player.Player)
	if !op.IsOp(p){
		output.Error("You don't have permission for this command.")
		return
	}
	p.Teleport(mgl64.Vec3{t.X, t.Y, t.Z})
	output.Printf("Teleported to X: %d Y: %d Z: %d", int(t.X), int(t.Y), int(t.Z) )
}

type TeleportPlayer struct {
	Target string
	Target2 string `optional:""`
}

func (t TeleportPlayer) Run(source cmd.Source, output *cmd.Output) {
	if t.Target != "" {
		to, ok := utils.PlayerByName(t.Target)
		if !ok {
			output.Errorf("%s not found", t.Target)
			return
		}
		if t.Target2 == "" {
			if p, ok := source.(*player.Player); ok {
				t.TeleportAnotherPlayer(p, to)
			}
		} else {
			p, ok := utils.PlayerByName(t.Target2)
			if !ok {
				output.Errorf("%s not found.", t.Target2)
				return
			}
			t.TeleportPlayerToAnotherPlayer(to, p)
			output.Printf("The player %s has been teleported to %s.", to.Name(), p.Name())
			return
		}
	}

	output.Error("Usage: /teleport <Player: string> <Player: string, optional>")
}

func (TeleportPlayer) TeleportAnotherPlayer(p, to *player.Player) {
	p.Teleport(to.Position())
	p.Message(fmt.Sprintf("You teleported to %s.", to.Name()))
}

func (TeleportPlayer) TeleportPlayerToAnotherPlayer(to, p *player.Player) {
	to.Teleport(p.Position())
	p.Message(fmt.Sprintf("The player %s has been teleported to your side.", to.Name()))
	to.Message(fmt.Sprintf("You have been teleported here by the %s.", p.Name()))
}