package main

import (
	"fmt"
	"strings"
	"time"

	u "github.com/adde/advent-of-code/utils"
	g "github.com/adde/advent-of-code/utils/grid"
	"github.com/adde/advent-of-code/utils/queue"
)

const (
	ROBOT      = '@'
	WALL       = '#'
	BOX        = 'O'
	BOX_LEFT   = '['
	BOX_RIGHT  = ']'
	SPACE      = '.'
	MULTIPLIER = 100
)

type Box struct {
	Pos  g.Point
	Char rune
}

type BoxRow struct {
	Row   int
	Boxes []Box
}

func main() {
	startTime := time.Now()
	file := u.ReadAll("input.txt")
	fileParts := strings.Split(file, "\n\n")
	warehouseP1 := g.CreateFromLines(strings.Split(fileParts[0], "\n"))
	warehouseP2 := g.CreateFromLines(strings.Split(fileParts[0], "\n"))
	moves := []rune(strings.ReplaceAll(fileParts[1], "\n", ""))

	fmt.Println("\nPart one:", partOne(warehouseP1, moves))
	fmt.Println("Part two:", partTwo(warehouseP2, moves))
	fmt.Printf("\nExecution time: %s\n\n", time.Since(startTime))
}

func partOne(warehouse g.Grid, moves []rune) int {
	robotPos := getRobotPos(warehouse)
	dirs := getDirections(moves)

	for _, dir := range dirs {
		newR, newC := robotPos.Row+dir.Row, robotPos.Col+dir.Col

		// If we hit a wall, continue to the next instruction
		if warehouse[newR][newC] == WALL {
			continue
		}

		// If we hit an empty space, move the robot
		if warehouse[newR][newC] == SPACE {
			warehouse[robotPos.Row][robotPos.Col] = SPACE
			robotPos = g.Point{Row: newR, Col: newC}
			warehouse[robotPos.Row][robotPos.Col] = ROBOT
		}

		// If we hit a box, try to move both the robot and the boxes
		if warehouse[newR][newC] == BOX {
			boxesToMove := []g.Point{{Row: newR, Col: newC}}
			newR2, newC2 := newR+dir.Row, newC+dir.Col

			for {
				// If we hit a wall, we can't move the boxes
				if warehouse[newR2][newC2] == WALL {
					break
				}

				// If we hit an empty space, move the robot
				// and the boxes one step in the direction
				if warehouse[newR2][newC2] == SPACE {
					// Move the robot
					warehouse[robotPos.Row][robotPos.Col] = SPACE
					robotPos = g.Point{Row: newR, Col: newC}
					warehouse[robotPos.Row][robotPos.Col] = ROBOT

					// Move the boxes
					for _, box := range boxesToMove {
						boxNewRow, boxNewCol := box.Row+dir.Row, box.Col+dir.Col
						warehouse[boxNewRow][boxNewCol] = BOX
					}

					break
				}

				// If we hit another box, add it to the list of boxes
				if warehouse[newR2][newC2] == BOX {
					boxesToMove = append(boxesToMove, g.Point{Row: newR2, Col: newC2})
				}

				newR2, newC2 = newR2+dir.Row, newC2+dir.Col
			}
		}
	}

	return getBoxGPSCoordinates(warehouse, false)
}

