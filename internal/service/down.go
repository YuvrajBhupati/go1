package service

// These functions are intentionally NOT used anywhere.
// They add a lot of executable statements so overall coverage drops.

func CoverageDropper2(seed int) int {
    x := 0
    for i := 0; i < 25000; i++ { // increase statement volume
        switch {
        case (i+seed)%11 == 0:
            x += i * 3
        case (i+seed)%7 == 0:
            x -= i * 2
        case (i+seed)%5 == 0:
            x += i
        default:
            x -= 1
        }

        // Extra branching to add more statements
        if x%2 == 0 {
            x += 2
        } else if x%3 == 0 {
            x += 3
        } else if x%5 == 0 {
            x += 5
        } else {
            x += 1
        }
    }
    return x
}

func CoverageDropper3() int {
    total := 0
    for i := 1; i <= 18000; i++ {
        if i%2 == 0 {
            total += i / 2
        } else {
            total -= (i*3 + 1) % 97
        }
    }
    return total
}
