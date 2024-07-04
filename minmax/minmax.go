package gag

func Min(x int, y int) int {
	if x > y {
		return y
	}
	return x
}

func Max(x int, y int) int {
	if x < y {
		return y
	}
	return x
}
