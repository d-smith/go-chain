package seed

import (
	"fmt"
	"math/big"
	"testing"
)

func TestSeedGenerated(t *testing.T) {
	seed := GetMnemonicPhrase()
	fmt.Println(seed)
	entropy := EntropyFromMnemonic(seed)
	dataBigInt := new(big.Int).SetBytes(entropy)
	fmt.Println("Entropy:         ", dataBigInt)
}
