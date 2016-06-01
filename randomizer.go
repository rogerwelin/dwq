package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"math/rand"
	"os"
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
	var path = flag.String("f", "log-path.txt", "string")
	var times = flag.Int("t", 10, "int")
	flag.Parse()
	fmt.Println("path:", *path)
	fmt.Println("times:", *times)

	f, err := os.OpenFile(*path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)

	ret_val := return_random(*times)
	fmt.Println(ret_val)
	log.Info(ret_val)

}
