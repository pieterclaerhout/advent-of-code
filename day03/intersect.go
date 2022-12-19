package day03

func intersect(a []string, b []string) []string {
	m := make(map[string]bool)
	for _, item := range a {
		m[item] = true
	}

	common := []string{}
	for _, item := range b {
		if _, exists := m[item]; exists {
			common = append(common, item)
		}
	}

	return common
}
