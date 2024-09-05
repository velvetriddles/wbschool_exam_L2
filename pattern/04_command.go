package pattern

import "fmt"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

type RemoteControl struct {
	command func()
}

func (r *RemoteControl) SetCommand(c func()) {
	r.command = c
}

func (r *RemoteControl) PressButton() {
	r.command()
}

// Получатель команды
type TV struct {
	IsOn bool
}

func (tv *TV) On() {
	tv.IsOn = true
	fmt.Println("TV is now ON")
}

func (tv *TV) Off() {
	tv.IsOn = false
	fmt.Println("TV is now OFF")
}

// func main() {
// 	// Создаем телевизор
// 	tv := &TV{}

// 	// Создаем пульт управления
// 	remote := &RemoteControl{}

// 	// Задаем команду включения телевизора через функцию
// 	remote.SetCommand(func() {
// 		tv.On()
// 	})
// 	remote.PressButton() // TV is now ON

// 	// Задаем команду выключения телевизора через функцию
// 	remote.SetCommand(func() {
// 		tv.Off()
// 	})
// 	remote.PressButton() // TV is now OFF
// }
