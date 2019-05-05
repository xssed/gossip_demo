package main

import (
	"sync"
)

//Key/Value数据结构
type Data struct {
	mtx   sync.RWMutex
	items map[string]string
}

//创建
func NewData() *Data {
	m := make(map[string]string)
	return &Data{items: m}
}

//获取Key数据
func (d *Data) Get(key string) string {
	d.mtx.RLock()
	defer d.mtx.RUnlock()
	val := d.items[key]
	return val
}

//获取整个数据集合
func (d *Data) GetItems() map[string]string {
	d.mtx.RLock()
	defer d.mtx.RUnlock()
	m := d.items
	return m
}

//设置Key数据
func (d *Data) Set(key, value string) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	d.items[key] = value
}

//删除Key数据
func (d *Data) Delete(key string) {
	d.mtx.Lock()
	defer d.mtx.Unlock()
	delete(d.items, key)
}
