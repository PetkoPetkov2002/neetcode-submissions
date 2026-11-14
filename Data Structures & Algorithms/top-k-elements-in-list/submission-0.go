func topKFrequent(nums []int, k int) []int {
	seenValues := make(map[int]int)
	bucket := make([][]int, len(nums)+1)
	var result []int
	for _, number := range nums{
		seenValues[number]++
	}
	for number, value := range seenValues{
		bucket[value] = append(bucket[value], number)
	}

	for i := len(bucket) - 1; i >= 0; i-- {
		for _, number := range bucket[i] {
			result = append(result, number)
			if len(result) == k {
				return result
			}
		}
	}

	return result
}
