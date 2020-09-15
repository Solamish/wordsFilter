package wordsfilter

import (
	"fmt"
	"testing"
)

func TestFilter(t *testing.T)  {
	sensitiveList := []string{"世界","全世界"}
	input := "你好，世界"
	filter := NewFilter(Opts{
		Path:       "",
		DataSource: sensitiveList,
	})
	if filter.IsMatch(input) == false {
		t.Errorf("Expected true, but got false")
	}
}

func TestFilter2(t *testing.T) {
	input := "的撒大QQ"
	filter := NewFilter(Opts{
		Path:       "广告.txt, 政治类.txt, 色情类.txt",
	})
	words := filter.store.Read()
	fmt.Println(len(words))
	if filter.IsMatch(input) == false {
		t.Errorf("Expected true, but got false")
	}
}