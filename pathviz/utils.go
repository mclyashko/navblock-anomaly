package main

func Normalize(val, min, max float64, scale int) int {
    if max == min {
        return 0
    }
    return int((val - min) / (max - min) * float64(scale-1))
}
