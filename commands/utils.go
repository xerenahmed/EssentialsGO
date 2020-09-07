package commands

import (
	"fmt"
	"github.com/df-mc/dragonfly/dragonfly/player"
	"github.com/df-mc/dragonfly/dragonfly/world/gamemode"
	"io/ioutil"
	"os"
	"strings"
)

func StringToGameMode(arg string) gamemode.GameMode{
	return map[string]gamemode.GameMode{
		"0": gamemode.Survival{}, "s": gamemode.Survival{}, "survival": gamemode.Survival{},
		"1": gamemode.Creative{}, "c": gamemode.Creative{}, "creative": gamemode.Creative{},
		"2": gamemode.Adventure{}, "a": gamemode.Adventure{}, "adventure": gamemode.Creative{},
	}[arg]
}

func GameModeToName(gm gamemode.GameMode) string{
	return map[gamemode.GameMode]string{
		gamemode.Survival{}: "survival",
		gamemode.Creative{}: "creative",
		gamemode.Adventure{}: "adventure",
	}[gm]
}

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

func IsOp(p *player.Player) bool{
	for _, p_ := range ops {
		if p_ == strings.ToLower(p.Name()){
			return true
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
		ops = SliceUnique(strings.Split(string(data), "\n"))
	}
}

func SaveOps(){
	ioutil.WriteFile("ops.txt", []byte(strings.Join(SliceUnique(ops), "\n")), 0655)
}

func SliceUnique(slice []string) []string {
	keys := make(map[string]bool)
	var list []string
	for _, entry := range slice {
		if _, value := keys[entry]; !value && entry != "" {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}