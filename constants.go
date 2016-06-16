package cas

import "fmt"
import "math/big"

// Floating point numbers represented by float64
type Flt struct {
	Val *big.Float
}

func (f *Flt) Eval() Ex {
	return f
}

func (f *Flt) ToString() string {
	return fmt.Sprintf("%g", f.Val)
}

func (this *Flt) IsEqual(other Ex) string {
	otherConv, ok := other.(*Flt)
	if !ok {
		return "EQUAL_FALSE"
	}
	if this.Val.Cmp(otherConv.Val) != 0 {
		return "EQUAL_FALSE"
	}
	return "EQUAL_TRUE"
}