func partTwo(warehouse g.Grid, moves []rune) int {
	// Resize the warehouse
	warehouse = doubleWarehouseWidth(warehouse)
	robotPos := getRobotPos(warehouse)
	dirs := getDirections(moves)

	for _, dir := range dirs {
		newR, newC := robotPos.Row+dir.Row, robotPos.Col+dir.Col

		// If we hit a wall, continue to the next instruction
		if warehouse[newR][newC] == WALL {
			continue
		}

		// If we hit an empty space, move the robot
		if warehouse[newR][newC] == SPACE {
			warehouse[robotPos.Row][robotPos.Col] = SPACE
			robotPos = g.Point{Row: newR, Col: newC}
			warehouse[robotPos.Row][robotPos.Col] = ROBOT
		}

		// If we hit a box, try to move both the robot and the boxes
		if warehouse[newR][newC] == BOX_LEFT || warehouse[newR][newC] == BOX_RIGHT {
			// When moving up or down we need to handle double width boxes
			if dir.Col == 0 {
				robotPos = moveDoubleBoxes(warehouse, robotPos, g.Point{Row: newR, Col: newC}, dir)
				continue
			}

			// Keep track of the boxes we are moving
			boxesToMove := []Box{{
				Pos:  g.Point{Row: newR, Col: newC},
				Char: warehouse[newR][newC],
			}}
			newR2, newC2 := newR+dir.Row, newC+dir.Col

			for {
				// If we hit a wall, we can't move the boxes
				if warehouse[newR2][newC2] == WALL {
					break
				}

				// If we hit an empty space, move the robot
				// and the boxes one step in the direction
				if warehouse[newR2][newC2] == SPACE {
					// Move the robot
					warehouse[robotPos.Row][robotPos.Col] = SPACE
					robotPos = g.Point{Row: newR, Col: newC}
					warehouse[robotPos.Row][robotPos.Col] = ROBOT

					// Move the boxes
					for _, box := range boxesToMove {
						boxNewRow, boxNewCol := box.Pos.Row+dir.Row, box.Pos.Col+dir.Col
						warehouse[boxNewRow][boxNewCol] = box.Char
					}

					break
				}

				// If we hit another box, add it to the list of boxes
				if warehouse[newR2][newC2] == BOX_LEFT || warehouse[newR2][newC2] == BOX_RIGHT {
					boxesToMove = append(boxesToMove, Box{
						Pos:  g.Point{Row: newR2, Col: newC2},
						Char: warehouse[newR2][newC2],
					})
				}

				newR2, newC2 = newR2+dir.Row, newC2+dir.Col
			}
		}
	}

	return getBoxGPSCoordinates(warehouse, true)
}

// Handle moving boxes that are double in width, moving up or down
func moveDoubleBoxes(warehouse g.Grid, robotPos, boxPos, dir g.Point) g.Point {
	// Initialize the first box row
	boxRow := BoxRow{
		Row: boxPos.Row,
		Boxes: []Box{{
			Pos:  g.Point{Row: boxPos.Row, Col: boxPos.Col},
			Char: warehouse[boxPos.Row][boxPos.Col],
		},
			{
				Pos:  g.Point{Row: boxPos.Row, Col: boxPos.Col + 1},
				Char: warehouse[boxPos.Row][boxPos.Col+1],
			},
		},
	}
	if warehouse[boxPos.Row][boxPos.Col] == BOX_RIGHT {
		boxRow = BoxRow{Row: boxPos.Row, Boxes: []Box{{
			Pos:  g.Point{Row: boxPos.Row, Col: boxPos.Col},
			Char: warehouse[boxPos.Row][boxPos.Col],
		}, {
			Pos:  g.Point{Row: boxPos.Row, Col: boxPos.Col - 1},
			Char: warehouse[boxPos.Row][boxPos.Col-1],
		}}}
	}

	queue := queue.New(boxRow)
	boxesToMove := []BoxRow{boxRow}

	// Run BFS to find all boxes that can be moved
	for !queue.IsEmpty() {
		current := queue.Pop()
		newBoxRow := BoxRow{Row: current.Row + dir.Row, Boxes: []Box{}}

		// If all adjacent cells are empty spaces, move the boxes
		if isAdjacentCellsEmpty(warehouse, current.Boxes, dir) {
			for i := len(boxesToMove) - 1; i >= 0; i-- {
				boxR := boxesToMove[i]

				for _, box := range boxR.Boxes {
					warehouse[box.Pos.Row+dir.Row][box.Pos.Col+dir.Col] = box.Char
					warehouse[box.Pos.Row][box.Pos.Col] = SPACE
				}
			}

			// Move the robot
			warehouse[robotPos.Row][robotPos.Col] = SPACE
			robotPos = g.Point{Row: robotPos.Row + dir.Row, Col: robotPos.Col + dir.Col}
			warehouse[robotPos.Row][robotPos.Col] = ROBOT

			break
		}

		// Find all adjacent boxes in the direction we are moving
		for _, box := range current.Boxes {
			newR, newC := box.Pos.Row+dir.Row, box.Pos.Col+dir.Col

			// Check if we are blocked by a wall
			if warehouse[newR][newC] == WALL {
				return robotPos
			}

			// Add the boxes to the new box row
			if warehouse[newR][newC] == BOX_LEFT {
				newBoxRow.Boxes = append(newBoxRow.Boxes,
					Box{Pos: g.Point{Row: newR, Col: newC}, Char: BOX_LEFT},
					Box{Pos: g.Point{Row: newR, Col: newC + 1}, Char: BOX_RIGHT})
			} else if warehouse[newR][newC] == BOX_RIGHT {
				newBoxRow.Boxes = append(newBoxRow.Boxes,
					Box{Pos: g.Point{Row: newR, Col: newC - 1}, Char: BOX_LEFT},
					Box{Pos: g.Point{Row: newR, Col: newC}, Char: BOX_RIGHT})
			}
		}

		queue.Append(newBoxRow)
		boxesToMove = append(boxesToMove, newBoxRow)
	}

	return robotPos
}

