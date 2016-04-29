package geojson_utils

import (
	"math"

	"github.com/paulmach/go.geojson"
)

var (
	RADIUS = float64(6378137)
)

func CalculateArea(geometry *geojson.Geometry) float64 {
	area := float64(0)
	switch geometry.Type {
	case geojson.GeometryPolygon:
		return polygonArea(geometry.Polygon)
	case geojson.GeometryMultiPolygon:
		for _, coords := range geometry.MultiPolygon {
			area += polygonArea(coords)
		}
	case geojson.GeometryCollection:
		for _, geom := range geometry.Geometries {
			area += CalculateArea(geom)
		}
	}
	return area
}

func polygonArea(coords [][][]float64) float64 {
	area := float64(0)

	if len(coords) > 0 {
		area += math.Abs(ringArea(coords[0]))
		for i := 1; i < len(coords); i++ {
			area -= math.Abs(ringArea(coords[i]))
		}
	}

	return area
}

func ringArea(ring [][]float64) float64 {
	p1, p2, p3 := []float64{}, []float64{}, []float64{}
	lowerIndex, middleIndex, upperIndex := 0, 0, 0

	area := float64(0)
	coordsLength := len(ring)

	if coordsLength > 2 {
		for i := 0; i < coordsLength; i++ {
			if i == (coordsLength - 2) {
				lowerIndex = coordsLength - 2
				middleIndex = coordsLength - 1
				upperIndex = 0
			} else if i == (coordsLength - 1) {
				lowerIndex = coordsLength - 1
				middleIndex = 0
				upperIndex = 1
			} else {
				lowerIndex = i
				middleIndex = i + 1
				upperIndex = i + 2
			}

			p1 = ring[lowerIndex]
			p2 = ring[middleIndex]
			p3 = ring[upperIndex]
			area += (rad(p3[0]) - rad(p1[0])) * math.Sin(rad(p2[1]))
		}

		area *= RADIUS * RADIUS / float64(2)
	}

	return area
}

func rad(p float64) float64 {
	return p * math.Pi / 180
}
