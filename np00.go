package np00

import (
	"fmt"
	"math"
)

type NNetwork map[string]NParray
type Array []float64
type NParray []Array

const ROWPART = 0
const COLPART = 1

func IdentityFunction(npa NParray) NParray {
	return npa
}

func sigmoid(f float64) float64 {
	return 1 / (1 + math.Exp(-f))
}

func Sigmoid(npa NParray) NParray {
	npr := MakeNParray(1, npa.Shape()[1])
	for n, e := range npa[0] {
		npr[0][n] = sigmoid(e)
	}
	return npr
}

func MaxFloatInSlice(fls []float64) (m float64) {

	m = fls[len(fls)-1]
	for _, e := range fls {
		if m <= e {
			m = e
		}
	}
	return m
}

func Sum(fls []float64) (s float64) {
	return SumOfSlice(fls)
}

func SumOfSlice(fls []float64) (s float64) {
	for _, e := range fls {
		s += e
	}
	return s
}

// Is 'long' better?
func (n NParray) Shape() [2]int {
	row := len(n)
	col := len(n[row-1])
	return [2]int{row, col}
}

func sumExpC(fls []float64) (s float64) {
	c := MaxFloatInSlice(fls)
	for _, e := range fls {
		s += math.Exp(e - c)
	}
	return s
}

func SoftMax(fls []float64) (sm []float64) {
	c := MaxFloatInSlice(fls)
	sum_exp_c := sumExpC(fls)
	sm = make([]float64, len(fls))

	for i, v := range fls {
		sm[i] = math.Exp(v-c) / sum_exp_c
	}
	return sm
}

func MakeNParray(row, col int) NParray {
	npa := make([]Array, row)
	for z := range npa {
		Array := make([]float64, col)
		npa[z] = Array
	}
	return npa
}

func (m NParray) ColsToArray(colid int) (fa []float64) {
	clen := m.Shape()[ROWPART]
	fa = make([]float64, clen)
	for r := range m {
		fa[r] = m[r][colid]
	}
	return
}

func (m NParray) RowsToArray(rowid int) (fa []float64) {
	rlen := m.Shape()[COLPART]
	fa = make([]float64, rlen)
	/* not required
	for rn:=0; rn<rlen;rn++ {
		fa[rn] = m[rowid][rn]
	}
	*/
	fa = m[rowid]
	return
}

func (a Array) Add(b Array) (f float64) {

	if len(a) == len(b) {
		for idx := range a {
			f += a[idx] * b[idx]
		}
	}
	return f
}

func (n NParray) String() string {

	var str string = ""
	str += fmt.Sprintf("(%v){\n", n.Shape())

	for r := range n {
		str += fmt.Sprintf("[ ")

		for k := range n[r] {
			str += fmt.Sprintf("%v ", n[r][k])
		}

		str += fmt.Sprintf("]\n")
	}
	str += fmt.Sprintf("}")
	return str
}

func Add(n NParray, m NParray) NParray {

	npa := make([]Array, n.Shape()[0])
	for z := range npa {
		Array := make([]float64, n.Shape()[1])
		npa[z] = Array
	}

	if n.Shape()[0] == m.Shape()[0] && n.Shape()[1] == m.Shape()[1] {
		for r := range n {
			for k := range n[r] {
				npa[r][k] = n[r][k] + m[r][k]
			}
		}
	} else {
		panic("Shape check error")
	}
	return npa
}

func (n NParray) Multi(f float64) NParray {

	npa := make([]Array, n.Shape()[0])
	for z := range npa {
		Array := make([]float64, n.Shape()[1])
		npa[z] = Array
	}

	for r := range n {
		for k := range n[r] {
			npa[r][k] = n[r][k] * f
		}
	}
	return npa
}

/** NParray Dot func
(1xN) [a1 a2 a3 a4.. aN]  .
(Nx1)   (c1
         c2
		 c3
		 c4
		 .
		 .
		 cN
		 )
===> a1*c1 + c2*c2 + a3*c3 + a4*c4 .... aN*cN (1x1)

(kxN) [a1 a2 a3 ... aN]
      [b1 b2 b3 ... bN]
            ...
	  [k1 k2 k3 ... kN]
	  .
(Nxm) (
	   [f1 g1 ... m1]
       [f2 g2 ... m2]
	   [f3 g3 ... m3]
       [...   ... m]
       [fN gN ... mN]
       )
	   A DOT B　requires:
	   B.col == A.row = N
	   Results :
	   new matrix has row k x , col m

img: https://s3.amazonaws.com/nkvd/pub/matrixDot.png
*/

func Dot(n NParray, m NParray) NParray {

	if n.Shape()[COLPART] != m.Shape()[ROWPART] {
		panic("error row x col")
	}

	npa := make([]Array, n.Shape()[ROWPART])
	for z := range npa {
		Array := make([]float64, m.Shape()[COLPART])
		npa[z] = Array
	}
	for ncol := range n {
		for mrow := range m[0] {
			npa[ncol][mrow] = n[ncol].Add(m.ColsToArray(mrow))
		}
	}
	return npa
}

func ReLU(x float64) float64 {
	return math.Max(0, x)
}
