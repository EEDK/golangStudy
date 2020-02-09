package calc

func fibonacci() func() int {
	fn := 0
	fn1 := 0
	return func() int {
		fn2 := fn1 + fn
		fn = fn1
		fn1 = fn2
		if fn1 == 0 {fn++}
		return fn2
	}
}
