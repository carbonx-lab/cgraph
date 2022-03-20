package storage

type Storage interface {
	Open(name string) error
	Close()

	Put(key []byte, value []byte)
	Delete(key []byte)
	Get(key []byte) []byte
	MultiGet(keys [][]byte) [][]byte
	KeyMayExist([]byte) bool
}