// Check if the adjacent cells are empty spaces
func isAdjacentCellsEmpty(warehouse g.Grid, boxes []Box, dir g.Point) bool {
	for _, box := range boxes {
		if warehouse[box.Pos.Row+dir.Row][box.Pos.Col+dir.Col] != SPACE {
			return false
		}
	}

	return true
}

// Double the width of the warehouse
func doubleWarehouseWidth(warehouse g.Grid) g.Grid {
	resizedWarehouse := make(g.Grid, len(warehouse))
	for i := range resizedWarehouse {
		resizedWarehouse[i] = make([]rune, len(warehouse[0])*2)
	}

	for i := 0; i < len(resizedWarehouse); i++ {
		for j := 0; j < len(resizedWarehouse[0])-1; j += 2 {
			cell := warehouse[i][j/2]

			if cell == ROBOT {
				// Robot will still only occupy one cell
				resizedWarehouse[i][j] = ROBOT
				resizedWarehouse[i][j+1] = SPACE
			} else if cell == BOX {
				// Convert the box to a double width box ([] instead of O)
				resizedWarehouse[i][j] = BOX_LEFT
				resizedWarehouse[i][j+1] = BOX_RIGHT
			} else {
				// If the cell is not a robot or a box, just copy it
				resizedWarehouse[i][j] = cell
				resizedWarehouse[i][j+1] = cell
			}
		}
	}

	return resizedWarehouse
}

// Get the GPS coordinates of the boxes in the warehouse
// (Goods Positioning System)
func getBoxGPSCoordinates(warehouse g.Grid, doubleWidth bool) int {
	sum := 0

	for i, row := range warehouse {
		for j, cell := range row {
			if cell == BOX || (cell == BOX_LEFT && doubleWidth) {
				sum += (MULTIPLIER * i) + j
			}
		}
	}

	return sum
}

// Get the position of the robot in the warehouse
func getRobotPos(warehouse g.Grid) g.Point {
	for i, row := range warehouse {
		for j, cell := range row {
			if cell == ROBOT {
				return g.Point{Row: i, Col: j}
			}
		}
	}

	return g.Point{Row: -1, Col: -1}
}

// Convert the direction characters to numerical directions
func getDirections(moves []rune) []g.Point {
	dirs := make([]g.Point, 0)

	dirMap := map[rune]g.Point{
		'^': {Row: -1, Col: 0},
		'>': {Row: 0, Col: 1},
		'v': {Row: 1, Col: 0},
		'<': {Row: 0, Col: -1},
	}

	for _, move := range moves {
		dirs = append(dirs, dirMap[move])
	}

	return dirs
}
