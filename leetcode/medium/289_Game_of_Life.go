package medium

/*
	Runtime: 0 ms, faster than 100.00% of Go online submissions for Game of Life.
	Memory Usage: 2.1 MB, less than 5.44% of Go online submissions for Game of Life.
*/
func gameOfLife(board [][]int)  {

	var newBoard [][]int
	newBoard = make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		newBoard[i] = make([]int, len(board[i]))
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			newBoard[i][j] = CheckLiveOrDead(board, board[i][j], i, j)
		}
	}

	for i := 0; i < len(newBoard); i++ {
		board[i] = newBoard[i]
	}

}

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

var lifeMap = [][]int{
	{0, 0, 0, 1, 0, 0, 0, 0, 0},
	{0, 0, 1, 1, 0, 0, 0, 0, 0},
}

func CheckLiveOrDead(board [][]int, state, i,j int) int {

	var cnt = 0
	for d := 0; d < 8; d++ {
		nx := i + dx[d]
		ny := j + dy[d]

		if nx < 0 || ny < 0 || nx >= len(board) || ny >= len(board[0]) {
			continue
		}

		if board[nx][ny] == 1 {
			cnt++
		}
	}

	return lifeMap[state][cnt]
}