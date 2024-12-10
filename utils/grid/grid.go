package grid

type Point struct {
	Row, Col int
}

type Grid [][]rune

func (g Grid) IsInsideBounds(r, c int) bool {
	return r >= 0 && r < len(g) && c >= 0 && c < len(g[0])
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
