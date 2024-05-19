package distance_find

import (
	"testing"
)

func Test_DistanceBetweenPoints(t *testing.T) {
	t.Run("Should give the distance between two point given 4 input values as coordinates", func(t *testing.T) {
		point1 := point{0, 0}
		point2 := point{3, 4}
		got := DistanceBetweenPoints(point1, point2)
		expected := float64(5)
		if got != expected {
			t.Errorf("Expected %f got %f", expected, got)
		}
	})
}
