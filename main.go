package main

import (
	"bytes"
	"crypto/sha256"
	"github.com/gofiber/fiber/v2"
)

// BlockChain represents a chain of blocks in a blockchain
type BlockChain struct {
   blocks []*Block // The slices of blocks in the blockchain
}

// Block represents a single block in a blockchain
type Block struct {
   Hash     []byte // Hash of the block's data
   Data     []byte // The actual data contained within the block
   PrevHash []byte // Hash of the previous block in the chain
}

// DeriveHash calculates and sets the hash for the block based on its data and previous hash
func (b *Block) DeriveHash() {
   info := bytes.Join([][]byte{b.Data, b.PrevHash}, []byte{})
   hash := sha256.Sum256(info)
   b.Hash = hash[:]
}

// CreateBlock creates a new block with the given data and previous hash
func CreateBlock(data string, prevHash []byte) *Block {
   block := &Block{[]byte{}, []byte(data), prevHash}
   block.DeriveHash()
   return block
}

// AddBlock appends a new block with the given data to the blockchain
func (chain *BlockChain) AddBlock(data string) {
   prevBlock := chain.blocks[len(chain.blocks)-1]
   new := CreateBlock(data, prevBlock.Hash)
   chain.blocks = append(chain.blocks, new)
}

// Genesis creates a genesis block for the blockchain
func Genesis() *Block {
   return CreateBlock("Genesis", []byte{})
}

// InitBlockChain initializes a new blockchain with a genesis block
func InitBlockChain() *BlockChain {
   return &BlockChain{[]*Block{Genesis()}}
}

func main() {
	app := fiber.New()
	chain := InitBlockChain()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello go fiber")
	})

	// GET endpoint to retrieve all blocks in the blockchain
	app.Get("/blocks", func(c *fiber.Ctx) error {
		return c.JSON(chain.blocks)
	})

	// POST endpoint to add a new block to the blockchain
	app.Post("/blocks", func(c *fiber.Ctx) error {
		data := c.FormValue("data")
		chain.AddBlock(data)
		return c.SendString("Block added successfully")
	})

	// Run the Fiber app on port  8080
	app.Listen(": 8080")
}