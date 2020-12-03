package _321_create_maximum_number

//
//func maxNumber(nums1 []int, nums2 []int, k int) []int {
//	arr1,arr2 := []int{},[]int{}
//	if len(nums1)<=k{
//		arr1 = nums1
//	}else{
//		for i,drop :=0,0;len(nums1)-drop > k&&i+1<len(nums1); {
//			if nums1[i] <nums1[i+1]{
//				arr1 =append(arr1,nums1[i])
//			}
//		}
//	}
//	if len(nums2)<=k{
//		arr2 = nums2
//	}
//
//}

//官方解法：
func maxSubsequence(a []int, k int) (s []int) {
	for i, v := range a {
		for len(s) > 0 && len(s)+len(a)-1-i >= k && v > s[len(s)-1] {
			s = s[:len(s)-1]
		}
		if len(s) < k {
			s = append(s, v)
		}
	}
	return
}

func lexicographicalLess(a, b []int) bool {
	for i := 0; i < len(a) && i < len(b); i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}
	return len(a) < len(b)
}

func merge(a, b []int) []int {
	merged := make([]int, len(a)+len(b))
	for i := range merged {
		if lexicographicalLess(a, b) {
			merged[i], b = b[0], b[1:]
		} else {
			merged[i], a = a[0], a[1:]
		}
	}
	return merged
}

func maxNumber(nums1, nums2 []int, k int) (res []int) {
	start := 0
	if k > len(nums2) {
		start = k - len(nums2)
	}
	for i := start; i <= k && i <= len(nums1); i++ {
		s1 := maxSubsequence(nums1, i)
		s2 := maxSubsequence(nums2, k-i)
		merged := merge(s1, s2)
		if lexicographicalLess(res, merged) {
			res = merged
		}
	}
	return
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/create-maximum-number/solution/pin-jie-zui-da-shu-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
