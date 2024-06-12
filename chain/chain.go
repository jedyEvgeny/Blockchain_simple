package chain

import (
	"fmt"

	"github.com/jedyEvgeny/Blockchain_simple/block"
)

type Chain struct {
	Blocks []*block.Block
}

func CreateBlockchain() *Chain {
	firstBlock := block.NewGenesisBlock()
	blockchain := []*block.Block{firstBlock}
	return &Chain{blockchain}
}

func (ch *Chain) AddBlock(data string) {
	lastIdxChain := len(ch.Blocks) - 1
	ptrPrevBlock := ch.Blocks[lastIdxChain]
	prevBlockHash := ptrPrevBlock.Hash
	newBlock := block.NewBlock(data, prevBlockHash)
	ch.Blocks = append(ch.Blocks, newBlock)
}

func PrintBlockChain(chain *Chain) {
	for _, block := range chain.Blocks {
		fmt.Printf("Предыдущий хеш: %x\n", block.PrevBlockHash)
		fmt.Printf("Данные: %s\n", block.Data)
		fmt.Printf("Хеш: %x\n", block.Hash)
		fmt.Println("-------------")
	}
}
