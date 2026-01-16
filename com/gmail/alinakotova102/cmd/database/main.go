package database

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const pgConn = "host=localhost user=postgres password=postgres dbname=gopg port=5433 sslmode=disable"

func GetDB() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(pgConn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateTables(db *gorm.DB) {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS postgis;`) // подключает расширение PostGIS (посчитает реальное расстояние между двумя точками на Земле в метрах)

	/*
		geography - тип данных для географических объектов на Земле (с учётом кривизны планеты)
		POINT → точка на карте (например: координаты кафе, дома, станции метро)
	*/
	db.Exec(`CREATE TABLE IF NOT EXISTS "favorite_places" (
	"id" INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	"name" TEXT NOT NULL,
	"coordinates" geography(POINT, 4326) NOT NULL);`)

	// LINESTRING — маршрут из множества точек
	db.Exec(`CREATE TABLE IF NOT EXISTS "routes" (
	"id" INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
	"path" geography(LINESTRING, 4326) NOT NULL);`)
}
