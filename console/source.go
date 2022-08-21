package console

import (
	"fmt"
	"github.com/df-mc/dragonfly/server/cmd"
	"github.com/df-mc/dragonfly/server/world"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/xerenahmed/essentialsgo/global"
)

type Console struct{}

func (Console) SendCommandOutput(output *cmd.Output) {
	for _, m := range output.Messages() {
		fmt.Println(m)
	}

	for _, e := range output.Errors() {
		fmt.Println(e.Error())
	}
}

func (Console) Name() string {
	return "Console"
}

func (Console) Position() mgl64.Vec3 {
	return mgl64.Vec3{}
}

func (Console) World() *world.World {
	return global.Server.World()
}
