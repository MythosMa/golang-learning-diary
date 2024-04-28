package main

import (
	"sync"
)

type UserInfo struct {
	mu   sync.Mutex
	name string
	age  int
}

func update(userInfo *UserInfo, newInfo *UserInfo) {
	userInfo.mu.Lock()
	userInfo.name = newInfo.name
	userInfo.age = newInfo.age
	userInfo.mu.Unlock()
}
