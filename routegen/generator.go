package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

func GenerateRoute(cfg RouteConfig, outDir string) (string, error) {
	dx := (cfg.ToX - cfg.FromX) / float64(cfg.Points)
	dy := (cfg.ToY - cfg.FromY) / float64(cfg.Points)

	points := make([]Point, cfg.Points)

	for i := 0; i < cfg.Points; i++ {
		timestamp := cfg.StartTime.Add(time.Duration(i) * time.Minute).Format(time.RFC3339)

		x := cfg.FromX + dx*float64(i)
		y := cfg.FromY + dy*float64(i)

		// шум в нормальных маршрутах
		if cfg.RouteType == "normal" {
			x += rand.Float64()*2 - 1
			y += rand.Float64()*2 - 1
		}

		// скорость
		speed := 10 + rand.Float64()*10

		// heading (по направлению движения)
		heading := int(math.Atan2(dy, dx)*180/math.Pi) % 360
		if cfg.RouteType == "normal" {
			heading += rand.Intn(11) - 5
		}

		// глубина
		depth := 100 + rand.Float64()*200

		// сигнал
		signal := 80 + rand.Intn(21)

		p := Point{
			Timestamp:      timestamp,
			X:              x,
			Y:              y,
			SpeedKnots:     speed,
			HeadingDeg:     heading,
			DepthM:         depth,
			SignalStrength: signal,
		}

		// если аномальный маршрут — применяем аномалию
		if cfg.RouteType == "abnormal" && cfg.Anomaly != "" {
			ApplyAnomaly(&p, i, cfg)
		}

		points[i] = p
	}

	// имя файла
	timestampNow := time.Now().Format("15-04-05.000") // часы-минуты-секунды-миллисекунды
	fileName := fmt.Sprintf("%s_%s.csv", cfg.RouteType, timestampNow)
	if cfg.RouteType == "abnormal" && cfg.Anomaly != "" {
		fileName = fmt.Sprintf("%s_%s_%s.csv", cfg.RouteType, cfg.Anomaly, timestampNow)
	}

	// чтобы имена не совпадали при быстром цикле
	fileName = fmt.Sprintf("%s_%03d.csv", fileName[:len(fileName)-4], rand.Intn(1000))
	filePath := filepath.Join(outDir, fileName)

	// сохраняем в CSV
	if err := os.MkdirAll(outDir, os.ModePerm); err != nil {
		return "", err
	}

	file, err := os.Create(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// заголовки
	writer.Write([]string{"timestamp", "x_coord", "y_coord", "speed_knots", "heading_deg", "depth_m", "signal_strength"})

	for _, p := range points {
		writer.Write([]string{
			p.Timestamp,
			fmt.Sprintf("%.2f", p.X),
			fmt.Sprintf("%.2f", p.Y),
			fmt.Sprintf("%.2f", p.SpeedKnots),
			fmt.Sprintf("%d", p.HeadingDeg),
			fmt.Sprintf("%.2f", p.DepthM),
			fmt.Sprintf("%d", p.SignalStrength),
		})
	}

	return filePath, nil
}
