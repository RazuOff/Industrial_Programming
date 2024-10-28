package main

import (
	"fmt"
	"strconv"
	"strings"
)

func triangleAngle(base float64, height float64) float64 {

	return base * .5 * height
}

func sortArray(arr []int) []int {
	if len(arr) < 2 {
		return arr
	}

	left, right := 0, len(arr)-1
	pivot := arr[len(arr)/2]

	for left <= right {
		for arr[left] < pivot {
			left++
		}
		for arr[right] > pivot {
			right--
		}
		if left <= right {
			arr[left], arr[right] = arr[right], arr[left]
			left++
			right--
		}
	}

	if right > 0 {
		sortArray(arr[:right+1])
	}
	if left < len(arr) {
		sortArray(arr[left:])
	}

	return arr
}

func sumOfSquares(n int) int {
	sum := 0
	for i := 2; i <= n; i += 2 {
		sum += i * i
	}
	return sum
}

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, " ", "")

	left, right := 0, len(s)-1

	for left < right {
		if s[left] != s[right] {
			return false
		}
		left++
		right--
	}

	return true
}

func isPrime(n int) bool {
	if n <= 1 {
		return false
	}
	// Проверяем делители от 2 до корня из n
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func generatePrimes(limit int) []int {
	var primes []int

	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			primes = append(primes, i)
		}
	}

	return primes
}

func toBinary(n int) string {
	return strconv.FormatInt(int64(n), 2)
}
func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}

	return max
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func sumArray(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

func main() {
	arr := []int{1, 5, 3, 9, 2, 10, 7}
	fmt.Println(findMax(arr))
}
