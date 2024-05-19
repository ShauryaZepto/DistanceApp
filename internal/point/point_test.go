package distance_find

import (
	"reflect"
	"testing"
)

func Test_NewPoint(t *testing.T) {
	t.Run("Should create a new point given two inputs", func(t *testing.T) {
		expect := &Point{1, 5}
		got := NewPoint(1, 5)
		if !reflect.DeepEqual(expect, got) {
			t.Errorf("Expected %v got %v", expect, got)
		}
	})
}
