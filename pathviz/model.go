package main

import (
	"fmt"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type MsgTick struct{}

type Model struct {
	Points        []DataPoint
	Index         int
	Width, Height int
	MinX, MaxX    float64
	MinY, MaxY    float64
}

func (m Model) Init() tea.Cmd {
	return tick()
}

func tick() tea.Cmd {
	return tea.Tick(time.Millisecond*300, func(time.Time) tea.Msg {
		return MsgTick{}
	})
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case MsgTick:
		if m.Index < len(m.Points)-1 {
			m.Index++
			return m, tick()
		} else {
			return m, tea.Quit
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height - 5
		return m, nil
	}
	return m, nil
}

func (m Model) View() string {
	if len(m.Points) == 0 {
		return "No points to display"
	}

	p := m.Points[m.Index]
	xNorm := Normalize(p.X, m.MinX, m.MaxX, m.Width)
	yNorm := Normalize(p.Y, m.MinY, m.MaxY, m.Height)

	// создаём пустую сетку
	grid := make([][]rune, m.Height)
	for i := range grid {
		grid[i] = make([]rune, m.Width)
		for j := range grid[i] {
			grid[i][j] = '.' // сетка
		}
	}

	// выбираем цвет текущей точки по скорости
	var colorFunc func(string) string
	if p.Speed < 10 {
		colorFunc = func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("#00AA00")).Render(s) // мягкий зеленый
		}
	} else if p.Speed < 20 {
		colorFunc = func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("#AAAA00")).Render(s) // мягкий желтый
		}
	} else {
		colorFunc = func(s string) string {
			return lipgloss.NewStyle().Foreground(lipgloss.Color("#AA0000")).Render(s) // мягкий красный
		}
	}

	// строим строку для отображения
	view := ""
	for i := m.Height - 1; i >= 0; i-- {
		for j := 0; j < m.Width; j++ {
			c := string(grid[i][j])
			if i == yNorm && j == xNorm {
				// текущий корабль яркий
				c = colorFunc("●")
			} else {
				// все остальные точки сетки приглушённые
				c = lipgloss.NewStyle().Foreground(lipgloss.Color("#444444")).Render(c)
			}
			view += c
		}
		view += "\n"
	}

	// текстовая панель с характеристиками
	info := fmt.Sprintf(
		"\nTime: %s | Speed: %.2f kn | Heading: %.1f° | Depth: %.2f m | Signal: %.0f",
		p.Timestamp, p.Speed, p.Heading, p.Depth, p.SignalStrength,
	)

	return view + info
}
