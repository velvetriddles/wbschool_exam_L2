package pattern

import "fmt"

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

// Интерфейс состояния
type State interface {
	Handle()
}

// Конкретное состояние - Красный свет
type RedLight struct{}

func (r *RedLight) Handle() {
	fmt.Println("Red Light - Stop the car!")
}

// Конкретное состояние - Желтый свет
type YellowLight struct{}

func (y *YellowLight) Handle() {
	fmt.Println("Yellow Light - Get ready to move!")
}

// Конкретное состояние - Зелёный свет
type GreenLight struct{}

func (g *GreenLight) Handle() {
	fmt.Println("Green Light - Go!")
}

// Контекст - светофор
type TrafficLight struct {
	state State
}

// Метод для установки текущего состояния светофора
func (t *TrafficLight) SetState(state State) {
	t.state = state
}

// Метод для выполнения текущего состояния
func (t *TrafficLight) Handle() {
	t.state.Handle()
}

// func main() {
// 	// Создаем светофор
// 	trafficLight := &TrafficLight{}

// 	// Устанавливаем Красный свет
// 	red := &RedLight{}
// 	trafficLight.SetState(red)
// 	trafficLight.Handle() // Output: Red Light - Stop the car!

// 	// Устанавливаем Желтый свет
// 	yellow := &YellowLight{}
// 	trafficLight.SetState(yellow)
// 	trafficLight.Handle() // Output: Yellow Light - Get ready to move!

// 	// Устанавливаем Зелёный свет
// 	green := &GreenLight{}
// 	trafficLight.SetState(green)
// 	trafficLight.Handle() // Output: Green Light - Go!
// }
