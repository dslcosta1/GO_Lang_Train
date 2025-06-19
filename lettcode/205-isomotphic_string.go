func isIsomorphic(s string, t string) bool {
	var mapS2T = make(map[byte]byte)
	var mapT2S = make(map[byte]byte)

	n := len(s)

	for i := 0; i < n; i++ {
		val, ok := mapS2T[s[i]]

		if ok {
			if val != t[i] {
				return false
			}
		}

		val, ok = mapT2S[t[i]]

		if ok {
			if val != s[i] {
				return false
			}
		}

		mapS2T[s[i]] = t[i]
		mapT2S[t[i]] = s[i]
	}

	return true
}