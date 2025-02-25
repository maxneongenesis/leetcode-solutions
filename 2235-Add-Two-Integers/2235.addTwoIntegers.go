package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin) // Сканер стандартного ввода
	if scanner.Scan() {
		input := scanner.Text()            //Ввод строки
		parts := strings.Split(input, " ") // Разделяем элементы строки знаком "пробел"
		num1, _ := strconv.Atoi(parts[0])  // Присваиваем
		num2, _ := strconv.Atoi(parts[1])
		sum := num1 + num2
		fmt.Printf("%d\n", sum) // Вывод суммы
	}
}
