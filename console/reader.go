package console

import (
	"bufio"
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"os"
	"strings"
	"time"
)

func StartConsole() {
	go func() {
		time.Sleep(time.Second * 1)
		source := &Console{}
		fmt.Println("Type help for commands.")
		// I don't use fmt.Scan() because the fmt package intentionally filters out whitespaces, this is how it is implemented.
		scanner := bufio.NewScanner(os.Stdin)
		for {
			if scanner.Scan() {
				commandString := scanner.Text()

				split := strings.Split(commandString, " ")
				if len(commandString) == 0 {
					continue
				}
				commandName := split[0]
				command, ok := cmd.ByAlias(commandName)

				if !ok {
					output := &cmd.Output{}
					output.Errorf("Unknown command '%v'", commandName)
					for _, e := range output.Errors() {
						fmt.Println(e)
					}
					continue
				}

				command.Execute(strings.TrimPrefix(strings.TrimPrefix(commandString, commandString), " "), source)
			}
		}
	}()
}
