package main

import (
	"fmt"
	"gography/core"
	"gography/geo"
	mapCore "gography/map"
)

func main() {
	origin := geo.Coords{Lat: 37.7749, Lng: -122.4194} // sf bbabyyy
	zoom := 12

	m_map := mapCore.NewMap(origin, zoom)
	m_map.On("viewchange", func(e core.Event) {
		origin, zoom := m_map.GetView()
		fmt.Printf("Map view changed. Origin: %v, Zoom: %d\n", origin, zoom)
	})

	newOrigin := geo.Coords{Lat: 42.2808, Lng: -83.7430} // ann arbor : )
	newZoom := 10
	m_map.SetView(newOrigin, newZoom)
}
