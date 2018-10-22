// Package dual provides computation for dual numbers.
package dual

type Dual struct {
	real float64
	eps  float64
}

func New(real, eps float64) Dual {
	return Dual{real, eps}
}

func FromFloat(real float64) Dual {
	return Dual{real, 0.0}
}

func Real(d Dual) float64 {
	return d.real
}

func Epsilon(d Dual) float64 {
	return d.epsilon
}

func Add(d1, d2 Dual) Dual {
	return Dual{d1.real + d2.real, d1.eps + d2.eps}
}

func Sub(d1, d2 Dual) Dual {
	return Dual{d1.real - d2.real, d1.eps - d2.eps}
}

func Mul(d1, d2 Dual) Dual {
	return Dual{d1.real * d2.real, d1.real*d2.eps + d1.eps*d2.real}
}

func Div(d1, d2 Dual) Dual {
	return Dual{
		real: d1.real / d2.real,
		eps:  (d1.eps*d2.real - d1.real*d2.eps) / (d2.real * d2.real),
	}
}
