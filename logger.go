package logger

import (
	"fmt"
	"time"
)

var formatTime = time.Now().Format(time.RFC1123)

func RegInfo(email string) {
	fmt.Println("---INFO---\n", formatTime, "user ", email, "registered successfully!")
}

func RegErr(email string, err error) {
	fmt.Println("---ERROR---\n", formatTime, "user ", email, "can't register with\n", err)
}

func ChangePassInfo(email string) {
	fmt.Println("---INFO---\n", formatTime, "user ", email, "changed successfully!")
}

func ChangePassErr(email string, err error) {
	fmt.Println("---ERROR---\n", formatTime, "user ", email, "can't change with\n", err)
}

func DBErr(message string, err error) {
	fmt.Println("---ERROR---\n", formatTime, message, "\n", err)
}

func PaidInfo(email string, payment float32) {
	fmt.Println("---INFO---\n", formatTime, "user ", email, " successfully paid ", payment)
}

func PaidErr(email string, payment float32, err error) {
	fmt.Println("---ERROR---\n", formatTime, "user ", email, "  payment failed ", payment, "\n", err)
}

func PaidWarning(message, email string, payment float32) {
	fmt.Printf("---WARN---\n%s \nUser %s, payment %.2f\n%s\n---WARN---", formatTime, email, payment, message)
}

func ApproveOrderInfo(email, address, payment, contact string) {
	fmt.Println("---INFO---\n", formatTime, "user ", email, " order successfully approved with address=",
		address, " payment=", payment, " contact: ", contact)
}

func ApproveOrderErr(email, payment string, err error) {
	fmt.Println("---ERROR---\n", formatTime, "user ", email,
		" with payment=", payment, " order not approved with error \n", err)
}
