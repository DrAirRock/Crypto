package Utils

import "math/big"

func Add(a, b *big.Int) *big.Int {
	return big.NewInt(0).Add(a, b)
}

func Mult(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mul(a, b)
}

func Sub(a, b *big.Int) *big.Int {
	return big.NewInt(0).Sub(a, b)
}

func Div(a, b *big.Int) *big.Int {
	return big.NewInt(0).Div(a, b)
}

func Mod(a, b *big.Int) *big.Int {
	return big.NewInt(0).Mod(a, b)
}

func Exp(x, y, m *big.Int) *big.Int {
	return big.NewInt(0).Exp(x, y, m)
}

func Eq(a, b *big.Int) bool {
	tmp := a.String()

	c, _ := big.NewInt(0).SetString(tmp, 10)

	return c.Cmp(b) == 0
}

func ModInv(a, N *big.Int) *big.Int {
	_, _, d := EGCD(N, a)

	if d.Cmp(big.NewInt(0)) == -1 {
		d = Sub(N, Abs(d))
	}

	return d
}

func Abs(a *big.Int) *big.Int {
	return big.NewInt(0).Abs(a)
}

func EGCD(a, b *big.Int) (*big.Int, *big.Int, *big.Int) {
	q := Div(a, b)
	r := Mod(a, b)

	x1 := big.NewInt(1)
	y1 := big.NewInt(0)
	x2 := big.NewInt(0)
	y2 := big.NewInt(1)

	for {
		if Eq(r, big.NewInt(0)) {
			break
		}
		a = b
		b = r

		oldQ := q

		q = Div(a, b)
		r = Mod(a, b)

		oldX1 := x1
		oldY1 := y1

		x1 = x2
		y1 = y2

		x2 = Sub(oldX1, Mult(oldQ, x2))
		y2 = Sub(oldY1, Mult(oldQ, y2))
	}
	return b, x2, y2
}
