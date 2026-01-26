package models

import "github.com/twpayne/go-geom"

type FavoritePlace struct {
	ID          int
	Name        string
	Coordinates geom.Point
}

type Route struct {
	ID   int
	Path geom.LineString
}

func CrearePoints() {
	/*
		geom.NewPoint(geom.XY) — создаёт точку в 2D пространстве (X = долгота, Y = широта)
		SetSRID(4326) — задаёт координатную систему. 4326 = WGS84, стандарт GPS (широта/долгота в градусах)
		MustSetCoords(geom.Coord{...}) — устанавливает координаты точки
	*/
	pt := geom.NewPoint(geom.XY).SetSRID(4326).MustSetCoords(geom.Coord{
		45.52515498907135,
		-73.57520521798992,
	})

	/*
		geom.NewLineString(geom.XY) — создаёт линию, соединяющую несколько точек (маршрут)
		SetSRID(4326) — опять указываем систему координат GPS
		MustSetCoords([]geom.Coord{...}) — передаём массив точек, из которых строится линия
	*/
	line := geom.NewLineString(geom.XY).SetSRID(4326).MustSetCoords([]geom.Coord{
		{45.52515498907135, -73.57520521798992},
		{45.52433716241833, -73.57591484495788},
		{45.523244861127694, -73.57358389203318},
		{45.52071221016285, -73.57593660297117},
	})

}
