package distance_find

import (
	distance_find "distance_find/internal/point"
	"testing"
)

func Test_DistanceBetweenPoints(t *testing.T) {
	t.Run("Should give the distance between two point given 4 input values as coordinates", func(t *testing.T) {
		point1 := &distance_find.Point{X: 0, Y: 0}
		point2 := &distance_find.Point{X: 3, Y: 4}
		got := DistanceBetweenPoints(point1, point2)
		expected := float64(5)
		if got != expected {
			t.Errorf("Expected %f got %f", expected, got)
		}
	})
}
