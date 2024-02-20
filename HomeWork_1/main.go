package main

import (
	"fmt"
)

type OrderCustomer struct {
	name     string
	surname  string
	phoneNum string
}

type OrderGood struct {
	name  string
	price float64
	count int
}

type Order struct {
	ID        int
	Customer  *OrderCustomer
	OrderGood []*OrderGood
}

var counter = make(map[int]*Order)

func main() {

	//1
	customer1 := OrderCustomer{name: "Alyosha", surname: "Debilny", phoneNum: "8-800-555-35-35"}

	//2
	var order1 = NewOrder(&customer1)

	//3
	AddGood(order1, "Бананы", 13.58, 3)
	AddGood(order1, "Кровь девственниц", 00.01, 666)
	AddGood(order1, "Какули единорога", 999.00, 18)

	//4
	fmt.Println(ToString(order1))

	//5
	customer2 := OrderCustomer{name: "Alyosha", surname: "Debilny", phoneNum: "8-800-555-35-35"}
	var order2 = NewOrder(&customer2)
	AddGood(order2, "Энергетик", 3.57, 18)
	AddGood(order2, "Метеоритные осколки", 87.01, 34)
	AddGood(order2, "Ватные палочки", 9.01, 1)
	fmt.Println(ToString(order2))

	//6
	var validOrder = GetOrderByID(1)
	if validOrder == nil {
		fmt.Println("Некорректный идентификатор заказа!")
	} else {
		fmt.Println(ToString(validOrder))
	}

	//7
	var invalidOrder = GetOrderByID(666)
	if invalidOrder == nil {
		fmt.Println("Некорректный идентификатор заказа!")
	} else {
		fmt.Println(ToString(invalidOrder))
	}
}

func NewOrder(customer *OrderCustomer) *Order {
	var i int = len(counter) + 1
	newOrder := Order{ID: i, Customer: customer}
	counter[i] = &newOrder
	return &newOrder
}

func AddGood(order *Order, name string, price float64, count int) {
	order.OrderGood = append(order.OrderGood, &OrderGood{name, price, count})
}

func ToString(order *Order) string {
	var resultString string
	s := make([]string, len(order.OrderGood))
	for i, v := range order.OrderGood {
		s[i] += fmt.Sprintf("%+v", v)
	}
	resultString = fmt.Sprintf("ID: %d\nCustomer: %s, Phone: %s\nGoods:\n%s", order.ID, order.Customer.name+" "+order.Customer.surname, order.Customer.phoneNum, s)
	return resultString
}

func GetOrderByID(id int) *Order {
	if v, ok := counter[id]; ok {
		return v
	}
	return nil
}
