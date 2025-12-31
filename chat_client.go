package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	// اتصال به سرور در پورت 8080
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting to server:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server on port 8080")

	// خواندن ورودی از کاربر
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter message: ")
		scanner.Scan()
		message := scanner.Text()

		if message == "exit" {
			fmt.Println("Exiting...")
			break
		}

		// ارسال پیام به سرور
		_, err := conn.Write([]byte(message + "\n"))
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}
	}
}
