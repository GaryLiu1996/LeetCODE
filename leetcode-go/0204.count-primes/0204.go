package _204_count_primes

// 暴力法：超出时间限制
//func countPrimes(n int) int {
//	if n<3{
//		return 0
//	}
//	account:=n-2
//	for i:=2;i<n;i++ {
//		for j:=2;j<i;j++ {
//			if i%j ==0{
//				account--
//				j=i
//			}
//		}
//	}
//	fmt.Println("acc",account)
//	return account
//}

// Sieve of Eratosthenes:厄拉多塞 筛法原理 也叫：埃氏筛
//	40 ms	7.1 MB
func countPrimes(n int) int {
	nMap := make(map[int]int, n)
	count := 0
	for i := 2; i < n; i++ {
		if nMap[i] == 0 {
			count++
			for j := i * i; j < n; j += i {
				//for j:=i+i;j<n;j+=i {
				nMap[j] = 1
			}
		}
	}
	return count
}

//线性筛：
//执行用时：
//40 ms
//, 在所有 Go 提交中击败了
//29.66%
//的用户
//内存消耗：
//7.1 MB
//, 在所有 Go 提交中击败了
//24.68%
//的用户
//func countPrimes(n int) int {
//	primes := []int{}
//	isPrime := make([]bool, n)
//	for i := range isPrime {
//		isPrime[i] = true
//	}
//	for i := 2; i < n; i++ {
//		if isPrime[i] {
//			primes = append(primes, i)
//		}
//		for _, p := range primes {
//			if i*p >= n {
//				break
//			}
//			isPrime[i*p] = false
//			if i%p == 0 {
//				break
//			}
//		}
//	}
//	return len(primes)
//}
