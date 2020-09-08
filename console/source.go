package console
// Eren5960 <ahmederen123@gmail.com>
import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
)

type Console struct {}

func (Console) SendCommandOutput(output *cmd.Output){
	for _, m := range output.Messages(){
		fmt.Println(m)
	}

	for _, e := range output.Errors(){
		fmt.Println(e.Error())
	}
}