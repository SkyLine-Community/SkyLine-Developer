package SkyLine_Backend

import (
	"fmt"
	"os"
)

func ReverseArrayForFileTraceback(a []string) {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func ReverseObjectArrayWithinCSCF(a []Object) {
	for k := len(a)/2 - 1; k >= 0; k-- {
		opp := len(a) - 1 - k
		a[k], a[opp] = a[opp], a[k]
	}
}

func CheckArgumentLength(ProperLength int, name string, Args []Object) {
	if len(Args) != ProperLength {
		fmt.Printf("Function %s must have %d positional arguments, function call got %d", name, ProperLength, len(Args))
		os.Exit(1)
	}
}
