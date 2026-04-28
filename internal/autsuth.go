package service

func BigUncoveredLogic() int {
    sum := 0
    for i := 0; i < 5000; i++ { // increase loops/branches
        if i%7 == 0 {
            sum += i * 2
        } else if i%5 == 0 {
            sum += i * 3
        } else {
            sum += i
        }
    }
    return sum
}
