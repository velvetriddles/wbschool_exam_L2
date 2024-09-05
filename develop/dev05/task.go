package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	after      = flag.Int("A", 0, "печатать +N строк после совпадения")
	before     = flag.Int("B", 0, "печатать +N строк до совпадения")
	context    = flag.Int("C", 0, "печатать ±N строк вокруг совпадения")
	count      = flag.Bool("c", false, "количество строк")
	ignoreCase = flag.Bool("i", false, "игнорировать регистр")
	invert     = flag.Bool("v", false, "вместо совпадения, исключать")
	fixed      = flag.Bool("F", false, "точное совпадение со строкой, не паттерн")
	lineNum    = flag.Bool("n", false, "печатать номер строки")
)

func main() {
	flag.Parse()

	// Получение паттерна поиска
	pattern := flag.Arg(0)

	// Проверка на наличие паттерна
	if pattern == "" {
		fmt.Println("Необходимо указать паттерн для поиска")
		os.Exit(1)
	}

	// Компиляция регулярного выражения
	var regex *regexp.Regexp
	if *fixed {
		pattern = regexp.QuoteMeta(pattern)
	}
	if *ignoreCase {
		regex = regexp.MustCompile(`(?i)` + pattern)
	} else {
		regex = regexp.MustCompile(pattern)
	}

	// Открытие файла для чтения
	file, err := os.Open(flag.Arg(1))
	if err != nil {
		fmt.Println("Ошибка при открытии файла:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Создание сканера для чтения файла построчно
	scanner := bufio.NewScanner(file)

	// Переменные для хранения контекста
	var contextLines []string
	contextCount := 0

	// Переменная для подсчета количества строк
	var totalCount int

	// Функция для печати строки с учетом флагов
	printLine := func(line string) {
		totalCount++
		if *count {
			return
		}
		if *lineNum {
			fmt.Printf("%d:", totalCount)
		}
		fmt.Println(line)
	}

	// Функция для печати строк контекста
	printContext := func() {
		for _, line := range contextLines {
			printLine(line)
		}
		contextLines = nil
		contextCount = 0
	}

	// Переменные для отслеживания совпадений
	var found bool
	var matchedLines []string

	// Переменная для подсчета строк до совпадения (флаг -B)
	beforeCount := *before

	// Просмотр каждой строки в файле
	for scanner.Scan() {
		line := scanner.Text()

		// Поиск совпадения
		matched := regex.MatchString(line)

		// Обработка совпадений
		if matched && !*invert {
			// Обработка флага -B (строки до совпадения)
			if *before > 0 {
				if contextCount < *before {
					contextLines = append(contextLines, line)
					contextCount++
					continue
				} else {
					printContext()
				}
			}
			// Печать строки совпадения
			printLine(line)
			// Печать строк после совпадения (флаг -A)
			for i := 0; i < *after; i++ {
				if scanner.Scan() {
					printLine(scanner.Text())
				}
			}
			// Подсчет строк после совпадения (флаг -C)
			contextCount = 0
			contextLines = nil
			found = true
			matchedLines = nil
		} else if !matched && *invert {
			// Если флаг -v (инвертирование) и совпадение не найдено
			printLine(line)
		} else {
			// Обработка флага -B (строки до совпадения)
			if *before > 0 && beforeCount > 0 {
				matchedLines = append(matchedLines, line)
				beforeCount--
				continue
			} else if beforeCount == 0 {
				// Печать строк до совпадения
				for _, l := range matchedLines {
					printLine(l)
				}
				beforeCount--
			}
			// Обработка флага -C (контекст)
			if *context > 0 && contextCount < *context {
				contextLines = append(contextLines, line)
				contextCount++
				continue
			}
			// Печать строк контекста (флаг -C)
			if *context > 0 && found {
				printContext()
			}
		}
	}

	// Печать строк контекста, если флаг -C не был указан
	if *context > 0 && found {
		printContext()
	}

	// Печать количества строк, если указан флаг -c
	if *count {
		fmt.Println("Количество строк:", totalCount)
	}
}
