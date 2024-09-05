package mysort

import (
	"flag"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type CMDParams struct {
	NColumn int
	Numeric bool
	Reverse bool
	Unique  bool
}

func ParseArgs() CMDParams {
	nColumn := flag.Int("k", 0, "Указание колонки для сортировки")
	numeric := flag.Bool("n", false, "Сортировать по числовому значению")
	reverse := flag.Bool("r", false, "Сортировать в обратном порядке")
	unique := flag.Bool("u", false, "Не выводить повторяющиеся строки")

	flag.Parse()

	return CMDParams{
		NColumn: *nColumn,
		Numeric: *numeric,
		Reverse: *reverse,
		Unique:  *unique,
	}
}


// Функция для удаления повторяющихся строк
func removeDuplicates(lines []string) []string {
	uniqueLines := make(map[string]bool)
	var result []string
	for _, line := range lines {
		if !uniqueLines[line] {
			uniqueLines[line] = true
			result = append(result, line)
		}
	}
	return result
}

func Run(params CMDParams, lines []string) {
	comparator := func(i, j int) bool {
		line1 := strings.Fields(lines[i])
		line2 := strings.Fields(lines[j])

		// Обработка ключа -k
		k := params.NColumn - 1
		if k >= len(line1) || k >= len(line2) {
			return false
		}

		// Преобразование к числовому значению, если указан ключ -n
		if params.Numeric {
			num1, err1 := strconv.Atoi(line1[k])
			num2, err2 := strconv.Atoi(line2[k])
			if err1 == nil && err2 == nil {
				if num1 != num2 {
					return num1 < num2
				}
			} else {
				// Если не удалось преобразовать к числу, то сравниваем как строки
				return lines[i] < lines[j]
			}
		}

		// Сравнение строк
		if lines[i] != lines[j] {
			if params.Reverse {
				return lines[i] > lines[j] // Инвертированный порядок
			}
			return lines[i] < lines[j]
		}
		return false
	}

	// Сортировка строк
	sort.SliceStable(lines, comparator)

	// Удаление повторяющихся строк, если указан флаг -u
	if params.Unique {
		lines = removeDuplicates(lines)
	}

	// Вывод отсортированных строк
	for _, line := range lines {
		fmt.Println(line)
	}
}