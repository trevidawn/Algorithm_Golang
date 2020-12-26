package medium

type SubrectangleQueries struct {
	rect [][]int
}

func Constructor(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rect: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int)  {

	var minX, minY, maxX, maxY int
	if col1 < col2 {
		minX = col1
		maxX = col2
	} else {
		minX = col1
		maxX = col2
	}

	if row1 < row2 {
		minY = row1
		maxY = row2
	} else {
		minY = row2
		maxY = row1
	}

	for y := minY; y <= maxY; y++ {
		for x := minX; x <= maxX; x++ {
			this.rect[y][x] = newValue
		}
	}
}


func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rect[row][col]
}

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * obj := Constructor(rectangle);
 * obj.UpdateSubrectangle(row1,col1,row2,col2,newValue);
 * param_2 := obj.GetValue(row,col);
 */