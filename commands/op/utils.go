package op

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/cmd"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/eren5960/essentialsgo/commands/utils"
	"github.com/eren5960/essentialsgo/console"
	"io/ioutil"
	"os"
	"strings"
)

var ops = make([]string, 0)

func GetOps() []string{
	return ops
}

func AddOp(p string){
	ops = append(ops, strings.ToLower(p))
	SaveOps()
}

func DelOp(p string){
	var newOps []string
	for _, p_ := range ops {
		if p_ != strings.ToLower(p) {
			newOps = append(newOps, p_)
		}
	}

	ops = newOps
	SaveOps()
}

func IsOp(s cmd.Source) bool{
	if _, ok := s.(*console.Console); ok {
		return true
	}
	if p, ok := s.(*player.Player); ok {
		for _, p_ := range ops {
			if p_ == strings.ToLower(p.Name()){
				return true
			}
		}
	}

	return false
}

func LoadOps(){
	if _, err := os.Stat("ops.txt"); os.IsNotExist(err) {
		if _, err = os.Create("ops.txt"); err != nil {
			fmt.Println("Error on creating ops.txt: " + err.Error())
			return
		}
	}
	data, err := ioutil.ReadFile("ops.txt")
	if err == nil {
		ops = utils.SliceUnique(strings.Split(string(data), "\n"))
	}
}

func SaveOps(){
	_ = ioutil.WriteFile("ops.txt", []byte(strings.Join(utils.SliceUnique(ops), "\n")), 0655)
}