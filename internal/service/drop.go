package service

func Drop1() int { x := 0; for i := 0; i < 3000; i++ { if i%2==0 { x+=i } else { x-=i } }; return x }
func Drop2() int { x := 0; for i := 0; i < 3000; i++ { if i%3==0 { x+=i*2 } else { x+=i } }; return x }
func Drop3() int { x := 0; for i := 0; i < 3000; i++ { if i%5==0 { x+=i*3 } else { x+=i } }; return x }
