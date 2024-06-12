package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  //Текущая временная метка
	Data          []byte //Основные данные
	PrevBlockHash []byte //Хеш предыдущего блока
	Hash          []byte //Хеш текущего блока
}

type Blockchain struct {
	blocks []*Block
}

func main() {
	bc := NewBlockchain()
	bc.AddBlock("Отправлено 1 BTC Евгению")
	bc.AddBlock("Отправлено ещё 30 BTC Евгению")
	for _, block := range bc.blocks {
		fmt.Printf("Предыдущий хеш: %x\n", block.PrevBlockHash)
		fmt.Printf("Данные: %s\n", block.Data)
		fmt.Printf("Хеш: %x\n", block.Hash)
		fmt.Println("-------------")
	}
}

// Вычисление хеша SHA-256
func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevBlockHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}}
	block.SetHash()
	return block
}

func (bc *Blockchain) AddBlock(data string) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	newBlock := NewBlock(data, prevBlock.Hash)
	bc.blocks = append(bc.blocks, newBlock)
}

// Создаём блок генезиса
func NewGenesisBlock() *Block {
	return NewBlock("Блок Генезиса", []byte{})
}

// Создаём блокчейн из блока Генезис
func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{NewGenesisBlock()}}
}
