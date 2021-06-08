package store

import "sync"

// Simple KV store taken from https://github.com/goraft/raftd/blob/master/db/db.go
type DB struct {
	data  map[string]string
	mutex sync.RWMutex
}

func NewDB() *DB {
	return &DB{
		data: make(map[string]string),
	}
}

func (db *DB) Size() int {
	return len(db.data)
}

func (db *DB) Get(key string) string {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	return db.data[key]
}

func (db *DB) Put(key string, value string) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.data[key] = value
}

func (db *DB) Clear() {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.data = make(map[string]string)
}
