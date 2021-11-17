package duplicates

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	ans := false
	arr := make([]int, n)
	hshmap := make(map[int]int)
	for i := range arr {
		fmt.Scan(&arr[i])
		hshmap[arr[i]] += 1
		if hshmap[arr[i]] > 1 {
			ans = true
		}
	}
	fmt.Print("Does your array have duplicates? ", ans)
}
