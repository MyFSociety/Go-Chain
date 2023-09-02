// entry point for the application
package main

import (
	"fmt"
	"harry/blockchain/utility"
)

func main() {
	chain := utility.NewBlockchain()
	chain.AddBlock("First Block after Genesis")
	chain.AddBlock("Second Block after Genesis")
	chain.AddBlock("Third Block after Genesis")
	chain.AddBlock("Fourth Block after Genesis")
	chain.AddBlock("Fifth Block after Genesis")

	for block_index, block := range chain.BlockArray {
		fmt.Printf("Block Number: %d\n", block_index)
		fmt.Printf("Previous Block Hash: %x\n", block.PrevBlockHash)
		fmt.Printf("Data in Block: %s\n", block.Data)
		fmt.Printf("Hash: %x\n", block.Hash)
		fmt.Println()
	}

}

/*
	OUTPUT:
		Block Number: 0
		Previous Block Hash:
		Data in Block: Genesis Block
		Hash: 5a19dc4d50692c48d3ad5cec3dd202cdeea8d7ecd2b3637e5852f2f99c7bf836

		Block Number: 1
		Previous Block Hash: 5a19dc4d50692c48d3ad5cec3dd202cdeea8d7ecd2b3637e5852f2f99c7bf836
		Data in Block: First Block after Genesis
		Hash: 65228225cb6367c7f5e81f28ccad36c07e92e43ad4b9bce0666afa14ea458d68

		Block Number: 2
		Previous Block Hash: 65228225cb6367c7f5e81f28ccad36c07e92e43ad4b9bce0666afa14ea458d68
		Data in Block: Second Block after Genesis
		Hash: c6cabc50b256e3753fdba82b2b78dd4c2185ffb8f141aa8dc174bcc9118ba0bf

		Block Number: 3
		Previous Block Hash: c6cabc50b256e3753fdba82b2b78dd4c2185ffb8f141aa8dc174bcc9118ba0bf
		Data in Block: Third Block after Genesis
		Hash: 397ea25629cafaeb4a57cfa2b584881bd89b012be81714da6cd2493bc2bcd932

		Block Number: 4
		Previous Block Hash: 397ea25629cafaeb4a57cfa2b584881bd89b012be81714da6cd2493bc2bcd932
		Data in Block: Fourth Block after Genesis
		Hash: 5bef99efed46b6627571da23a1d8778d1ea69f2a89d90128f0de4afcab127464

		Block Number: 5
		Previous Block Hash: 5bef99efed46b6627571da23a1d8778d1ea69f2a89d90128f0de4afcab127464
		Data in Block: Fifth Block after Genesis
		Hash: 3092bb7b8a0649c08a572029c28aed8e41cb315134e82a1dd32d2b6489cead1f
*/
