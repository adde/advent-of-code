package main

import (
	"fmt"
	"math"
	"time"

	u "github.com/adde/advent-of-code/utils"
)

type Player struct {
	HP, Damage, Armor int
}

type Item struct {
	Name   string
	Cost   int
	Damage int
	Armor  int
}

var weapons = []Item{
	{"Dagger", 8, 4, 0},
	{"Shortsword", 10, 5, 0},
	{"Warhammer", 25, 6, 0},
	{"Longsword", 40, 7, 0},
	{"Greataxe", 74, 8, 0},
}

var armor = []Item{
	{"None", 0, 0, 0},
	{"Leather", 13, 0, 1},
	{"Chainmail", 31, 0, 2},
	{"Splintmail", 53, 0, 3},
	{"Bandedmail", 75, 0, 4},
	{"Platemail", 102, 0, 5},
}

var rings = []Item{
	{"None", 0, 0, 0},
	{"Damage +1", 25, 1, 0},
	{"Damage +2", 50, 2, 0},
	{"Damage +3", 100, 3, 0},
	{"Defense +1", 20, 0, 1},
	{"Defense +2", 40, 0, 2},
	{"Defense +3", 80, 0, 3},
}

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	numbers := u.GetIntsFromString(file, false)
	boss := Player{numbers[0], numbers[1], numbers[2]}

	fmt.Println("\nPart one:", partOne(boss))
	fmt.Println("Part two:", partTwo(boss))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(initialBoss Player) int {
	lowestCost := math.MaxInt64

	for _, weapon := range weapons {
		for _, armor := range armor {
			for _, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1 == ring2 && ring1.Name != "None" {
						continue
					}

					cost := weapon.Cost + armor.Cost + ring1.Cost + ring2.Cost
					damage := weapon.Damage + armor.Damage + ring1.Damage + ring2.Damage
					armor := weapon.Armor + armor.Armor + ring1.Armor + ring2.Armor

					player := Player{100, damage, armor}
					boss := initialBoss

					for {
						boss.HP -= u.Max(1, player.Damage-boss.Armor)
						if boss.HP <= 0 {
							if cost < lowestCost {
								lowestCost = cost
							}
							break
						}

						player.HP -= u.Max(1, boss.Damage-player.Armor)
						if player.HP <= 0 {
							break
						}
					}
				}
			}
		}
	}

	return lowestCost
}

func partTwo(initialBoss Player) int {
	highestCost := 0

	for _, weapon := range weapons {
		for _, armor := range armor {
			for _, ring1 := range rings {
				for _, ring2 := range rings {
					if ring1 == ring2 && ring1.Name != "None" {
						continue
					}

					cost := weapon.Cost + armor.Cost + ring1.Cost + ring2.Cost
					damage := weapon.Damage + armor.Damage + ring1.Damage + ring2.Damage
					armor := weapon.Armor + armor.Armor + ring1.Armor + ring2.Armor

					player := Player{100, damage, armor}
					boss := initialBoss

					for {
						boss.HP -= u.Max(1, player.Damage-boss.Armor)
						if boss.HP <= 0 {
							break
						}

						player.HP -= u.Max(1, boss.Damage-player.Armor)
						if player.HP <= 0 {
							if cost > highestCost {
								highestCost = cost
							}
							break
						}
					}
				}
			}
		}
	}

	return highestCost
}
