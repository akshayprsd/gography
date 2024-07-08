package geo

import (
	"fmt"
	"gography/core"
	"math"
)

type Coords struct {
	Lat float64
	Lng float64
}

func (c Coords) string() string {
	return fmt.Sprintf("LatLng(%f, %f)", c.Lat, c.Lng)
}

func (c Coords) Equals(other Coords) bool {
	return c.Lat == other.Lat && c.Lng == other.Lng
}

func (c Coords) Distance(other Coords) float64 {
	const R = 6371 // Earth radius in kilometers
	lat1 := core.ToRad(c.Lat)
	lat2 := core.ToRad(other.Lat)
	deltaLat := core.ToRad(other.Lat - c.Lat)
	deltaLng := core.ToRad(other.Lng - c.Lng)

	a := math.Sin(deltaLat/2)*math.Sin(deltaLat/2) +
		math.Cos(lat1)*math.Cos(lat2)*
			math.Sin(deltaLng/2)*math.Sin(deltaLng/2)
	circumference := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	distance := R * circumference
	return distance // in kms
}

func (c Coords) Midpoint(other Coords) Coords {
	lat1 := core.ToRad(c.Lat)
	lng1 := core.ToRad(c.Lng)
	lat2 := core.ToRad(other.Lat)

	dLng := core.ToRad(other.Lng - c.Lng)

	bx := math.Cos(lat2) * math.Cos(dLng)
	by := math.Cos(lat2) * math.Sin(dLng)
	lat3 := math.Atan2(math.Sin(lat1)+math.Sin(lat2),
		math.Sqrt((math.Cos(lat1)+bx)*(math.Cos(lat1)+bx)+by*by))
	lng3 := lng1 + math.Atan2(by, math.Cos(lat1)+bx)

	return Coords{Lat: core.ToDeg(lat3), Lng: core.ToDeg(lng3)}
}

func (c Coords) Validate() bool {
	return c.Lat >= -90 && c.Lat <= 90 && c.Lng >= -180 && c.Lng <= 180
}
