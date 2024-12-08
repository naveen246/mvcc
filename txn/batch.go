package txn

type KV struct {
	Key   []byte
	Value []byte
}

func newKV(key, value []byte) KV {
	return KV{key, value}
}

type Batch struct {
	kvPairs []KV
}
