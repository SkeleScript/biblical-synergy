package main

import "fmt"

func main() {
	fmt.Println("Hello, Colly")

	_, idx := BinSearch([]int{1, 2, 3, 4, 5, 6, 7, 8, 9}, 2)
	fmt.Println("idx is ", idx)
}

func BinSearch(nums []int, want int) (err error, idx int) {
	if (len(nums) == 0) {
		fmt.Println("empty")
		err = fmt.Errorf("empty")
		return
	}

	mid := len(nums) / 2
	if (nums[mid] == want) {
		idx = mid
		fmt.Println("idx is mid")
	}

	if nums[mid] > want {
		err, idx = BinSearch(nums[:mid], want)
		fmt.Println("idx is great mid")
	} else {
		err, idx = BinSearch(nums[mid+1:], want)
		fmt.Println("idx less mid")
		idx += mid + 1
	}

	fmt.Println("idx is ", idx)

	return
}
