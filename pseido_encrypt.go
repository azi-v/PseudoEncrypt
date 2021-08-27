package pseudoencrypt

func PseudoEncrypt(value int64) int64 {
	var l1, l2, r1, r2 int64 = 0, 0, 0, 0
	l1 = (value >> l2) & (4096 - 1)
	r2 = value & (4096 - 1)
	for i := 0; i < 3; i++ {
		l2 = r1
		r2 = l1 ^ ((((1366*r1 + 150889) % 714025) / 714025.0) * (4096 - 1))
		l1 = l2
		r1 = r2
	}
	return (l1 << l2) + r1
}

func BoundedPseudoEncrypt(value, max,min int64) (res int64) {
	for {
		if res = PseudoEncrypt(value);res<=max&&res>=min{
			return res
		}
	}
}
