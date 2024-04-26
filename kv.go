package main

import "sync"

type KV struct {
	mu   *sync.RWMutex
	data map[string][]byte
}

func NewKV() *KV {
	return &KV{
		data: make(map[string][]byte),
		mu:   &sync.RWMutex{},
	}
}

func (kv *KV) Set(key string, val string) error {
	kv.mu.Lock()
	defer kv.mu.Unlock()
	kv.data[key] = []byte(val)
	return nil
}

func (kv *KV) Get(key string) ([]byte, bool) {
	kv.mu.RLock()
	defer kv.mu.RUnlock()
	val, ok := kv.data[key]
	return val, ok
}
