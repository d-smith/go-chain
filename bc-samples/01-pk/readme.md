# Bitcoin private keys

A private key can be almost any random 256 bit number

* Must be not be higher than the max value that is the number of points on the elliptic curve
* Bitcoin eliptic curve described by Secp256k1
* Max value for the random key is n-1, where n is 1.1578*10**77
* Make sure you use a crptographically secure psuedo random number generator with a seed from a source of sufficient entropy
* Bitcoin - edcsa with secp256k1 curve order and sha256 hash function.


From the public derived from the private key we can generate bit coin addresses - https://learnmeabitcoin.com/technical/address

* https://github.com/bitcoinbook/bitcoinbook/blob/develop/ch04.asciidoc

