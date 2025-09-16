package main

import (
    "flag"
    "log"

    tea "github.com/charmbracelet/bubbletea"
)

func main() {
    filePath := flag.String("file", "", "Path to CSV file")
    flag.Parse()

    if *filePath == "" {
        log.Fatal("Please provide -file path")
    }

    points, err := LoadCSV(*filePath)
    if err != nil {
        log.Fatal(err)
    }

    m := Model{
        Points: points,
        Index:  0,
        MinX:   points[0].X,
        MaxX:   points[0].X,
        MinY:   points[0].Y,
        MaxY:   points[0].Y,
    }

    for _, p := range points {
        if p.X < m.MinX {
            m.MinX = p.X
        }
        if p.X > m.MaxX {
            m.MaxX = p.X
        }
        if p.Y < m.MinY {
            m.MinY = p.Y
        }
        if p.Y > m.MaxY {
            m.MaxY = p.Y
        }
    }

    program := tea.NewProgram(m, tea.WithAltScreen())
    if err := program.Start(); err != nil {
        log.Fatal(err)
    }
}
