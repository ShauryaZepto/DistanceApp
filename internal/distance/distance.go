package distance_find

import (
	distance_find "distance_find/internal/point"
	"math"
)

func DistanceBetweenPoints(p1, p2 *distance_find.Point) float64 {
	return math.Sqrt(float64((p1.X-p2.X)*(p1.X-p2.X) + (p1.Y-p2.Y)*(p1.Y-p2.Y)))
}
