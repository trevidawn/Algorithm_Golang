package main

import "fmt"

var nums [13]int
var selected [13]int
var K int

func main() {

	for {
		fmt.Scanf("%d", &K)
		if K == 0 {
			break
		}

		var num int
		for i := 0; i < K; i++ {
			fmt.Scanf("%d", &num)
			nums[i] = num
		}

		select_num(0, 0)
		fmt.Printf("\n")
	}
}

func select_num(cnt int, idx int) {
	if idx ==  K && cnt != 6 {
		return
	}
	if cnt == 6 {
		for i := 0; i < 6; i++ {
			fmt.Printf("%d ", selected[i])
		}
		fmt.Printf("\n")
		return
	}

	selected[cnt] = nums[idx]
	select_num(cnt+1,idx+1)
	select_num(cnt, idx+1)
}