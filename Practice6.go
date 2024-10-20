package main

import (
	"fmt"
	"math"
	"strings"

	"math/rand"
	"time"
)

func BubbleSort(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
			fmt.Printf("Шаг %d.%d: %v\n", i, j, arr)
		}
		if !swapped {
			break
		}
	}
}

func TaskThree() {
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	fmt.Printf("Исходный массив: %v\n", numbers)
	BubbleSort(numbers)
	fmt.Printf("Отсортированный массив: %v\n", numbers)
}

func TaskFour() {
	for i := 1; i <= 10; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%4d", i*j)
		}
		fmt.Println()
	}
}

func PascalTriangle(levels int) {
	triangle := make([][]int, levels)

	for i := 0; i < levels; i++ {
		triangle[i] = make([]int, i+1)
		triangle[i][0], triangle[i][i] = 1, 1

		for j := 1; j < i; j++ {
			triangle[i][j] = triangle[i-1][j-1] + triangle[i-1][j]
		}

		fmt.Println(triangle[i])
	}
}

func TaskSeven() {
	var levels int
	fmt.Print("Введите количество уровней треугольника Паскаля: ")
	fmt.Scan(&levels)

	PascalTriangle(levels)
}

func FindMinMax(arr []int) (int, int) {
	min := math.MaxInt64
	max := math.MinInt64

	for _, num := range arr {
		if num > max {
			max = num
		}

		if num < min {
			min = num
		}
	}

	return max, min
}

func TaskNine() {
	numbers := []int{15, 3, 9, -2, 7, 19, 1, 8, -10}
	max, min := FindMinMax(numbers)
	fmt.Printf("Максимум: %d, Минимум: %d\n", max, min)
}

func GuessNumber() {

	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(100) + 1

	var guess, attempts int
	maxAttempts := 7

	fmt.Println("Я загадал число от 1 до 100. Угадайте его!")

	for attempts = 0; attempts < maxAttempts; attempts++ {
		fmt.Printf("Попытка %d: Введите ваше число: ", attempts+1)
		fmt.Scan(&guess)

		if guess > secretNumber {
			fmt.Println("Загаданное число меньше.")
		} else if guess < secretNumber {
			fmt.Println("Загаданное число больше.")
		} else {
			fmt.Printf("Поздравляю! Вы угадали число %d за %d попыток.\n", secretNumber, attempts+1)
			return
		}
	}

	fmt.Printf("Увы! Вы исчерпали количество попыток. Загаданное число было: %d\n", secretNumber)
}

func TaskTen() {
	GuessNumber()
}

func TaskTwelve() {
	fmt.Println("Введите строку:")
	var input string
	fmt.Scanln(&input)

	words := strings.Fields(strings.ToLower(input))

	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	fmt.Printf("Количество уникальных слов: %d\n", len(wordCount))

	for word, count := range wordCount {
		fmt.Printf("Слово '%s' встречается %d раз(а)\n", word, count)
	}
}

const (
	width  = 10 // ширина игрового поля
	height = 10 // высота игрового поля
)

// Функция для создания пустого игрового поля
func createBoard() [][]bool {
	board := make([][]bool, height)
	for i := range board {
		board[i] = make([]bool, width)
	}
	return board
}

// Функция для инициализации игрового поля случайными значениями
func randomizeBoard(board [][]bool) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			board[i][j] = rand.Intn(2) == 1
		}
	}
}

// Функция для вывода игрового поля
func printBoard(board [][]bool) {
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			if board[i][j] {
				fmt.Print("O ") // живая клетка
			} else {
				fmt.Print(". ") // мёртвая клетка
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

// Функция для подсчета живых соседей клетки
func countAliveNeighbors(board [][]bool, x, y int) int {
	count := 0
	// Проверяем все восемь соседей клетки
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if i == 0 && j == 0 {
				continue // не считаем саму клетку
			}
			// Используем цикличность для выхода за границы
			neighborX := (x + i + width) % width
			neighborY := (y + j + height) % height
			if board[neighborY][neighborX] {
				count++
			}
		}
	}
	return count
}

// Функция для обновления игрового поля по правилам игры
func nextGeneration(board [][]bool) [][]bool {
	newBoard := createBoard()

	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			aliveNeighbors := countAliveNeighbors(board, j, i)

			// Правила игры
			if board[i][j] && (aliveNeighbors == 2 || aliveNeighbors == 3) {
				newBoard[i][j] = true // клетка остаётся живой
			} else if !board[i][j] && aliveNeighbors == 3 {
				newBoard[i][j] = true // мёртвая клетка оживает
			} else {
				newBoard[i][j] = false // клетка умирает или остаётся мёртвой
			}
		}
	}
	return newBoard
}

func TaskThirteen() {
	board := createBoard()
	randomizeBoard(board)

	generations := 10

	for gen := 0; gen < generations; gen++ {
		fmt.Printf("Поколение %d:\n", gen+1)
		printBoard(board)
		board = nextGeneration(board)
		time.Sleep(1 * time.Second)
	}
}

func sumOfDigits(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

func digitalRoot(n int) int {
	for n >= 10 {
		n = sumOfDigits(n)
	}
	return n
}

func TaskFourteen() {
	fmt.Println("Введите число:")
	var number int
	fmt.Scan(&number)

	root := digitalRoot(number)

	fmt.Printf("Цифровой корень числа %d: %d\n", number, root)
}
func intToRoman(num int) string {
	vals := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""

	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			roman += syms[i]
			num -= vals[i]
		}
	}

	return roman
}

func TaskFiveteen() {
	fmt.Println("Введите арабское число:")
	var number int
	fmt.Scan(&number)

	roman := intToRoman(number)

	fmt.Printf("Римское представление числа %d: %s\n", number, roman)
}

func main() {
	TaskFour()
}
