package main

import (
	"flag"
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
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

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "index here")

}

func req_calc(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	numId := vars["numId"]
	numInt, err := strconv.Atoi(numId)
	if err != nil {
		log.Error(err)
	}
	retVal := return_random(numInt)

	fmt.Fprintln(w, "Results:", retVal)
}

func main() {
	var path = flag.String("f", "log-path.txt", "string")
	var times = flag.Int("t", 10, "int")
	var http_bool = flag.Bool("http", false, "string")

	flag.Parse()
	//fmt.Println("path:", *path)
	//fmt.Println("times:", *times)

	f, err := os.OpenFile(*path, os.O_WRONLY|os.O_CREATE, 0755)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(f)

	if *http_bool == true {
		router := mux.NewRouter().StrictSlash(true)
		router.HandleFunc("/", index)
		router.HandleFunc("/req/{numId}", req_calc)
		log.Fatal(http.ListenAndServe(":8080", router))
	} else {
		ret_val := return_random(*times)
		fmt.Println(ret_val)
		log.Info(ret_val)
	}

}
