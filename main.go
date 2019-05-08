package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/xssed/owlcache/queue"
)

var (
	data    = NewData()
	members = flag.String("members", "", "comma seperated list of members")
	port    = flag.Int("port", 4002, "http port")
	nh      = NewHandler()
	q       = queue.New()
)

func init() {
	flag.Parse()
}

func main() {
	if err := nh.Start(); err != nil {
		fmt.Println(err)
	}

	go func() {
		for {
			time.Sleep(time.Microsecond * 8)
			size := q.Size()
			if size >= 1 {
				e := q.Pop()
				//fmt.Println("结果:", e)
				if e != nil {

					var result Execute
					v, convert_ok := e.(string)
					if convert_ok {
						//fmt.Println("string:", v)
						if err := json.Unmarshal([]byte(v), &result); err != nil {
							fmt.Println(err)
						}
						//fmt.Println("json to map ", result)
					}

					switch result["cmd"] {
					case "add":
						data.Set(result["key"], result["val"])
					case "del":
						data.Delete(result["key"])
					}

				}
			}
		}
	}()

	http.HandleFunc("/add", addHandler)
	http.HandleFunc("/del", delHandler)
	http.HandleFunc("/get", getHandler)
	fmt.Printf("Listening on :%d\n", *port)
	if err := http.ListenAndServe(fmt.Sprintf(":%d", *port), nil); err != nil {
		fmt.Println(err)
	}

}
