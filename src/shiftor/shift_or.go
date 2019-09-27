package shiftor

var mask = [128]int{0: ^0}

func Search(txt string, pat string) (occ []int) {
	// make it faster by reusing mask
	if mask[0] == ^0 {
		mask = buildMask(pat)
	}

	m := len(pat)
	n := len(txt)
	s := 1 << uint(m)

	for i := 0; i < n; i++ {
		s <<= 1
		s |= mask[txt[i]]

		if (s >> uint(m - 1)) & 1 == 0 && i >= m {
			occ = append(occ, i - m + 1)
		}
	}

	return
}

func buildMask(pat string) (mask [128]int) {
	m := len(pat)
	stamp := 1 << uint(m)


	init :=  (1 << uint(m)) - 1
	for i := 0; i < 128; i++ {
		mask[i] = init
	}

	// a is a rune which is a 32bit int
	for _, a := range pat {
		checkRune(a)
		mask[a] &= stamp
		stamp = (stamp << 1) | 1
	}

	return
}

// TO-DO remove this check
func checkRune(r rune) {
	if r > 127 {
		panic("Unsupported rune size")
	}
}