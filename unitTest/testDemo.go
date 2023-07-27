package unittest

func HelloTom() string {
	return "Jack"
}

// 判断是否及格
func RatingScore(score uint8) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "E"
	}
}
