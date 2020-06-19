package main

func main() {
	// 这题用手机做的，没保留测试用例
}

//----------------------------------------下面才是解题代码----------------------------------------

func sortColors(nums []int) {
	l := 0
	r := len(nums) - 1
	for i := 0; i <= r && l < r; i++ {
		if nums[i] == 0 {
			swap(nums, l, i)
			l++
		} else if nums[i] == 2 {
			swap(nums, i, r)
			r--
			i--
		}
	}
}

func swap(nums []int, i, j int) {
	if i == j {
		return
	}
	nums[i] ^= nums[j]
	nums[j] ^= nums[i]
	nums[i] ^= nums[j]
}
