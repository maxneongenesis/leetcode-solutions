package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Создаем новый сканер для чтения из стандартного ввода
	scanner := bufio.NewScanner(os.Stdin)

	// Запрашиваем у пользователя ввод чисел через пробел
	fmt.Println("Введите числа через пробел:")

	// Считываем строку
	if !scanner.Scan() {
		fmt.Println("Ошибка: не удалось прочитать ввод.")
		return
	}

	// Получаем введённую строку
	input := scanner.Text()

	// Разделяем строку на отдельные элементы
	parts := strings.Fields(input)

	// Создаем срез для хранения чисел
	var arr []int

	// Преобразуем строки в числа
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			fmt.Printf("Ошибка: '%s' не является числом.\n", part)
			return
		}
		arr = append(arr, num)
	}

	result := runningSum(arr)
	writer := bufio.NewWriter(os.Stdout)
	for _, num := range result {
		writer.WriteString(strconv.Itoa(num) + " ")
	}
	writer.WriteString("\n")
	writer.Flush()
}

func runningSum(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		arr[i] += arr[i-1]
	}
	return arr
}
