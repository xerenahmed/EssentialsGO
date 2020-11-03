package commands

import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/eren5960/essentialsgo/commands/op"
	"github.com/eren5960/essentialsgo/global"
)

type Stop struct{}

func (t Stop) Run(source cmd.Source, output *cmd.Output) {
	if !op.IsOp(source) {
		output.Error("You don't have permission for this command.")
		return
	}
	output.Printf("Stopping server.")
	if err := global.Server.Close(); err != nil {
		output.Printf("error shutting down server: %v", err)
	}
}
