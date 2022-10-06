// package assignment01bca
package main

import (
	"crypto/sha256"
	"fmt"
	"strconv"
	"strings"
)

var counter int

type block struct {
	transaction string
	nonce int
	hash string
	previousHash string
} 

type blockchain struct{
	list []*block
}

func CalculateHash (stringToHash string) string {

	hash := sha256.Sum256([]byte(stringToHash))
	return fmt.Sprintf("%x", hash)
	
}



func NewBlock(transaction string, nonce int, previousHash string) *block {
	temp := new(block)
	temp.transaction = transaction
	temp.previousHash = previousHash
	temp.nonce = nonce
	str := strconv.Itoa(nonce)
	final := str + transaction + previousHash
	temp.hash = CalculateHash(final)
	counter++
	return temp
}

func (b *blockchain) ListBlocks() {
	for i := 0; i < len(b.list); i++ {
		fmt.Printf("%s List %d %s \n", strings.Repeat("=", 25), i, strings.Repeat("=", 25))
		fmt.Println(b.list[i].transaction)
		fmt.Println(b.list[i].previousHash)
		fmt.Println(b.list[i].hash)
		fmt.Println(b.list[i].nonce)
	}
}

func (b *blockchain)AddBlock( temp *block){
	b.list = append(b.list, temp)
}

func (b *blockchain)ChangeBlock(index int, transaction string) {
	b.list[index].transaction = transaction
}
func (b *blockchain)VerifyChain() bool {
	for i:= 0; i< len(b.list); i++ {
		str := strconv.Itoa(b.list[i].nonce)
		final := str + b.list[i].transaction + b.list[i].previousHash
		hash := CalculateHash(final)

		if hash != b.list[i].hash{
			fmt.Println("The BlockChain transaction was tempered in block no.", i)
			return false
		}
	}
	fmt.Println("The BlockChain is safe and untempered")
	return true
}

func main(){
	counter = 0

	chain := new(blockchain)
	//a := NewBlock("alice to bob", 3, "000")
	chain.AddBlock(NewBlock("alice to bob", 3, "000"));
	chain.AddBlock(NewBlock("bob to fatima", 7, chain.list[counter-1].hash));
	chain.AddBlock(NewBlock("fatima to adam", 7, chain.list[counter-1].hash));
	chain.ListBlocks()
	chain.VerifyChain()

	chain.ChangeBlock(1, "yousaf to fatima")
	chain.ListBlocks()
	chain.VerifyChain()
}
