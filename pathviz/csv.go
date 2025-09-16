package main

import (
    "encoding/csv"
    "os"
    "strconv"
)

type DataPoint struct {
    Timestamp      string
    X, Y           float64
    Speed          float64
    Heading        float64
    Depth          float64
    SignalStrength float64
}

func LoadCSV(path string) ([]DataPoint, error) {
    f, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer f.Close()

    r := csv.NewReader(f)
    records, err := r.ReadAll()
    if err != nil {
        return nil, err
    }

    var points []DataPoint
    for i, rec := range records {
        if i == 0 {
            continue // skip header
        }
        x, _ := strconv.ParseFloat(rec[1], 64)
        y, _ := strconv.ParseFloat(rec[2], 64)
        speed, _ := strconv.ParseFloat(rec[3], 64)
        heading, _ := strconv.ParseFloat(rec[4], 64)
        depth, _ := strconv.ParseFloat(rec[5], 64)
        signal, _ := strconv.ParseFloat(rec[6], 64)
        points = append(points, DataPoint{
            Timestamp:      rec[0],
            X:              x,
            Y:              y,
            Speed:          speed,
            Heading:        heading,
            Depth:          depth,
            SignalStrength: signal,
        })
    }
    return points, nil
}
