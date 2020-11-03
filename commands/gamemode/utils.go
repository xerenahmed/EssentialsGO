package gamemode

import (
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	"reflect"
)

func StringToGameMode(arg string) gamemode.GameMode {
	return map[string]gamemode.GameMode{
		"0": gamemode.Survival{}, "s": gamemode.Survival{}, "survival": gamemode.Survival{},
		"1": gamemode.Creative{}, "c": gamemode.Creative{}, "creative": gamemode.Creative{},
		"2": gamemode.Adventure{}, "a": gamemode.Adventure{}, "adventure": gamemode.Creative{},
	}[arg]
}

func GameModeToName(gm gamemode.GameMode) string {
	return map[gamemode.GameMode]string{
		gamemode.Survival{}:  "survival",
		gamemode.Creative{}:  "creative",
		gamemode.Adventure{}: "adventure",
	}[gm]
}

type mode string

func (mode) Type() string {
	return "mode"
}

func (mode) Options() []string {
	return []string{"0", "1", "2", "s", "c", "a", "survival", "creative", "adventure"}
}

func (mode) SetOption(option string, r reflect.Value) {
	r.SetString(option)
}
