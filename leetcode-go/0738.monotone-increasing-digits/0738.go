package _738_monotone_increasing_digits

func monotoneIncreasingDigits(N int) int {
	low := 9
	for i := 1; i <= N; i *= 10 {
		value := (N / i) % 10
		if value > low {
			N = N/i*i - 1
			value = (N / i) % 10
		}
		low = value
	}
	return N
}
