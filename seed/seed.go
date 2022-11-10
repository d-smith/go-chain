package seed

import (
	"crypto/rand" 
	"crypto/sha256"
    "encoding/hex"
    "fmt"
	"math/big"
	"io/ioutil"
	"strings"
)

var (
    wordlist []string
    wordMap = map[string]int{}
)


func init() {
    readWordList()
}

func readWordList() {
    file, _ := ioutil.ReadFile("wordlist.txt")
    wordlist = strings.Split(string(file), "\n")

    for i, v := range wordlist {
		wordMap[v] = i
	}
}



func GetMnemonicPhrase() string {

	// Three steps to creating a mnenomic sentance
	// 1. Generate entropy
	// 2. Convert to mnemonic
	// 3. Mnemonic to see

	// Entropy --------------------


	bytes := make([]byte, 32) // [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
    rand.Read(bytes)
    fmt.Println("Entropy:", hex.EncodeToString(bytes))
    fmt.Println("Entropy:", bytes)
    fmt.Printf("Entropy: %08b\n", bytes)
    fmt.Println()

	// We now create a checksum to help detect errors when deriving the 
	// entropy from the seed phrase. CHecksum is formed by taking the sha256
	// of the entropy, and inlcuding 1 bit of that hash for every 32 bits of
	// entropy. Note 11 bits are mapped to one of the keywords, 2^11 is
	// the number of seed words available in the standard, 1 bit for every
	// 32 gives us multiples of 11...

	h := sha256.New()                        // hash function
    h.Write(bytes)                           // write bytes to hash function
    checksum := h.Sum(nil)                   // get result as a byte slice
    fmt.Printf("Checksum: %08b\n", checksum) // 10000101

	// Get specific number of bits from the checksum
    size := len(bytes) * 8 / 32       // 1 bit for every 32 bits of entropy (1 byte = 8 bits)
    fmt.Println("Bits Wanted:", size) // how many bits from the checksum we want
    fmt.Println()

	// The entropy bytes are converted into a big integer
	// to allow arithmetic to be used for inserting the 
	// checksum bits - byte slices don't allow for appending
	// single gits
	dataBigInt := new(big.Int).SetBytes(bytes)
    fmt.Println("Entropy:         ", dataBigInt)

	// This loop lifted verbatim from the learnmeabitcode site
	// Run through the number of bits you want from the checksum, manually adding each bit to the entropy (through arithmetic)
    for i := uint8(0); i < uint8(size); i++ {
        // Add a zero bit to the end for every bit of checksum we add
        //
        //          --->
        //          01001101
        // |entropy|0|
        //
        dataBigInt.Mul(dataBigInt, big.NewInt(2)) // multiplying an integer by two is like adding a 0 bit to the end

        // Use bitwise AND mask to check if each bit of the checksum is set
        //
        // checksum[0] = 01001101
        //           AND 10000000 = 0
        //           AND  1000000 = 1000000
        //           AND   100000 = 0
        //           AND    10000 = 0
        //
        mask := 1 << (7 - i)
        set := uint8(checksum[0]) & uint8(mask) // e.g. 100100100 AND 10000000 = 10000000

        if set > 0 {
            // If the bit is set, change the last zero bit to a 1 bit
            //          10001101
            // |entropy|1|
            //
            dataBigInt.Or(dataBigInt, big.NewInt(1)) // Use bitwise OR to toggle last bit (basically adds 1 to the integer)
        }
    }

    fmt.Println("Entropy+Checksum:", dataBigInt)
    fmt.Println()


	// 3. Map to seed words from the standard
	

	// How many 11 bit pieces are there?
	pieces := ((len(bytes) * 8) + size) / 11

    // Create an array of strings to hold words
    words := make([]string, pieces)

	// Loop through every 11 bits of entropy+checksum and convert to corresponding word from wordlist
    for i := pieces - 1; i >= 0; i-- {

        // Use bit mask (bitwise AND) to split integer in to 11-bit pieces
        //
        //            right to left          big.NewInt(2047) = bit mask
        //          <----------------          <--------->
        // 11111111111|11111111111|11111111111|11111111111
        //
        word := big.NewInt(0)                  // hold result of 11 bit mask
        word.And(dataBigInt, big.NewInt(2047)) // bit mask last 11 bits (2047 = 0b11111111111)

        // Add corresponding word to array
        //
        // 11100111000 = 1848 = train
        //
        words[i] = wordlist[word.Int64()] // insert word from wordlist in to array (need to convert big.Int to int64)

        // Remove those 11 bits from end of big integer by bit shifting
        //
        // 11111111111|11111111111|11111111111|11111111111
        //                                    /            - dividing is the same as bit shifting
        //                                    100000000000 = big.NewInt(2048)
        // 11111111111|11111111111|11111111111|
        //
        dataBigInt.Div(dataBigInt, big.NewInt(2048)) // dividing by 2048 is the same as bit shifting 11 bits

    }

	mnemonic := strings.Join(words, " ")
    fmt.Println("Mnemonic:", mnemonic)
    fmt.Println()

	return mnemonic;
}

func entropyFromMnemonic(phrase string) []byte {
    return make([]byte, 32);
}