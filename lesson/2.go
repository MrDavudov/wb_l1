package lesson

func Two(s, f int) (res int) {
	for i:=s;i<f;i++ {
		if i%2 == 0 {
			res = res + i * i
		}
	}
	return res
}