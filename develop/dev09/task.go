package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	// Проверяем, что передан URL-адрес для скачивания
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <URL>")
		return
	}

	url := os.Args[1] // URL-адрес для скачивания
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching URL: %v\n", err)
		return
	}
	defer resp.Body.Close()

	// Создаем файл для сохранения содержимого
	file, err := os.Create("downloaded_file.txt")
	if err != nil {
		fmt.Printf("Error creating file: %v\n", err)
		return
	}
	defer file.Close()

	// Копируем содержимое из ответа в файл
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Printf("Error copying content to file: %v\n", err)
		return
	}

	fmt.Println("File downloaded successfully!")
}
