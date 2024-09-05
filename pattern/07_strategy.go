package pattern

import "fmt"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

// Интерфейс для стратегий оплаты
type PaymentStrategy interface {
	Pay(amount float64) string
}

// Стратегия оплаты кредитной картой
type CreditCard struct {
	owner      string
	cardNumber string
	expiration string
	cvv        string
}

func (cc *CreditCard) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Credit Card [%s]", amount, cc.cardNumber)
}

// Стратегия оплаты через PayPal
type PayPal struct {
	email string
}

func (pp *PayPal) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using PayPal [%s]", amount, pp.email)
}

// Стратегия оплаты через Google Pay
type GooglePay struct {
	accountID string
}

func (gp *GooglePay) Pay(amount float64) string {
	return fmt.Sprintf("Paid %.2f using Google Pay [%s]", amount, gp.accountID)
}

// Контекст, который использует стратегию оплаты
type PaymentContext struct {
	strategy PaymentStrategy
}

// Установка стратегии
func (p *PaymentContext) SetPaymentStrategy(strategy PaymentStrategy) {
	p.strategy = strategy
}

// Вызов метода оплаты через выбранную стратегию
func (p *PaymentContext) Pay(amount float64) {
	fmt.Println(p.strategy.Pay(amount))
}

// func main() {
// 	// Создаем контекст для оплаты
// 	context := &PaymentContext{}

// 	// Оплата через кредитную карту
// 	creditCard := &CreditCard{
// 		owner:      "John Doe",
// 		cardNumber: "1234-5678-9876-5432",
// 		expiration: "12/24",
// 		cvv:        "123",
// 	}
// 	context.SetPaymentStrategy(creditCard)
// 	context.Pay(100.0) // Output: Paid 100.00 using Credit Card [1234-5678-9876-5432]

// 	// Оплата через PayPal
// 	paypal := &PayPal{
// 		email: "john@example.com",
// 	}
// 	context.SetPaymentStrategy(paypal)
// 	context.Pay(200.0) // Output: Paid 200.00 using PayPal [john@example.com]

// 	// Оплата через Google Pay
// 	googlePay := &GooglePay{
// 		accountID: "john-google-123",
// 	}
// 	context.SetPaymentStrategy(googlePay)
// 	context.Pay(150.0) // Output: Paid 150.00 using Google Pay [john-google-123]
// }
