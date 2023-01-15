package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	msg := "你嘅運財號碼係: "
	rand.Seed(time.Now().UnixNano()) //create random seed
	luckyNumbers := rand.Perm(49)    //create unique random numbers from 0 to 48
	for i := 0; i < 6; i++ {         //for saving memory, we only need the first 6 numbers
		luckyNumbers[i] = luckyNumbers[i] + 1     //add 1 to all numbers to make it become 1 to 49
		msg = msg + strconv.Itoa(luckyNumbers[i]) //convert int to string
		if i == 5 {
			break
		}
		msg = msg + ", "
	}

	fmt.Println(msg)

}
