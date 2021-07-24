package commands

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/xerenahmed/essentialsgo/commands/op"
	"github.com/xerenahmed/essentialsgo/console"
	"github.com/go-gl/mathgl/mgl64"
)

type TeleportXYZ struct {
	X, Y, Z float64
}

func (t TeleportXYZ) Run(source cmd.Source, output *cmd.Output) {
	if _, ok := source.(*console.Console); ok {
		output.Error("Use in game!")
		return
	}
	p, _ := source.(*player.Player)
	if !op.IsOp(p) {
		output.Error("You don't have permission to run this command.")
		return
	}
	p.Teleport(mgl64.Vec3{t.X, t.Y, t.Z})
	output.Printf("Teleported to X: %d Y: %d Z: %d", int(t.X), int(t.Y), int(t.Z))
}

type TeleportPlayer struct {
	Player []cmd.Target
	Target []cmd.Target `optional:""`
}

func (t TeleportPlayer) Run(source cmd.Source, output *cmd.Output) {
	if len(t.Player) < 1 {
		output.Error("Usage: /teleport <Player: target> [Target: target]")
	}

	for _, target := range t.Player {
		if tp1, ok := target.(*player.Player); ok {
			if len(t.Target) < 1 {
				if p, ok := source.(*player.Player); ok {
					t.TeleportAnotherPlayer(p, tp1)
				}
			} else if tp2, ok := t.Target[0].(*player.Player); ok {
				t.TeleportPlayerToAnotherPlayer(tp1, tp2)
				output.Printf("%s has been teleported to %s.", tp1.Name(), tp2.Name())
			} else {
				output.Errorf("Second target is not a player!")
			}
		} else {
			output.Errorf("First target is not a player!")
		}
	}
}

func (TeleportPlayer) TeleportAnotherPlayer(p, to *player.Player) {
	p.Teleport(to.Position())
	p.Message(fmt.Sprintf("You teleported to %s.", to.Name()))
}

func (TeleportPlayer) TeleportPlayerToAnotherPlayer(to, p *player.Player) {
	to.Teleport(p.Position())
	p.Message(fmt.Sprintf("%s has been teleported to your side.", to.Name()))
	to.Message(fmt.Sprintf("You have been teleported here by %s.", p.Name()))
}
