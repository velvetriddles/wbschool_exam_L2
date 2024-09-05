package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

/*
=== Взаимодействие с ОС ===

Необходимо реализовать собственный шелл

встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах

Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func main() {
	fmt.Println("Welcome to Go Shell!")

	// Перехватываем сигналы SIGINT (Ctrl+C) и SIGQUIT (Ctrl+\)
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGQUIT)

	for {
		// Выводим приглашение командной строки
		fmt.Print("> ")

		// Читаем ввод пользователя
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Ошибка чтения ввода:", err)
			continue
		}

		// Удаляем символ новой строки из ввода
		input = strings.TrimSpace(input)

		// Разбиваем ввод на аргументы
		args := strings.Fields(input)
		if len(args) == 0 {
			continue
		}

		// Проверяем встроенные команды
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Println("Использование: cd <directory>")
				continue
			}
			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Ошибка при переходе в директорию:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Ошибка при получении текущей директории:", err)
				continue
			}
			fmt.Println(dir)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("Использование: kill <pid>")
				continue
			}
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Println("Неверный формат PID:", err)
				continue
			}
			p, err := os.FindProcess(pid)
			if err != nil {
				fmt.Println("Ошибка при поиске процесса:", err)
				continue
			}
			err = p.Kill()
			if err != nil {
				fmt.Println("Ошибка при завершении процесса:", err)
			}
		case "ps":
			cmd := exec.Command("ps", "aux")
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Ошибка при выполнении ps:", err)
			}
		default:
			fmt.Println("Неизвестная команда:", args[0])
		}

		// Ожидаем сигналы SIGINT и SIGQUIT
		select {
		case sig := <-signals:
			fmt.Printf("Получен сигнал: %v\n", sig)
			os.Exit(0)
		default:
		}
	}
}
