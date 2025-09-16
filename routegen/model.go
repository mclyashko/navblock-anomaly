package main

import "time"

// Point — точка маршрута
type Point struct {
	Timestamp      string
	X              float64
	Y              float64
	SpeedKnots     float64
	HeadingDeg     int
	DepthM         float64
	SignalStrength int
}

// RouteConfig — настройки генерации маршрута
type RouteConfig struct {
	FromX, FromY float64
	ToX, ToY     float64
	Points       int
	RouteType    string // normal | abnormal
	Anomaly      string // если abnormal — тип аномалии
	StartTime    time.Time
}
