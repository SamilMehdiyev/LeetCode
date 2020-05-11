package solutions

// FloodFill function is a solution for the
// Problem #733 - Flood Fill
// from leetcode.com
func FloodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	if image == nil || len(image) == 0 {
		return image
	}

	fillNewColor(image, sr, sc, newColor, image[sr][sc])

	return image
}

func fillNewColor(image [][]int, sr int, sc int, newColor int, originalColor int) {

	if sr < 0 || sr >= len(image) || sc < 0 || sc >= len(image[sr]) {
		return
	}

	if image[sr][sc] == newColor || image[sr][sc] != originalColor {
		return
	}

	image[sr][sc] = newColor

	fillNewColor(image, sr-1, sc, newColor, originalColor)
	fillNewColor(image, sr, sc+1, newColor, originalColor)
	fillNewColor(image, sr+1, sc, newColor, originalColor)
	fillNewColor(image, sr, sc-1, newColor, originalColor)
}
