package service

// Dummy function with many statements.
// If no test calls this, coverage will drop.
func CoverageDropper(x int) int {
    total := 0
    for i := 0; i < 200; i++ { // lots of statements
        if (x+i)%2 == 0 {
            total += i
        } else {
            total -= i
        }

        if total%3 == 0 {
            total += 1
        } else if total%5 == 0 {
            total += 2
        } else {
            total += 3
        }
    }
    return total
}
