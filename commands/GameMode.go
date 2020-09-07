package commands

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	"github.com/sandertv/gophertunnel/minecraft/text"
	"reflect"
)

type GameMode struct {
	GameMode mode
	Target   target
}

// Run ...
func (t GameMode) Run(source cmd.Source, output *cmd.Output) {
	spew.Dump(t)
	p := source.(*player.Player)
	mode := map[mode]gamemode.GameMode{
		"0": gamemode.Survival{}, "s": gamemode.Survival{},
		"1": gamemode.Creative{}, "c": gamemode.Creative{},
		"2": gamemode.Adventure{}, "a": gamemode.Adventure{},
	}[t.GameMode]

	p.SetGameMode(mode)
	output.Printf(text.Green()("Changed gamemode."))
}

type mode string

func (mode) Type() string {
	return "mode"
}

func (mode) Options() []string {
	return []string{"0", "1", "2", "s", "c", "a"}
}

func (mode) SetOption(option string, r reflect.Value) {
	r.SetString(option)
}

type target string

func (target) Type() string {
	return "target"
}

func (target) Options() []string {
	var players []string

	for _, p := range Server.Players() {
		players = append(players, p.Name())
	}

	return players
}

func (target) SetOption(option string, r reflect.Value) {
	r.SetString(option)
}