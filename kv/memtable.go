package kv

import (
	"github.com/huandu/skiplist"
	"sync"
)

type Memtable struct {
	sync.RWMutex
	skl *skiplist.SkipList
}

func NewMemtable() *Memtable {
	return &Memtable{
		skl: skiplist.New(skiplist.Bytes),
	}
}

func (m *Memtable) Put(key VersionedKey, value Value) {
	m.Lock()
	defer m.Unlock()
	m.skl.Set(key.Bytes(), value.Data)
}

func (m *Memtable) Get(key VersionedKey) (Value, bool) {
	m.RLock()
	defer m.RUnlock()
	elem := m.skl.Get(key.Bytes())
	if elem == nil {
		return Value{}, false
	}
	return Value{elem.Value.([]byte)}, true
}
