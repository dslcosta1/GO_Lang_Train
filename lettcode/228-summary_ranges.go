import "strconv"

func summaryRanges(nums []int) []string {
	var result []string

	var limit = len(nums)
	if limit == 0 {
		return result
	} else if limit == 1 {
		result = append(result, strconv.Itoa(nums[0]))
		return result
	}

	var start int = nums[0]
	var end int = nums[0]

	for i := 1; i < limit; i++ {
		if end == nums[i]-1 {
			end = nums[i]
		} else {
			if start == nums[i-1] {
				result = append(result, strconv.Itoa(start))
			} else {
				result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
			}
			start = nums[i]
			end = nums[i]
		}
	}

	if start == nums[limit-1] {
		result = append(result, strconv.Itoa(start))
	} else {
		result = append(result, strconv.Itoa(start)+"->"+strconv.Itoa(end))
	}

	return result
}