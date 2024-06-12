package block

import (
	"bytes"
	"crypto/sha256"
	"strconv"
	"time"
)

type Block struct {
	Timestamp     int64  // Текущая временная метка
	Data          []byte // Основные данные
	PrevBlockHash []byte // Хеш предыдущего блока
	Hash          []byte // Хеш текущего блока
}

func NewGenesisBlock() *Block {
	return NewBlock("Блок Генезиса", []byte{})
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	now := time.Now().Unix()
	block := &Block{
		Timestamp:     now,
		Data:          []byte(data),
		PrevBlockHash: prevBlockHash,
	}
	block.SetHash()
	return block
}

func (b *Block) SetHash() {
	timestamp := []byte(strconv.Itoa(int(b.Timestamp)))
	arrAllFieldsBlock := [][]byte{b.PrevBlockHash, b.Data, timestamp}
	headers := bytes.Join(arrAllFieldsBlock, []byte{})
	hash := sha256.Sum256(headers)
	b.Hash = hash[:]
}
