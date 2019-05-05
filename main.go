package main

import (
	"flag"
	"fmt"
	"net/http"
)

var (
	data    = NewData()
	members = flag.String("members", "", "comma seperated list of members")
	port    = flag.Int("port", 4004, "http port")
	nh      = NewHandler()
)

func init() {
	flag.Parse()
}

func main() {
	if err := nh.Start(); err != nil {
		fmt.Println(err)
	}

	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/del", delHandler)
	http.HandleFunc("/get", getHandler)
	fmt.Printf("Listening on :%d\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		fmt.Println(err)
	}
}
