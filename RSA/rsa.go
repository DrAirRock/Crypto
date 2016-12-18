package RSA

import (
	"crypto/rand"
	"math/big"

	"github.com/lannonbr/Crypto/Utils"
)

// Generate a single prime with a specific bit size
func genPrime(bitsize int) (*big.Int, error) {
	return rand.Prime(rand.Reader, bitsize)
}

// Generate keys needed for RSA in the form of two arrays, [N, e] and [N, d]
func genRSAKeys(bitSize int) ([]*big.Int, []*big.Int) {
	p, _ := genPrime(bitSize)
	q, _ := genPrime(bitSize)

	N := Utils.Mult(p, q)
	phi := Utils.Mult(Utils.Sub(p, big.NewInt(1)), Utils.Sub(q, big.NewInt(1)))

	e := big.NewInt(17) // It's prime, good enough

	d := Utils.ModInv(e, phi)

	return []*big.Int{N, e}, []*big.Int{N, d}
}

// Encrypt a string with RSA to an array where each string in the output is a single encrypted character
func encryptRSA(plaintext string, N, e *big.Int) []string {
	plaintextBytes := []byte(plaintext)

	var cipherArr []string

	for _, character := range plaintextBytes {
		encryptedCharacter := big.NewInt(int64(character))
		encryptedCharacter = encryptedCharacter.Exp(encryptedCharacter, e, N)
		cipherArr = append(cipherArr, encryptedCharacter.String())
	}

	return cipherArr
}

// Decrypt the encrypted cipher array into a plaintext string as output
func decryptRSA(cipher []string, N, d *big.Int) string {
	out := ""

	for _, line := range cipher {
		decryptedCharacter, _ := big.NewInt(0).SetString(line, 10)
		decryptedCharacter = Utils.Exp(decryptedCharacter, d, N)

		out += string(decryptedCharacter.Int64())
	}

	return out
}
