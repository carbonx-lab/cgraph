package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_storage_Open(t *testing.T) {
	s := NewStorage()

	err := s.Open("db1")
	assert.Nil(t, err)
}

func Test_storage_Close(t *testing.T) {
	s := NewStorage()

	s.Close()
}

func Test_storage_Delete(t *testing.T) {

}

func Test_storage_Get(t *testing.T) {

}

func Test_storage_KeyMayExist(t *testing.T) {

}

func Test_storage_MultiGet(t *testing.T) {

}

func Test_storage_Put(t *testing.T) {

}
