package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	health int
}

func takeDamage(player *Player) {
	rand.Seed(time.Now().UnixNano())
	dmg := rand.Intn(10)
	player.health -= dmg
}

func main() {
	player := &Player{health: 100}
	fmt.Println("Player health before: ", player.health)
	takeDamage(player)
	fmt.Println("Player health after: ", player.health)
}
