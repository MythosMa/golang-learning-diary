package main

import (
	"mythosma.poke/module"

)

func main() {
	players := []string{"刘德华", "张学友", "黎明", "郭富城"}
    game := new(core.Game)
	game.NewGame(players)
}