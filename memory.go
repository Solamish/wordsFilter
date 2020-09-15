package wordsfilter

import (
	"bufio"
	. "github.com/orcaman/concurrent-map"
	"log"
	"os"
	"strings"
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

// WriteFile can load multiple dictionary files,
// the file name separated by "," or ", "
// the front of the dictionary preferentially load the participle,
//	such as: "user_dictionary.txt,common_dictionary.txt"
func (ms *MemoryStore) WriteFile(path string) error{
	if path == "" {
		return nil
	}

	dictPaths := DicPaths(path)
	if len(dictPaths) > 0 {
		for i := 0; i < len(dictPaths); i++ {
			file, err := os.Open(dictPaths[i])
			defer file.Close()
			if err != nil {
				log.Println("Open file error:",err)
				return err
			}

			scanner := bufio.NewScanner(file)
			for scanner.Scan(){
				word := scanner.Text()
    			ms.data.Set(word,1)
			}
		}
	}
	return nil
}

// DictPaths get the dict's paths
func DicPaths(filePath string) (paths []string) {
	var dicPath string

	var fileName []string
	if strings.Contains(filePath,", ") {
		fileName = strings.Split(filePath,", ")
	} else  {
		fileName = strings.Split(filePath, ",")
	}

	for i := 0; i < len(fileName); i++ {
		dicPath = fileName[i]

		if dicPath != ""{
			paths = append(paths,dicPath)
		}
	}

	return
}
