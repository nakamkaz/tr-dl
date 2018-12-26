package main

/**
p.69 - ゼロから作るDeep Learning 3.5.2 ソフトマックス関数実装上の注意

そのまま代入すると指数関数の結果がオーバーフローしてしまうので、
あらかじめexpに入れる前に入力の最大値を引いておく。
というのをgolangで
*/

import (
	"fmt"
	"math"
)

func MaxFloatInSlice(fls []float64) (m float64) {

	m = fls[len(fls)-1]
	for _, e := range fls {
		if m <= e {
			m = e
		}
	}
	return m
}

func Sum(fls []float64) float64 {
	var s float64 = 0
	for _, e := range fls {
		s += e
	}
	return s
}

func SumExpC(fls []float64) float64 {
	var s float64 = 0
	c := MaxFloatInSlice(fls)
	for _, e := range fls {
		s += math.Exp(e - c)
	}
	return s
}

func SoftMax(fls []float64) (sm []float64) {
	c := MaxFloatInSlice(fls)
	sum_exp_c := SumExpC(fls)
	sm = make([]float64, len(fls))

	for i, v := range fls {
		sm[i] = math.Exp(v-c) / sum_exp_c
	}
	return sm
}

func main() {
	a := []float64{1010, 1000, 990}
	fmt.Println("Max: ", MaxFloatInSlice(a))
	fmt.Println("Sum of a: ", Sum(a))
	fmt.Printf("SoftMax: %v\n", SoftMax(a))
	fmt.Printf("Sum of SoftMax %v\n", Sum(SoftMax(a)))
}

/**
http://goiduuid.appspot.com/p/6mkmFXck0k

Max:  1010
Sum of a:  3000
SoftMax: [0.999954600070331 4.539786860886666e-05 2.061060046209062e-09]
Sum of SoftMax 1

*/
