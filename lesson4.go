// usr/bin/go
// author: Ali
// date  : 21/09/2023
package main

import "fmt"

func main() {
	// Go dasturlash tilida asosiy ma'lumot turlari
	/*
		bool   -> true yoki false                                    | 1 bayt
		strnig -> belgilar ketma-ketligi                             | 1 bayt dan 4 baytgacha
		byte   -> byte == int8                                       | 8 bit
		rune   -> rune == int32                                      | 32 bit

		int8   -> -128 dan 127 gacha                                 | 8 bit unsigned
		int16  -> -32768 dan 32767 gacha                             | 16 bit unsigned
		int32  -> -2147483648 dan 2147483647 gacha                   | 32 bit unsigned
		int64  -> -9223372036854775808 dan 9223372036854775807 gacha | 64 bit unsigned

		uint8  -> 0 dan 255 gacha									 | 8 bit
		uint16 -> 0 dan 65535 gacha                       	    	 | 16 bit
		uint32 -> 0 dan 4294967295 gacha						     | 32 bit
		uint64 -> 0 dan 18446744073709551615 gacha			     	 | 64 bit

		float32 -> -3.4e+38 dan 3.4e+38 gacha						 | 64 bit
		float64 -> -1.7e+308 dan +1.7e+308 gacha                     | 32 bit
	*/

	// bool
	var isGopher bool = true
	var isPythonist bool = false
	fmt.Println("We are Go developer: ", isGopher)
	fmt.Println("We are Pythonist:    ", isPythonist)

	// string
	var gostring string = "It's Gopher!"
	fmt.Println(gostring)

	// int
	var a int = 2023
	var b int = 2005
	fmt.Println(a + b)

	var maxInt uint64 = 1<<63 - 1
	fmt.Printf("Mening pullarim %d$ :)\n", maxInt)
}
