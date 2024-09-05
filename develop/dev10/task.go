package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"time"
)

/*
=== Утилита telnet ===

Реализовать примитивный telnet клиент:
Примеры вызовов:
go-telnet --timeout=10s host port go-telnet mysite.ru 8080 go-telnet --timeout=3s 1.1.1.1 123

Программа должна подключаться к указанному хосту (ip или доменное имя) и порту по протоколу TCP.
После подключения STDIN программы должен записываться в сокет, а данные полученные и сокета должны выводиться в STDOUT
Опционально в программу можно передать таймаут на подключение к серверу (через аргумент --timeout, по умолчанию 10s).

При нажатии Ctrl+D программа должна закрывать сокет и завершаться. Если сокет закрывается со стороны сервера, программа должна также завершаться.
При подключении к несуществующему сервер, программа должна завершаться через timeout.
*/

func main() {
	// Определяем флаги командной строки
	timeout := flag.Duration("timeout", 10*time.Second, "connection timeout")
	flag.Parse()

	// Проверяем, что переданы хост и порт
	if flag.NArg() < 2 {
		fmt.Println("Usage: go-telnet [--timeout=10s] host port")
		return
	}

	host := flag.Arg(0)
	port := flag.Arg(1)

	// Устанавливаем соединение с сервером
	conn, err := net.DialTimeout("tcp", host+":"+port, *timeout)
	if err != nil {
		fmt.Printf("Error connecting to %s:%s: %v\n", host, port, err)
		return
	}
	defer conn.Close()

	// Определяем канал для обработки сигнала Ctrl+C или Ctrl+D
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Канал для обработки завершения работы программы
	done := make(chan struct{})

	// Запускаем горутину для чтения данных из соединения и вывода их в STDOUT
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := conn.Read(buffer)
			if err != nil {
				fmt.Println("Connection closed by server.")
				close(done)
				return
			}
			fmt.Print(string(buffer[:n]))
		}
	}()

	// Читаем данные из STDIN и записываем их в соединение
	go func() {
		buffer := make([]byte, 1024)
		for {
			n, err := os.Stdin.Read(buffer)
			if err != nil {
				close(done)
				return
			}
			_, err = conn.Write(buffer[:n])
			if err != nil {
				fmt.Printf("Error writing to connection: %v\n", err)
				close(done)
				return
			}
		}
	}()

	// Ожидаем сигнала об окончании работы программы
	select {
	case <-interrupt:
	case <-done:
	}
}
