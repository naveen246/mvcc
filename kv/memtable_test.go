package kv

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPutGet(t *testing.T) {
	kvPairs := []struct {
		key   VersionedKey
		value Value
	}{
		{
			NewVersionedKey([]byte("aaa"), 1),
			NewValue([]byte("111")),
		},
		{
			NewVersionedKey([]byte("aaa"), 2),
			NewValue([]byte("222")),
		},
	}

	memtable := NewMemtable()
	for _, kvPair := range kvPairs {
		memtable.Put(kvPair.key, kvPair.value)
		val, ok := memtable.Get(kvPair.key)
		assert.True(t, ok)
		assert.Equal(t, kvPair.value, val)
	}
}
