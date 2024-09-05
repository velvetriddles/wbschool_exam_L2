package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

// Интерфейс для всех обработчиков
type Handler interface {
	SetNext(handler Handler) Handler // Метод для установки следующего обработчика
	HandleRequest(request string)    // Метод для обработки запроса
}

// Базовый обработчик
type BaseHandler struct {
	next Handler
}

// Метод для установки следующего обработчика
func (h *BaseHandler) SetNext(handler Handler) Handler {
	h.next = handler
	return handler
}

// Метод для передачи запроса следующему обработчику
func (h *BaseHandler) HandleRequest(request string) {
	if h.next != nil {
		h.next.HandleRequest(request)
	}
}

// Обработчик службы поддержки
type SupportHandler struct {
	BaseHandler // Композиция базового обработчика
}

func (h *SupportHandler) HandleRequest(request string) {
	if request == "basic issue" {
		fmt.Println("Support handled the request.")
	} else {
		fmt.Println("Support passed the request to the next handler.")
		h.BaseHandler.HandleRequest(request) // Передаем запрос дальше напрямую через BaseHandler
	}
}

// Обработчик для менеджера
type ManagerHandler struct {
	BaseHandler
}

func (h *ManagerHandler) HandleRequest(request string) {
	if request == "medium issue" {
		fmt.Println("Manager handled the request.")
	} else {
		fmt.Println("Manager passed the request to the next handler.")
		h.BaseHandler.HandleRequest(request) // Передаем запрос дальше напрямую через BaseHandler
	}
}

// Обработчик для директора
type DirectorHandler struct {
	BaseHandler
}

func (h *DirectorHandler) HandleRequest(request string) {
	if request == "complex issue" {
		fmt.Println("Director handled the request.")
	} else {
		fmt.Println("Director couldn't handle the request.")
	}
}

// func main() {
// 	// Создаем обработчики
// 	support := &SupportHandler{}
// 	manager := &ManagerHandler{}
// 	director := &DirectorHandler{}

// 	// Устанавливаем цепочку: support -> manager -> director
// 	support.SetNext(manager).SetNext(director)

// 	// Пример: Запрос "basic issue"
// 	fmt.Println("Request: basic issue")
// 	support.HandleRequest("basic issue")

// 	// Пример: Запрос "medium issue"
// 	fmt.Println("\nRequest: medium issue")
// 	support.HandleRequest("medium issue")

// 	// Пример: Запрос "complex issue"
// 	fmt.Println("\nRequest: complex issue")
// 	support.HandleRequest("complex issue")
// }
