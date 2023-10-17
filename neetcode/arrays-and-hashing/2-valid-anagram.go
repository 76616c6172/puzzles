package neetcode

func IsAnagram(s, t string) bool {
	m1 := make(map[int32]int)
	m2 := make(map[int32]int)

	for _, letter := range s {
		m1[letter]++
	}
	for _, letter := range t {
		m2[letter]++
	}

	if len(m1) != len(m2) {
		return false
	}

	for k := range m1 {
		if _, exists := m2[k]; !exists {
			return false
		}

		if m1[k] != m2[k] {
			return false
		}
	}

	return true
}
