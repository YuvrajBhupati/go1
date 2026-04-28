package coverage_dropper

func UncoveredA() int {
    return 10
}

func UncoveredB(x int) int {
    if x > 5 {
        return x * 2
    }
    return x - 3
}

func UncoveredC() string {
    return "not tested"
}
