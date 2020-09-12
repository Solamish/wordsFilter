package store

import (
	"bufio"
	."github.com/orcaman/concurrent-map"
	"log"
	"os"
)

type MemoryStore struct {
	data ConcurrentMap
}

type Opts struct {
	Path string
	DataSource []string
}

func NewMemoryStore(opts Opts) *MemoryStore {
	memoryStore := &MemoryStore{
		data: New(),
	}
	if opts.Path != "" {
		memoryStore.WriteFile(opts.Path)
	} else {
		for i := range opts.DataSource {
			memoryStore.Write(opts.DataSource[i])
		}
	}

	return memoryStore
}

// Write writes words to memory
func (ms *MemoryStore) Write(words ...string) {
	if len(words) == 0 {
		return
	}

	for i := range words {
		ms.data.Set(words[i],1)
	}
}

// Read reads words from memory
func (ms *MemoryStore) Read() []string{
	keys := ms.data.Keys()
	return keys
}

// Remove removes words from memory
func (ms *MemoryStore) Remove(words ...string) {
	if len(words) == 0 {
		return
	}

	for i := range words {
		ms.data.Remove(words[i])
	}
}

func (ms *MemoryStore) IsExist(word string) bool {
	_, isExist := ms.data.Get(word)
	return isExist
}

// WriteFile writes file to memory
func (ms *MemoryStore) WriteFile(path string) {
	if path == "" {
		return
	}

	file, err := os.Open(path)
	if err != nil {
		log.Println("Open file error:",err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text != "" {
			ms.data.Set(text,1)
		}
	}
}
