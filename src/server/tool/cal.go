package tool

import "math/rand"

func C_M_N(m int32, n int32) []int32 {
	res := make([]int32, 0)
	origin := make([]int32, 0)

	var i int32
	var j int32
	for i = 0; i < m; i++ {
		origin = append(origin, i)
	}

	for j = 0; i < n; j++ {
		index := C_M_1(int32(len(origin)))
		res = append(res, origin[index])
		if index == int32(len(origin))-1 {
			origin = origin[0:index]
		} else {
			origin = append(origin[0:index], origin[index+1:]...)
		}
	}
	return res
}

func C_M_1(m int32) int32 {
	res := rand.Int31n(m)
	return res
}
