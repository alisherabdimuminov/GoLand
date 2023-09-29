// usr/bin/go
// author: Ali
// date  : 29/09/2023
package main

import (
	"fmt"
	"time"
)

func main() {
	// Go dasturlash tilida time
	// import (... "time")
	now := time.Now()
	fmt.Println(now)

	// bugungi kunni olish
	day := now.Day()
	fmt.Println("Kun:", day)

	// hozirgi oyni olish
	month := now.Month()
	fmt.Println("Oy:", month)

	// hozirgi yilni olish
	year := now.Year()
	fmt.Println("Yil:", year)

	// hozirgi soat
	hour := now.Hour()
	fmt.Println("Soat:", hour)

	// hozirgi soat
	minute := now.Minute()
	fmt.Println("Soat:", minute)

	// hozirgi soat
	second := now.Second()
	fmt.Println("Soat:", second)

	fmt.Printf("%d-%d-%d %d:%d:%d\n", year, month, day, hour, minute, second)

	// format yordamida vaqtni chiraqish
	const (
		format = "2006-01-02 15:04:05"
	)
	fmt.Println(now.Format(format))

}
