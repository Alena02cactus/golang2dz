package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Вставьте ваш код здесь
	letterCount := make(map[rune]int)
	totalLetters := 0

	for _, char := range strings.ToLower(text) {
		if char >= 'a' && char <= 'z' {
			letterCount[char]++
			totalLetters++
		}
	}

	fmt.Println("Результат:")
	for letter, count := range letterCount {
		frequency := float64(count) / float64(totalLetters) * 100
		fmt.Printf("%c - %d %.2f\n", letter, count, frequency)
	}
}
