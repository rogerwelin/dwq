package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

func return_random(num int) string {
	alphabet := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	var bytes = make([]byte, num)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = alphabet[b%byte(len(alphabet))]
	}

	return string(bytes)

}

func main() {
	randomInt := os.Args[1]
	val_toi, err := strconv.Atoi(randomInt)

	if err != nil {
		fmt.Println("Run program with only one integer argument")
		log.Fatal(err)
	}

	ret_val := return_random(val_toi)
	fmt.Println(ret_val)

}
