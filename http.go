package main

import (
	"encoding/json"
	"net/http"
)

func addHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	val := r.Form.Get("val")
	data.Set(key, val)

	b, err := json.Marshal([]*Execute{
		&Execute{
			Cmd: "add",
			Data: map[string]string{
				key: val,
			},
		},
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	//发送数据到集群
	nh.QueueBroadcast(b)
}

func delHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	data.Delete(key)

	b, err := json.Marshal([]*Execute{
		&Execute{
			Cmd: "del",
			Data: map[string]string{
				key: "",
			},
		},
	})

	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	//发送数据到集群
	nh.QueueBroadcast(b)
}

func getHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	key := r.Form.Get("key")
	val := data.Get(key)
	w.Write([]byte(val))
}
