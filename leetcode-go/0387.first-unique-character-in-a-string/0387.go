package _387_first_unique_character_in_a_string

func firstUniqChar(s string) int {
	a := make(map[byte]int, len(s))
	for i := 0; i < len(s); i++ {
		a[s[i]]++
	}
	for i := 0; i < len(s); i++ {
		if a[s[i]] == 1 {
			return i
		}
	}
	return -1
}
