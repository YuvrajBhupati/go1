package dummy

func A(n int) int { r := 0; for i := 0; i < n; i++ { r += i }; return r }
func B(n int) int { r := 1; for i := 1; i < n; i++ { r *= i }; return r }
func C(n int) int { r := 0; for i := 0; i < n; i++ { if i%2==0 { r++ } else { r-- } }; return r }
