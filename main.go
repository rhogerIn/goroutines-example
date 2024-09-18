package main

import (
	"fmt"
	"time"
)

func processPayment(paymentId string, amount float64, ch chan<- string) {
	time.Sleep(1 * time.Second)
	ch <- fmt.Sprintf("Payment %s of $%.2f processed.", paymentId, amount)
}

func main() {
	payments := []struct {
		id		string
		amount	float64
	}{
		{"PAYMENT123", 87.90},
		{"PAYMENT456", 187.90},
		{"PAYMENT789", 817.90},
	}
	ch := make(chan string)
	for _, payment := range payments {
		go processPayment(payment.id, payment.amount, ch)
	}
	for range payments {
		fmt.Println(<-ch)
	}
}
