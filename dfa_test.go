package wordsfilter

import (
	"testing"
)

func TestIsMatch(t *testing.T) {
	sensitiveList := []string{"世界","全世界"}
	input := "你好，世界"

	dfa := NewDFAUtil(sensitiveList)
	if dfa.IsMatch(input) == false {
		t.Errorf("Expected true, but got false")
	}
}

func TestHandleWord(t *testing.T) {
	sensitiveList := []string{"世界","世界的"}
	input := "你好,世界"

	dfa := NewDFAUtil(sensitiveList)
	newInput := dfa.HandleWord(input,'*')
	expected := "你好,**"
	if newInput != expected {
		t.Errorf("Expected %s, but got %s", expected, newInput)
	}
}

