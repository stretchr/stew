package slice

func MinusStrings(s, minus []string) []string {

	a := []string{}

	l := len(s)
	for i := 0; i < l; i++ {
		if !ContainsString(minus, s[i]) {
			a = append(a, s[i])
		}
	}

	return a
}
