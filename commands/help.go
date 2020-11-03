package commands

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/eren5960/essentialsgo/global"
	"math"
)

type Help struct {
	Page int `optional:""`
}

func (t Help) Run(_ cmd.Source, output *cmd.Output) {
	page := int(math.Max(float64(t.Page), 1))
	output.Printf("--- Help Page %d ---", page)
	max := page * 5

	for i, c := range global.Commands {
		if i < max-5 {
			continue
		} else if i >= max {
			break
		}
		output.Printf("%s: %s", c.Name(), c.Description())
	}

	if output.MessageCount() == 1 {
		output.Errorf("There are only have %d pages.", int(math.Round(float64(len(global.Commands))/5)))
	}
}
