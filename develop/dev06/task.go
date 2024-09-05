package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

var (
	fieldsFlag    string
	delimiterFlag string
	separatedFlag bool
)

func init() {
	flag.StringVar(&fieldsFlag, "f", "", "выбрать поля (колонки)")
	flag.StringVar(&delimiterFlag, "d", "\t", "использовать другой разделитель")
	flag.BoolVar(&separatedFlag, "s", false, "только строки с разделителем")
	flag.Parse()
}

func main() {
	// Получение запрашиваемых полей
	fields := strings.Split(fieldsFlag, ",")
	for i, field := range fields {
		fields[i] = strings.TrimSpace(field)
	}

	// Создание сканера для чтения STDIN
	scanner := bufio.NewScanner(os.Stdin)

	// Обработка каждой строки STDIN
	for scanner.Scan() {
		line := scanner.Text()

		// Разбиение строки на колонки по разделителю
		columns := strings.Split(line, delimiterFlag)

		// Проверка, содержит ли строка разделитель
		if separatedFlag && len(columns) == 1 {
			continue
		}

		// Формирование и вывод строки с выбранными полями
		var result []string
		for _, field := range fields {
			index := atoi(field)
			if index > 0 && index <= len(columns) {
				result = append(result, columns[index-1])
			}
		}
		fmt.Println(strings.Join(result, delimiterFlag))
	}
}

// atoi преобразует строку в целое число.
// Возвращает 0, если преобразование не удалось.
func atoi(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
