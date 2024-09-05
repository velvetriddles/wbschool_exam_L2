Что выведет программа? Объяснить вывод программы.

```go
package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}
```

Ответ:
```
"error"
error - это интерфейс. Функция test возвращает указатель на тип данных customError, переменной err типа error присваивается интерфейс, который хранит о себе информацию о типе данных, поэтому переменная err не равняется nil.


```
