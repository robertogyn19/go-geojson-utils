package geojson_utils

import (
	"testing"

	"fmt"
	"github.com/paulmach/go.geojson"
	"github.com/stretchr/testify/assert"
)

var (
	p1 = []float64{30.585937499999996, 35.746512259918504}
	p2 = []float64{-66.796875, -18.979025953255253}
	p3 = []float64{16.171875, -17.97873309555617}

	pl1 = [][][]float64{[][]float64{p1, p2, p3, p1}}
	a1  = 26214016516278.5430

	p4 = []float64{-37.6171875, 68.13885164925573}
	p5 = []float64{-2.109375, 56.559482483762245}
	p6 = []float64{15.8203125, 73.22669969306126}

	pl2 = [][][]float64{[][]float64{p4, p5, p6, p4}}
	a2  = 2146389996641.2275

	point        = geojson.NewPointGeometry(p1)
	multiPoint   = geojson.NewMultiPointGeometry(p1, p2, p3, p4, p5, p6)
	line         = geojson.NewLineStringGeometry([][]float64{p1, p2})
	multiLine    = geojson.NewMultiLineStringGeometry([][]float64{p1, p2}, [][]float64{p3, p4})
	polygon1     = geojson.NewPolygonGeometry(pl1)
	polygon2     = geojson.NewPolygonGeometry(pl2)
	multiPolygon = geojson.NewMultiPolygonGeometry(pl1, pl2)
	collection   = geojson.NewCollectionGeometry(polygon1, polygon2)

	complexGeometry = []byte(`
  {
    "type": "Polygon",
    "coordinates": [
      [
        [
          -50.33935546875,
          -15.654775665159686
        ],
        [
          -51.56982421875,
          -16.85711965391805
        ],
        [
          -47.98828124999999,
          -16.878146994732155
        ],
        [
          -48.394775390625,
          -14.785505314974664
        ],
        [
          -50.196533203125,
          -18.35452552912664
        ],
        [
          -47.823486328125,
          -17.853290114098012
        ],
        [
          -50.33935546875,
          -15.654775665159686
        ]
      ]
    ]
  }`)
	a3 = 58668714469.1091842651
)

func TestCalculateArea(t *testing.T) {
	executeTest(t, point, float64(0))
}

func TestCalculateArea2(t *testing.T) {
	executeTest(t, multiPoint, float64(0))
}

func TestCalculateArea3(t *testing.T) {
	executeTest(t, line, float64(0))
}

func TestCalculateArea4(t *testing.T) {
	executeTest(t, multiLine, float64(0))
}

func TestCalculateArea5(t *testing.T) {
	executeTest(t, polygon1, a1)
	executeTest(t, polygon2, a2)
}

func TestCalculateArea6(t *testing.T) {
	executeTest(t, multiPolygon, a1+a2)
}

func TestCalculateArea7(t *testing.T) {
	executeTest(t, collection, a1+a2)
}

func TestCalculateArea8(t *testing.T) {
	geom, _ := geojson.UnmarshalGeometry(complexGeometry)
	executeTest(t, geom, a3)
}

func executeTest(t *testing.T, geom *geojson.Geometry, expected float64) {
	area := CalculateArea(geom)
	assert.Equal(t, fmt.Sprintf("%.10f", expected), fmt.Sprintf("%.10f", area))
}
