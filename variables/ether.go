package variables

import "math/big"

var c = big.NewInt(1000000000000)

func Ether(value float32) *big.Int {
	r := int64(1000000 * value)
	return big.NewInt(0).Mul(c, big.NewInt(r))
}
