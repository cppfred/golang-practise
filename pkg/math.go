package pkg

// MinInt64 return minimum int64 value
func MinInt64(x int64, y int64) int64 { // Min int function
	if x > y {
		return y
	}
	return x
}
