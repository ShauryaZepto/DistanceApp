package distance_find

import (
	"testing"
)

func Test_DistanceBetweenPoints(t *testing.T) {
	t.Run("Should give the distance between two point given 4 input values as coordinates", func(t *testing.T) {
		got := DistanceBetweenPoints(0, 0, 3, 4)
		expected := float64(5)
		if got != expected {
			t.Errorf("Expected %f got %f", expected, got)
		}
	})
}
