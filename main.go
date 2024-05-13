package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите предложение")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Вставьте ваш код здесь
}
