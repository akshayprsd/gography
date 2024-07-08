package mapCore

import (
	"gography/core"
	"gography/geo"
)

type Map struct {
	core.Evented
	origin geo.Coords
	zoom   int
	// add layer array
}

func NewMap(origin geo.Coords, zoom int) *Map {
	return &Map{
		origin: origin,
		zoom:   zoom,
	}
}

func (m *Map) SetView(origin geo.Coords, zoom int) {
	m.origin = origin
	m.zoom = zoom
	m.Fire("viewchange", nil)
}

func (m *Map) GetView() (geo.Coords, int) {
	return m.origin, m.zoom
}

func (m *Map) ZoomIn(amount int) {
	m.zoom += amount
	m.Fire("viewchange", nil)
}

func (m *Map) ZoomOut(amount int) {
	m.zoom -= amount
	m.Fire("viewchange", nil)
}

func (m *Map) Pan(deltaLat, deltaLng float64) {
	m.origin.Lat += deltaLat
	m.origin.Lng += deltaLng
	m.Fire("viewchange", nil)
}

func (m *Map) ResetView(defaultOrigin geo.Coords, defaultZoom int) {
	m.origin = defaultOrigin
	m.zoom = defaultZoom
	m.Fire("viewchange", nil)
}

func (m *Map) OnViewChange(callback func(core.Event)) {
	m.On("viewchange", callback)
}
