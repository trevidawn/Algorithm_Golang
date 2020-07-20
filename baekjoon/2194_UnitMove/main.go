package main

import (
	"bufio"
	"fmt"
	"os"
)

type pair struct {
	x int
	y int

	depth int
}

var (
	N int
	M int
	A int
	B int
	K int

	dx = [4]int{-1, 0, 1, 0}
	dy = [4]int{0, 1, 0, -1}

	start pair
	end   pair

	q      [30000]pair
	qHead  int
	qTail  int
	qEmpty int
	Map    [501][501]int
	NewMap [501][501]int
	visit  [501][501]bool
)

func main() {

	r := bufio.NewReader(os.Stdin)

	fmt.Fscanf(r, "%d %d %d %d %d\n", &N, &M, &A, &B, &K)

	var n, m int
	for i := 0; i < K; i++ {
		fmt.Fscanf(r, "%d %d\n", &n, &m)
		Map[n][m] = 1
	}

	fmt.Fscanf(r, "%d %d\n", &start.x, &start.y)
	fmt.Fscanf(r, "%d %d\n", &end.x, &end.y)

	for i := 1; i <= N; i++ {
		for j := 1; j <= M; j++ {
			if Map[i][j] == 1 {
				for ii := i - A + 1; ii <= i; ii++ {
					for jj := j - B + 1; jj <= j; jj++ {
						if ii < 1 || ii > N || jj < 1 || jj > M {
							continue
						}

						NewMap[ii][jj] = 1
					}
				}
			}
		}
	}

	N = N - A + 1
	M = M - B + 1

	//for i := 1; i <= N; i++ {
	//	for j := 1; j <= M; j++ {
	//		fmt.Printf("%d ", NewMap[i][j])
	//	}
	//	fmt.Printf("\n")
	//}

	var ret int
	if start.x == end.x && start.y == end.y {
		ret = 0
	} else {
		ret = bfs()
	}

	fmt.Println(ret)
}

func bfs() int {
	q[qTail] = start
	qTail++
	qEmpty++
	visit[start.x][start.y] = true

	for qEmpty > 0 {
		front := q[qHead]
		qEmpty--
		qHead++

		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				nx := front.x + dx[i]
				ny := front.y + dy[i]
				if nx < 1 || nx > N || ny < 1 || ny > M || visit[nx][ny] {
					continue
				}
				if nx == end.x && ny == end.y {
					return front.depth + 1
				}

				if NewMap[nx][ny] == 0 {
					q[qTail] = pair{nx, ny, front.depth + 1}
					qTail++
					qEmpty++
					visit[nx][ny] = true
				}
			}
		}
	}

	return -1
}
