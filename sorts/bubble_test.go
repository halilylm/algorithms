package sorts

import (
	"reflect"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	arr := []int{4, -2, 6, 5}
	BubbleSort(&arr)
	expected := []int{-2, 4, 5, 6}
	if !reflect.DeepEqual(arr, expected) {
		t.Errorf("expected %v got %v", expected, arr)
	}
}
