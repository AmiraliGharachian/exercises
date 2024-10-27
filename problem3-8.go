package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Weapon struct {
	damage          []int
	strengthReq     int
	agilityReq      int
	intelligenceReq int
}

type Consumable struct {
	duration        int
	hitpointsEffect int
	strengthEffect  int
	agilityEffect   int
	intellectEffect int
}

type Gopher struct {
	hitpoints      int
	weapons        []Weapon
	inventory      []Consumable
	strength       int
	agility        int
	intellect      int
	coins          int
	equippedWeapon *Weapon
}

func randomInRange(min, max int) int {
	return rand.Intn(max-min+1) + min
}

func (g *Gopher) attack(other *Gopher) {
	damage := 1
	if g.equippedWeapon != nil {
		damage = randomInRange(g.equippedWeapon.damage[0], g.equippedWeapon.damage[1])
	}
	other.hitpoints -= damage
	fmt.Printf("Gopher attacks for %d damage. Other gopher HP: %d\n", damage, other.hitpoints)
}

func (g *Gopher) work() {
	earnings := randomInRange(5, 15)
	g.coins += earnings
	fmt.Printf("Gopher goes to work and earns %d coins!\n", earnings)
}

func (g *Gopher) buy(weapon Weapon) {
	if g.coins >= 10 && weapon.strengthReq <= g.strength && weapon.agilityReq <= g.agility && weapon.intelligenceReq <= g.intellect {
		g.coins -= 10
		g.equippedWeapon = &weapon
		fmt.Println("Gopher bought a weapon!")
	} else {
		fmt.Println("Gopher cannot afford this weapon or does not meet the requirements.")
	}
}

func (g *Gopher) train(skill string) {
	if g.coins < 5 {
		fmt.Println("Not enough coins to train.")
		return
	}
	g.coins -= 5
	switch skill {
	case "strength":
		g.strength += 2
		fmt.Println("Gopher trains strength!")
	case "agility":
		g.agility += 2
		fmt.Println("Gopher trains agility!")
	case "intellect":
		g.intellect += 2
		fmt.Println("Gopher trains intellect!")
	default:
		fmt.Println("Unknown skill.")
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	gopher1 := Gopher{hitpoints: 30, coins: 20}
	gopher2 := Gopher{hitpoints: 30, coins: 20}

	knife := Weapon{damage: []int{2, 3}, strengthReq: 0, agilityReq: 0, intelligenceReq: 0}
	ninjaku := Weapon{damage: []int{1, 7}, strengthReq: 0, agilityReq: 2, intelligenceReq: 0}

	fmt.Println("Welcome to Gopher RPG")
	gopher1.buy(knife)
	gopher2.buy(ninjaku)

	for gopher1.hitpoints > 0 && gopher2.hitpoints > 0 {

		fmt.Println("Gopher 1's turn")
		gopher1.attack(&gopher2)
		if gopher2.hitpoints <= 0 {
			fmt.Println("Gopher 2 is defeated!")
			break
		}

		fmt.Println("Gopher 2's turn")
		gopher2.attack(&gopher1)
		if gopher1.hitpoints <= 0 {
			fmt.Println("Gopher 1 is defeated!")
			break
		}
	}
}
