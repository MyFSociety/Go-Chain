// proof of work implementation for the blockchain

package utility

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"math/big"
	"math/rand"
)

// proof of work struct

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

// for demonstration purposes, we will set the target to 24 bits
const targetBits = 24

// max nounce value
const maxNounce = 100000000000

// creating a new proof of work
func NewProofOfWork(b *Block) *ProofOfWork {
	rand_int := rand.Int63n(1000000000000000000)
	target := big.NewInt(rand_int)
	// left shift 256 - targetBits
	target.Lsh(target, uint(256-targetBits))
	return &ProofOfWork{b, target}
}

func (pow *ProofOfWork) prepareData(nounce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PrevBlockHash,
			pow.block.Data,
			IntToByte(pow.block.Timestamp),
			IntToByte(int64(targetBits)),
			IntToByte(int64(nounce)),
		},
		[]byte{},
	)
	return data
}

// convert int to byte array
func IntToByte(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		panic(err)
	}
	return buff.Bytes()
}

// run the proof of work algorithm
func (pow *ProofOfWork) Run() (int, []byte, error) {
	var intHash big.Int
	var hash [32]byte
	nounce := 0
	for nounce < maxNounce {
		data := pow.prepareData(nounce)
		hash = sha256.Sum256(data)
		intHash.SetBytes(hash[:])
		if intHash.Cmp(pow.target) == -1 {
			break
		} else {
			nounce++
		}
	}
	return nounce, hash[:], nil
}

// validate the proof of work
// func (pow *ProofOfWork) Validate() bool {
// 	var intHash big.Int
// 	data := pow.prepareData(pow.block.Nounce)
// 	hash := sha256.Sum256(data)
// 	intHash.SetBytes(hash[:])
// 	return intHash.Cmp(pow.target) == -1
// }
