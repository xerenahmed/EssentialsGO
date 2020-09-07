package commands
// Eren5960 <ahmederen123@gmail.com>
import (
	"github.com/df-mc/dragonfly/dragonfly/cmd"
)

type Stop struct {}

func (t Stop) Run(source cmd.Source, output *cmd.Output) {
	output.Printf("Stopping server.")
	if err := Server.Close(); err != nil {
		output.Printf("error shutting down server: %v", err)
	}
}