package main

import (
	"math/rand"
)

// ApplyAnomaly изменяет точку в зависимости от выбранной аномалии
func ApplyAnomaly(p *Point, i int, cfg RouteConfig) {
	switch cfg.Anomaly {
	case "zigzag":
		if i%10 < 5 {
			p.X += rand.Float64()*10 - 5
			p.Y += rand.Float64()*10 - 5
		}
	case "wrong_heading":
		// курс меняем на 90°
		p.HeadingDeg = (p.HeadingDeg + 90) % 360
	case "lost_signal":
		// сигнал пропадает после половины пути
		if i > cfg.Points/2 {
			p.SignalStrength = 0
		}
	case "depth_spike":
		// глубина скачет каждые 20 точек
		if i%20 == 0 {
			p.DepthM = 400 + rand.Float64()*100
		}
	}
}
