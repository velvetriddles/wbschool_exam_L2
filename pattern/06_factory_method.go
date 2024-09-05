package pattern

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

// Интерфейс для транспорта
type Transport interface {
	Drive() string
}

// Реализация автомобиля
type Car struct{}

func (c *Car) Drive() string {
	return "Driving a car"
}

// Реализация мотоцикла
type Motorcycle struct{}

func (m *Motorcycle) Drive() string {
	return "Riding a motorcycle"
}

// Интерфейс фабрики для транспорта
type TransportFactory interface {
	CreateTransport() Transport
}

// Фабрика для создания автомобилей
type CarFactory struct{}

func (cf *CarFactory) CreateTransport() Transport {
	return &Car{}
}

// Фабрика для создания мотоциклов
type MotorcycleFactory struct{}

func (mf *MotorcycleFactory) CreateTransport() Transport {
	return &Motorcycle{}
}

// func main() {
// 	// Создаем фабрику для автомобилей
// 	var carFactory TransportFactory = &CarFactory{}
// 	car := carFactory.CreateTransport()
// 	fmt.Println(car.Drive()) // Output: Driving a car

// 	// Создаем фабрику для мотоциклов
// 	var motorcycleFactory TransportFactory = &MotorcycleFactory{}
// 	motorcycle := motorcycleFactory.CreateTransport()
// 	fmt.Println(motorcycle.Drive()) // Output: Riding a motorcycle
// }
