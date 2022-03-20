package storage

type storage struct {
}

func NewStorage() Storage {
	s := storage{}

	return &s
}

func (s *storage) Open(name string) error {
	return nil
}

func (s *storage) Close() {}

func (s *storage) Put(key []byte, value []byte) {

}

func (s *storage) Delete(key []byte) {

}

func (s *storage) Get(key []byte) []byte {
	return nil
}

func (s *storage) MultiGet([][]byte) [][]byte {
	return nil
}

func (s *storage) KeyMayExist(key []byte) bool {
	return false
}
