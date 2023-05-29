package tasks

import (
	"fmt"
	"time"

	"tabelf/backend/service/app"
)

func SayHello(ctx JobContext, config app.Config) {
	fmt.Print("倒计时···\r\n")
	for count := 5; count > 0; count-- {
		fmt.Printf("%d\r\n", count)
		time.Sleep(time.Second)
	}
	fmt.Print("Hello\r\n")
}
