package calc

func Sum(nums []int) int {
	var total int
	for _, num := range nums {
		total += num
	}
	return total
}
