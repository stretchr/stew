package slice

// CommonStrings gets a new []string that contains strings that are
// present in both specified slices.
func CommonStrings(s1, s2 []string) []string {

	var common []string
	for _, v := range s1 {
		if ContainsString(s2, v) {
			common = append(common, v)
		}
	}

	return common

}
