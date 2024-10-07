package main

import (
	"fmt"
	"math"
)

func NumberSum(number int) int {
	firstDigit := number / 1000
	secondDigit := number / 100 % 10
	thirdDigit := number / 10 % 10
	lastDigit := number % 10
	return firstDigit + thirdDigit + secondDigit + lastDigit
}

func CelciaToFar(number float64) float64 {
	return (number * (9. / 5.)) + 32.
}

func FarToCelcia(number float64) float64 {
	return (number - 32.) * (5. / 9.)
}

func DoubleValues(nums []int) []int {
	return []int{nums[0] * 2, nums[1] * 2, nums[2] * 2, nums[3] * 2}
}

func ConcatStr(str1 string, str2 string) string {
	newString := str1 + " " + str2
	return newString
}

func GetDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	return math.Sqrt(math.Pow(x2-x1, 2.) + math.Pow(y2-y1, 2.))
}

func IsLeapYear(year int) bool {
	return (year%400 == 0) || (year%4 == 0 && year%100 != 0)
}

func Factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	return n * Factorial(n-1)
}
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func ReverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}

func TaskOne() {
	var number int
	fmt.Println("Enter Number:")
	fmt.Scanln(&number)
	fmt.Println(NumberSum(number))
}
func TaskTwo() {
	var number float64
	fmt.Println("Введите градусы по цельсию:")
	fmt.Scanln(&number)
	fmt.Println("Градусов по фарегейту", CelciaToFar(number))
	fmt.Println("Градусов по цельсию", FarToCelcia(CelciaToFar(number)))
}

func TaskThree() {
	var input1, input2, input3, input4 int
	fmt.Printf("Введите первое число ->")
	fmt.Scan(&input1)
	fmt.Printf("Введите второе число->")
	fmt.Scan(&input2)
	fmt.Printf("Введите третье число->")
	fmt.Scan(&input3)
	fmt.Printf("Введите четвёртое число->")
	fmt.Scan(&input4)
	array := []int{input1, input2, input3, input4}
	fmt.Println(DoubleValues(array))
}

func TaskFour() {
	var input1, input2 string
	fmt.Printf("Введите первую строку->")
	fmt.Scan(&input1)
	fmt.Printf("Введите вторую строку->")
	fmt.Scan(&input2)

	fmt.Printf(ConcatStr(input1, input2))
}

func TaskFive() {
	var x1, x2, y1, y2 float64
	fmt.Printf("x1->")
	fmt.Scan(&x1)
	fmt.Printf("y1->")
	fmt.Scan(&y1)
	fmt.Printf("x2->")
	fmt.Scan(&x2)
	fmt.Printf("y2->")
	fmt.Scan(&y2)
	fmt.Print(GetDistance(x1, y1, x2, y2))
}

func TaskTwoOne() {
	var input int
	fmt.Printf("Введите число->")
	fmt.Scan(&input)
	if input%2 == 0 {
		fmt.Printf("Число чётное")
	} else {
		fmt.Printf("Число нечётное")
	}
}

func TaskTwoTwo() {
	var year int
	fmt.Print("Введите год: ")
	fmt.Scan(&year)

	if IsLeapYear(year) {
		fmt.Println("Високосный")
	} else {
		fmt.Println("Не високосный")
	}

}

func TaskTwoThree() {
	var a, b, c int
	fmt.Print("Введите три числа: ")
	fmt.Scan(&a, &b, &c)

	max := a

	if b > max {
		max = b
	}
	if c > max {
		max = c
	}

	fmt.Println("Наибольшее число:", max)
}

func TaskTwoFour() {
	var age int
	fmt.Print("Введите возраст: ")
	fmt.Scan(&age)

	if age >= 0 && age <= 12 {
		fmt.Println("Ребенок")
	} else if age >= 13 && age <= 17 {
		fmt.Println("Подросток")
	} else if age >= 18 && age <= 64 {
		fmt.Println("Взрослый")
	} else if age >= 65 {
		fmt.Println("Пожилой")
	} else {
		fmt.Println("Некорректный возраст")
	}
}

func TaskTwoFive() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	if number%3 == 0 && number%5 == 0 {
		fmt.Println("Делится")
	} else {
		fmt.Println("Не делится")
	}
}

func TaskThreeOne() {
	var number int
	fmt.Print("Введите число: ")
	fmt.Scan(&number)

	result := Factorial(number)
	fmt.Printf("Факториал числа %d равен %d\n", number, result)
}

func TaskThreeTwo() {
	var n int
	fmt.Print("Введите количество чисел Фибоначчи: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("Некорректное значение. Введите положительное число.")
		return
	}

	fmt.Print("Первые ", n, " чисел Фибоначчи: ")
	for i := 0; i < n; i++ {
		fmt.Print(Fibonacci(i), " ")
	}
	fmt.Println()
}

func TaskThreeThree() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println("Исходный массив:", arr)

	reversedArr := ReverseArray(arr)
	fmt.Println("Перевёрнутый массив:", reversedArr)
}

func TaskThreeFour() {
	var n int
	fmt.Print("Введите число: ")
	fmt.Scan(&n)

	fmt.Print("Простые числа до ", n, ": ")
	for i := 2; i <= n; i++ {
		if IsPrime(i) {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

func TaskThreeFive() {
	arr := []int{1, 2, 3, 4, 5}
	sum := 0

	for _, value := range arr {
		sum += value
	}

	fmt.Println("Сумма чисел в массиве:", sum)
}

func main() {
	TaskThreeFive()
}
