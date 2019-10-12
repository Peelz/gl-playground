package problems

func twoSum(nums []int, target int) []int {
	table := make(map[int]int)
	for i, v := range nums {

		initial := target - v
		if val, ok := table[initial]; ok {
			return []int{val, i}
		}
		table[v] = i
	}
	return []int{}

}
