package solutions

// SubrectangleQueries type is a solution for the
// Question #2 - Make Two Arrays Equal by Reversing SubArrays
// from leetcode.com
type SubrectangleQueries struct {
	rectangle [][]int
}

// Constructor for the type SubrectangleQueries
func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle}
}

// UpdateSubrectangle function
func (obj *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			obj.rectangle[i][j] = newValue
		}
	}
}

// GetValue function
func (obj *SubrectangleQueries) GetValue(row int, col int) int {
	return obj.rectangle[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */
