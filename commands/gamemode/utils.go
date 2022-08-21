package gamemode

import (
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
)

func StringToGameMode(arg string) world.GameMode {
	return map[string]world.GameMode{
		"0": world.GameModeSurvival, "s": world.GameModeSurvival, "survival": world.GameModeSurvival,
		"1": world.GameModeCreative, "c": world.GameModeCreative, "creative": world.GameModeCreative,
		"2": world.GameModeAdventure, "a": world.GameModeAdventure, "adventure": world.GameModeAdventure,
		"3": world.GameModeSpectator, "v": world.GameModeSpectator, "spectator": world.GameModeSpectator,
	}[arg]
}

func GameModeToName(gm world.GameMode) string {
	return map[world.GameMode]string{
		world.GameModeSurvival:  "survival",
		world.GameModeCreative:  "creative",
		world.GameModeAdventure: "adventure",
		world.GameModeSpectator: "spectator",
	}[gm]
}

type mode string

func (mode) Type() string {
	return "mode"
}

func (mode) Options(source cmd.Source) []string {
	return []string{"0", "1", "2", "3", "s", "c", "a", "v", "survival", "creative", "adventure", "spectator"}
}
