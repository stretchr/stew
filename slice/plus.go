package slice

func PlusStrings(s, plus []string) []string {

	var l int

	a := []string{}

	l = len(s)
	for i := 0; i < l; i++ {
		a = append(a, s[i])
	}

	l = len(plus)
	for i := 0; i < l; i++ {
		a = append(a, plus[i])
	}

	return a

}
