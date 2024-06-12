package main

import (
	"github.com/jedyEvgeny/Blockchain_simple/chain"
)

func main() {
	bc := chain.CreateBlockchain()
	bc.AddBlock("Отправлено 1 BTC Евгению")
	bc.AddBlock("Отправлено ещё 30 BTC Евгению")
	chain.PrintBlockChain(bc)
}
