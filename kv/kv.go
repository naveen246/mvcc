package kv

type VersionedKey struct {
	Key     []byte
	Version uint64
}

func NewVersionedKey(key []byte, version uint64) VersionedKey {
	return VersionedKey{
		Key:     key,
		Version: version,
	}
}

func (k VersionedKey) Bytes() []byte {
	data := make([]byte, 0)
	data = append(data, k.Key...)
	data = append(data, byte(k.Version))
	return data
}

type Value struct {
	Data []byte
}

func NewValue(b []byte) Value {
	return Value{b}
}
