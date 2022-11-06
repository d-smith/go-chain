package seed

import (
	"testing"
	"fmt"
)

func TestSeedGenerated(t *testing.T) {
	seed := GetMnemonicPhrase()
	fmt.Println(seed)
}