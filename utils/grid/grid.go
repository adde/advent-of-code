package grid

func CreateFromLines(lines []string) [][]rune {
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
