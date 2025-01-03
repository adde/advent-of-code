package main

import (
	"container/heap"
	"fmt"
	"time"

	u "github.com/adde/advent-of-code/utils"
	pq "github.com/adde/advent-of-code/utils/priorityqueue"
)

type Player struct {
	HP, Mana, Damage, Armor int
}

type Spell struct {
	Name                      string
	Cost                      int
	Damage, Heal, Armor, Mana int
	Duration                  int
}

type State struct {
	Wizard, Boss Player
	ManaSpent    int
	Turn         int
	Effects      map[string]int
}

func (s State) GetCost() int {
	return s.ManaSpent
}

const (
	WIZARD_HP   = 50
	WIZARD_MANA = 500
)

var spells = map[string]Spell{
	"Magic Missile": {"Magic Missile", 53, 4, 0, 0, 0, 0},
	"Drain":         {"Drain", 73, 2, 2, 0, 0, 0},
	"Shield":        {"Shield", 113, 0, 0, 7, 0, 6},
	"Poison":        {"Poison", 173, 3, 0, 0, 0, 6},
	"Recharge":      {"Recharge", 229, 0, 0, 0, 101, 5},
}

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	numbers := u.GetIntsFromString(file, false)
	boss := Player{numbers[0], 0, numbers[1], 0}

	fmt.Println("\nPart one:", playGame(boss, false))
	fmt.Println("Part two:", playGame(boss, true))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

// Simulate the game and find the minimum mana spent to win
func playGame(initialBoss Player, isHard bool) int {
	q := &pq.PriorityQueue{State{
		Wizard:    Player{WIZARD_HP, WIZARD_MANA, 0, 0},
		Boss:      initialBoss,
		ManaSpent: 0,
		Turn:      0,
		Effects:   make(map[string]int),
	}}
	heap.Init(q)

	for q.Len() > 0 {
		state := heap.Pop(q).(State)

		// If Boss is dead, return mana spent
		if state.Boss.HP <= 0 {
			return state.ManaSpent
		}

		// If Wizard is dead, continue
		if state.Wizard.HP <= 0 {
			continue
		}

		// Player turn
		if state.Turn == 0 {
			// If hard mode, reduce Wizard HP by 1 each turn
			if isHard {
				state.Wizard.HP--
				if state.Wizard.HP <= 0 {
					continue
				}
			}

			// Try casting spells
			for _, spell := range spells {
				// Copy state
				newState := copyState(state)
				newState.Turn = 1

				// Apply effects
				applyEffects(&newState)

				// If spell is too expensive, continue
				if newState.Wizard.Mana <= spell.Cost {
					continue
				}

				// Apply mana cost
				newState.Wizard.Mana -= spell.Cost
				newState.ManaSpent += spell.Cost

				// Cast spell
				switch spell.Name {
				case "Magic Missile":
					newState.Boss.HP -= spell.Damage
				case "Drain":
					newState.Boss.HP -= spell.Damage
					newState.Wizard.HP += spell.Heal
				default:
					if _, ok := newState.Effects[spell.Name]; !ok {
						newState.Effects[spell.Name] = spell.Duration
					}
				}

				heap.Push(q, newState)
			}
		} else { // Boss turn
			// Copy state
			newState := copyState(state)
			newState.Turn = 0

			// Apply effects
			applyEffects(&newState)

			// Boss attack
			damage := newState.Boss.Damage - newState.Wizard.Armor
			if damage < 1 {
				damage = 1
			}
			newState.Wizard.HP -= damage

			heap.Push(q, newState)
		}
	}

	return 0
}

// Create a new state object with the same values
func copyState(state State) State {
	newState := State{
		Wizard:    state.Wizard,
		Boss:      state.Boss,
		ManaSpent: state.ManaSpent,
		Turn:      state.Turn,
		Effects:   make(map[string]int),
	}

	for k, v := range state.Effects {
		newState.Effects[k] = v
	}

	return newState
}

// Apply spell effects and remove expired effects
func applyEffects(state *State) {
	// Reset armor before applying effects
	state.Wizard.Armor = 0

	// Apply active effects
	for name, turns := range state.Effects {
		switch name {
		case "Shield":
			state.Wizard.Armor = spells[name].Armor
		case "Poison":
			state.Boss.HP -= spells[name].Damage
		case "Recharge":
			state.Wizard.Mana += spells[name].Mana
		}
		state.Effects[name] = turns - 1

		// Remove expired effects
		if state.Effects[name] <= 0 {
			delete(state.Effects, name)
		}
	}
}
