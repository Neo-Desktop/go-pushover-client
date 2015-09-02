package main

import (
	"fmt"
	"log"
	"strconv"

	client "github.com/Neo-Desktop/go-pushover-client"
)

var (
	key      string
	device   string
	message  string
	priority string
	err      error
	i        int64
)

func main() {
	fmt.Printf("%s\n:", "Please enter your application key")

	_, err = fmt.Scanln(&key)
	if err != nil {
		log.Fatal(err)
	}

	pushover := client.New(key)

	fmt.Printf("\n%s\n:", "Please enter a device token")
	_, err = fmt.Scanln(&device)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n:", "Please enter a message to send")
	_, err = fmt.Scanln(&message)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n%s\n:", "Please enter a priority to send this message")
	_, err = fmt.Scanln(&priority)
	if err != nil {
		log.Fatal(err)
	}

	i, err = strconv.ParseInt(priority, 10, 64)
	if i > 2 || i < -2 {
		log.Println("Proiority should be between -2 and 2")
		log.Fatal("Message aborted")
	}

	pushover.Send(device, message, priority)
	log.Println("Message sent")

}
