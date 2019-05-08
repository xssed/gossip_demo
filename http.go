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

	exedata := make(Execute)
	exedata["cmd"] = "add"
	exedata["key"] = key
	exedata["val"] = val

	var exedata_list []*Execute
	exedata_list = append(exedata_list, &exedata)

	b, err := json.Marshal(exedata_list)

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

	exedata := make(Execute)
	exedata["cmd"] = "del"
	exedata["key"] = key
	exedata["val"] = ""

	var exedata_list []*Execute
	exedata_list = append(exedata_list, &exedata)

	b, err := json.Marshal(exedata_list)

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
