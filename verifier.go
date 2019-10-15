package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	arg := os.Args[1]

	checkValidNumber( arg )

}

func checkValidNumber( arg string ) {

	cardNumber := strings.ReplaceAll( arg, " ", "" )

	var sumNumbers int64 = 0

	var slicedNumber int64

	if len(cardNumber) != 16 {
		fmt.Println("NO")
		return
	}

	for i := 0; i < 16; i++ {
		slicedNumber, _ = strconv.ParseInt(cardNumber[i:i+1], 10, 64)
		if i % 2 == 0 {
			var dubbedNumber int64 = slicedNumber * 2
			if dubbedNumber > 9 {
				slicedNumber = dubbedNumber - 9
			} else {
				slicedNumber = dubbedNumber
			}
		}
		sumNumbers = sumNumbers + slicedNumber
	}

	if sumNumbers % 10 == 0 {
		fmt.Println("YES")
	}

}
