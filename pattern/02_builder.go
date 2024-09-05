package pattern

/*
	Реализовать паттерн «строитель».

Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.

	https://en.wikipedia.org/wiki/Builder_pattern
*/

// Структура House, которую мы будем создавать через Builder
type House struct {
	windows int
	doors   int
	garage  bool
	pool    bool
}

// Строитель для объекта House
type HouseBuilder struct {
	house House
}

// Конструктор для строителя
func NewHouseBuilder() *HouseBuilder {
	return &HouseBuilder{House{}}
}

// Методы строителя для задания параметров

func (b *HouseBuilder) SetWindows(windows int) *HouseBuilder {
	b.house.windows = windows
	return b
}

func (b *HouseBuilder) SetDoors(doors int) *HouseBuilder {
	b.house.doors = doors
	return b
}

func (b *HouseBuilder) SetGarage(garage bool) *HouseBuilder {
	b.house.garage = garage
	return b
}

func (b *HouseBuilder) SetPool(pool bool) *HouseBuilder {
	b.house.pool = pool
	return b
}

// Метод для построения и возврата объекта House
func (b *HouseBuilder) Build() House {
	return b.house
}

// func main() {
// 	// Создаем объект с помощью строителя
// 	builder := NewHouseBuilder()
// 	house := builder.SetWindows(10).
// 		SetDoors(5).
// 		SetGarage(true).
// 		SetPool(true).
// 		Build()

// 	fmt.Printf("House: %+v\n", house)
// }
