package grid

import "fmt"

type Point struct {
	Row, Col int
}

type Grid [][]rune

func (g Grid) Get(r, c int) rune {
	return g[r][c]
}

func (g Grid) Set(r, c int, v rune) {
	g[r][c] = v
}

func (g Grid) IsInsideBounds(r, c int) bool {
	return r >= 0 && r < len(g) && c >= 0 && c < len(g[0])
}

func (g Grid) Print() {
	for _, r := range g {
		for _, c := range r {
			fmt.Print(string(c))
		}
		fmt.Println()
	}
	fmt.Println()
}

func (g *Grid) Clear(v rune) {
	for i := 0; i < len(*g); i++ {
		for j := 0; j < len((*g)[i]); j++ {
			(*g)[i][j] = v
		}
	}
}

func Create(rows, cols int, v rune) Grid {
	grid := make(Grid, rows)

	for r := range grid {
		grid[r] = make([]rune, cols)

		for c := range grid[r] {
			grid[r][c] = v
		}
	}

	return grid
}

func CreateFromLines(lines []string) Grid {
	grid := make([][]rune, len(lines))

	for r, line := range lines {
		grid[r] = make([]rune, len(line))

		for c, v := range line {
			grid[r][c] = v
		}
	}

	return grid
}

func IsInsideBounds(grid [][]rune, r, c int) bool {
	return r >= 0 && r < len(grid) && c >= 0 && c < len(grid[0])
}
