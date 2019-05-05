package main

import (
	"encoding/json"
)

//代表
type delegate struct{}

//节点元数据
func (d *delegate) NodeMeta(limit int) []byte {
	return []byte{}
}

//通知消息
func (d *delegate) NotifyMsg(b []byte) {
	if len(b) == 0 {
		return
	}

	//将通讯数据的头字节取出(自定义数据协议)
	switch b[0] {
	case 'd':
		//定义数据包列表
		var executes []*Execute
		//绑定数据到数据包列表
		if err := json.Unmarshal(b[1:], &executes); err != nil {
			return
		}
		//遍历取单个数据包操作
		for _, u := range executes {
			//把map[string]string的Data数据取出
			for k, v := range u.Data {
				switch u.Cmd {
				case "add":
					data.Set(k, v)
				case "del":
					data.Delete(k)
				}
			}
		}
	}
}

//获取广播
func (d *delegate) GetBroadcasts(overhead, limit int) [][]byte {
	return nh.broadcasts.GetBroadcasts(overhead, limit)
}

//本地状态，将数据转化成JSON数据返回
func (d *delegate) LocalState(join bool) []byte {
	m := data.GetItems()
	b, _ := json.Marshal(m)
	return b
}

//合并远程状态
func (d *delegate) MergeRemoteState(buf []byte, join bool) {
	if len(buf) == 0 {
		return
	}
	if !join {
		return
	}
	var m map[string]string
	if err := json.Unmarshal(buf, &m); err != nil {
		return
	}
	for k, v := range m {
		data.Set(k, v)
	}
}
