package main

import (
	"bytes"
	"io"
)

type Store interface {
	Open(string) (io.ReadWriteCloser, error)
}

type StorageType int

const (
	DiskStorage StorageType = 1 << iota
	TempStorage
	MemoryStorage
)

type memoryStorage struct {
	data *bytes.Buffer
}

func newMemoryStorage() *memoryStorage {
	return &memoryStorage{data: new(bytes.Buffer)}
}

func (m *memoryStorage) Open(string) (io.ReadWriteCloser, error) {
	return m, nil
}

func (m *memoryStorage) Read(p []byte) (n int, err error) {
	return m.data.Read(p)
}

func (m *memoryStorage) Write(p []byte) (n int, err error) {
	return m.data.Write(p)
}

func (m *memoryStorage) Close() error {
	return nil
}

func NewStore(t StorageType) Store {
	switch t {
	case MemoryStorage:
		return newMemoryStorage()
	// case DiskStorage:
	// return newDiskStorage( /*...*/ )
	default:
		return nil
	}
}

func main() {
	// 메모리 저장소 생성
	s := NewStore(MemoryStorage)
	f, _ := s.Open("file")

	// 데이터 쓰기
	n, _ := f.Write([]byte("data"))
	defer func(f io.ReadWriteCloser) {
		err := f.Close()
		if err != nil {
			panic(err)
		}
	}(f)

	// 결과 출력
	println(n, "bytes written")
}
