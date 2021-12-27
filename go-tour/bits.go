package main

import (
	"fmt"
	"strconv"
)

func bits() {
	fmt.Println(0b100)                                      // 4
	fmt.Println(byte(0b100))                                // 4
	fmt.Println(byte(4))                                    // 4
	fmt.Printf("Print as binary using b %b \n", 4)          // 100
	fmt.Printf("Print as binary using b %b \n", 0b100)      // 100
	fmt.Printf("Right shift binary 1001 %b \n", 0b1001>>1)  // 100
	fmt.Printf("Right shift decimal 1001 %b \n", 0b1001>>1) // 100

	//decNum := 6
	//var binNum string = fmt.Sprintf("%b", decNum)
	//fmt.Printf("Convert dec -> bin %v \n", binNum) // 100

	// Get decimal number from binary string
	parsedInt, _ := strconv.ParseInt("1101", 2, 64)
	fmt.Printf("%d", parsedInt)
}
