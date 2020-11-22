package _242_valid_anagram

//Hash
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap, tMap := make(map[string]int), make(map[string]int)
	for _, value := range s {
		sMap[string(value)]++
	}
	for _, value := range t {
		tMap[string(value)]++
	}

	i := 0
	for key, value := range sMap {
		if tMap[key] != value {
			i++
		}
	}

	if i != 0 {
		return false
	}
	return true
}

////精简版
//func isAnagram(s string, t string) bool {
//	if len(s)!=len(t){
//		return false
//	}
//
//	sMap:=make(map[string]int)
//	for _,value := range s{
//		sMap[string(value)]++
//	}
//	for _,value := range t{
//		if sMap[string(value)] >= 0 {
//			sMap[string(value)]--
//		}
//		if sMap[string(value)]<0{
//			return false
//		}
//	}
//	return true
//}

////再精简 去掉类型转换节约大量内存 提高速度
//执行用时：
//8 ms
//, 在所有 Go 提交中击败了
//59.89%
//的用户
//内存消耗：
//2.8 MB
//, 在所有 Go 提交中击败了
//66.27%
//的用户
//func isAnagram(s string, t string) bool {
//	if len(s)!=len(t){
//		return false
//	}
//
//	sMap:=make(map[rune]int)
//	for _,value := range s{
//		sMap[value]++
//	}
//	for _,value := range t{
//		if sMap[value] >= 0 {
//			sMap[value]--
//		}
//		if sMap[value]<0{
//			return false
//		}
//	}
//	return true
//}

////标准Hash
//func isAnagram(s, t string) bool {
//	var c1, c2 [26]int
//	for _, ch := range s {
//		c1[ch-'a']++
//	}
//	for _, ch := range t {
//		c2[ch-'a']++
//	}
//	return c1 == c2
//}
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/valid-anagram/solution/you-xiao-de-zi-mu-yi-wei-ci-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

//Sort
//func isAnagram(s, t string) bool {
//	s1, s2 := []byte(s), []byte(t)
//	sort.Slice(s1, func(i, j int) bool { return s1[i] < s1[j] })
//	sort.Slice(s2, func(i, j int) bool { return s2[i] < s2[j] })
//	return string(s1) == string(s2)
//}
//
//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/valid-anagram/solution/you-xiao-de-zi-mu-yi-wei-ci-by-leetcode-solution/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
