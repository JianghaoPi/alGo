package problem217

// Problem: https://leetcode.com/problems/contains-duplicate/
// Use a map to record if a num has shown before

func containsDuplicate(nums []int) bool {
	if len(nums) <= 1 {
		return false
	}

	record := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := record[nums[i]]; ok {
			return true
		}
		record[nums[i]] = true
	}
	return false
}
