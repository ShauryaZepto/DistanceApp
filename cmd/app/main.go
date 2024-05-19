package main

import (
	distance_find "distance_find/internal/distance"
	distance_point "distance_find/internal/point"
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Println("Enter the coordinate of the first point (x1 y1):")
	fmt.Scanln(&x1, &y1)
	fmt.Println("Enter the coordinate of the second point (x2 y2):")
	fmt.Scanln(&x2, &y2)
	point1 := distance_point.NewPoint(x1, y1)
	point2 := distance_point.NewPoint(x2, y2)
	distance := distance_find.DistanceBetweenPoints(point1, point2)
	fmt.Printf("The Distance between %v and %v is %f \n", *point1, *point2, distance)

}
