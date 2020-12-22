package diffsquares

func SumOfSquares( n int) int {
	sumsq := 0
	for i:=1; i<=n; i++ {
		sumsq += i*i
	}
	return sumsq
}

func SquareOfSum( n int) int {
	sum := 0
	for i:=1; i<=n; i++ {
		sum += i
	}
	return sum*sum
}

func Difference( n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}