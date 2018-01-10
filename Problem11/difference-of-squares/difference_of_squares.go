//Package diffsquares
// Find the difference between the square of the sum and the sum of the squares of the first N natural numbers.
package diffsquares

// SquareOfSums is the square of the sum of the first 10 natural numbers is
// (1 + 2 + ... + 10)² = 55² = 3025.
func SquareOfSums(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		total += i
	}

	return total * total
}

// SumOfSquares is the sum of the squares of the first ten natural numbers is
// 1² + 2² + ... + 10² = 385.
func SumOfSquares(n int) int {
	total := 0

	for i := 1; i <= n; i++ {
		total += i * i
	}

	return total
}

// Difference is the difference between the square of the sum of the first
// ten natural numbers and the sum of the squares of the first ten
// natural numbers is 3025 - 385 = 2640.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
