package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	fromX := flag.Float64("fromX", 0, "Start X coord")
	fromY := flag.Float64("fromY", 0, "Start Y coord")
	toX := flag.Float64("toX", 10, "End X coord")
	toY := flag.Float64("toY", 10, "End Y coord")
	points := flag.Int("points", 100, "Number of points in route")
	routeType := flag.String("type", "normal", "Route type: normal or abnormal")
	anom := flag.String("anom", "", "Anomaly type if abnormal (zigzag, lost_signal, wrong_heading, depth_spike)")
	outDir := flag.String("out", "./data", "Output directory")
	count := flag.Int("count", 1, "Number of routes to generate")

	flag.Parse()

	for i := 0; i < *count; i++ {
		cfg := RouteConfig{
			FromX:     *fromX,
			FromY:     *fromY,
			ToX:       *toX,
			ToY:       *toY,
			Points:    *points,
			RouteType: *routeType,
			Anomaly:   *anom,
			StartTime: time.Now().UTC(),
		}

		filePath, err := GenerateRoute(cfg, *outDir)
		if err != nil {
			panic(err)
		}

		fmt.Printf("[%d/%d] File generated: %s\n", i+1, *count, filePath)
	}
}
