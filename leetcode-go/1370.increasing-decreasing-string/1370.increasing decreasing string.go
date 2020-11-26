package _1370_increasing_decreasing_string

import (
	"fmt"
	"reflect"
)

////自己的解法： 哈希表：效率低
//执行用时：
//12 ms
//, 在所有 Go 提交中击败了
//32.69%
//的用户
//内存消耗：
//5 MB
//, 在所有 Go 提交中击败了
//23.08%
//的用户
//func sortString(s string) string {
//	sortMap := make(map[byte]int, len(s))
//	for _, value := range s {
//		sortMap[byte(value)]++
//	}
//	arr := []byte{}
//	for len(arr) < len(s) {
//		for i := byte('a'); i < byte('z')+1; i++ {
//			if sortMap[i] > 0 {
//				arr = append(arr,i)
//				sortMap[i]--
//			}
//		}
//		for i := byte('z'); i > byte('a')-1; i-- {
//			if sortMap[i] > 0 {
//				arr = append(arr,i)
//				sortMap[i]--
//			}
//		}
//	}
//	return string(arr)
//}

////官方答案：桶计数
//执行用时：
//4 ms
//, 在所有 Go 提交中击败了
//79.81%
//的用户
//内存消耗：
//3 MB
//, 在所有 Go 提交中击败了
//91.35%
//的用户
func sortString(s string) string {
	//Go 语言的 string 是用 utf-8 进行编码的，英文字母占用一个字节8BIT ，而中文字母占用 3个字节
	fmt.Println("'a'---type:", reflect.TypeOf('a'))                                                                //97-------输出的是a的unicode编码,int32
	cnt := ['z' + 1]int{}                                                                                          //GO的字符有两种 byte(utf-8)、rune(int32)，'z'+1类型是int32。
	fmt.Println("'z'+1,1,byte('z')----type:", reflect.TypeOf('z'+1), reflect.TypeOf(1), reflect.TypeOf(byte('z'))) //int32 int uint8
	for _, ch := range s {
		fmt.Println("ch-type:", reflect.TypeOf(ch))
		cnt[ch]++
	}
	n := len(s)
	ans := make([]byte, 0, n)
	for len(ans) < n {
		for i := byte('a'); i <= 'z'; i++ {
			if cnt[i] > 0 { //TODO:为什么这里KEY定义的是INT32而I是BYTE仍然没有报错？
				ans = append(ans, i)
				cnt[i]--
			}
		}
		for i := byte('z'); i >= 'a'; i-- {
			if cnt[i] > 0 {
				ans = append(ans, i)
				cnt[i]--
			}
		}
	}
	return string(ans)
}

//作者：LeetCode-Solution
//链接：https://leetcode-cn.com/problems/increasing-decreasing-string/solution/shang-sheng-xia-jiang-zi-fu-chuan-by-leetcode-solu/
//来源：力扣（LeetCode）
//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
