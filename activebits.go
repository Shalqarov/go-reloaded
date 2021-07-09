package student

func ActiveBits(n int) uint {
	bits := uint(0)
	for n != 0 {
		if n%2 == 1 {
			bits++
		}
		n = n >> 1
	}
	return bits
}
