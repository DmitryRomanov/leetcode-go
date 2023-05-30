package flood_fill

func floodFill(image [][]int, sr int, sc int, color int) [][]int {
	if image[sr][sc] != color {
		dfs(image, sr, sc, image[sr][sc], color)
	}
	return image
}

func dfs(image [][]int, row int, col int, source_color int, color int) {
	if row < 0 || col < 0 || row >= len(image) || col >= len(image[0]) || source_color != image[row][col] {
		return
	}

	image[row][col] = color

	dfs(image, row+1, col, source_color, color)
	dfs(image, row-1, col, source_color, color)
	dfs(image, row, col+1, source_color, color)
	dfs(image, row, col-1, source_color, color)
}
