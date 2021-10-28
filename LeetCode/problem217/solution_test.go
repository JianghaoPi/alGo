package problem217

import "testing"

func TestContainsDuplicate(t *testing.T)  {
	nums1 := []int{1, 2, 3, 1}
	if ans := containsDuplicate(nums1); ans != true {
		t.Errorf("check %v contains duplicate expected be true, but false got", nums1)
	}

	nums2 := []int{1, 2, 3, 4}
	if ans := containsDuplicate(nums2); ans != false {
		t.Errorf("check %v contains duplicate expected be false, but true got", nums2)
	}
}
