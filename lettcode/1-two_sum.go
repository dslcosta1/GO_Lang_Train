func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)

	for index, num := range nums {
		value, ok := numsMap[target-num]

		if ok {
			return []int{value, index}
		} else {
			numsMap[num] = index
		}
	}

	return []int{-1, -1}
}