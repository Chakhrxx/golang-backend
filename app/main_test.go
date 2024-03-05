package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func TestCreateBlock(t *testing.T) {
	prevHash := []byte("PreviousHash")
	data := "Test Data"
	block := CreateBlock(data, prevHash)

	assert.NotNil(t, block)
	assert.Equal(t, block.Data, []byte(data))
	assert.Equal(t, block.PrevHash, prevHash)

	info := bytes.Join([][]byte{[]byte(data), prevHash}, []byte{})
	hash := sha256.Sum256(info)
	expectedHash := hex.EncodeToString(hash[:])

	assert.Equal(t, hex.EncodeToString(block.Hash), expectedHash)
}

func TestAddBlock(t *testing.T) {
	chain := InitBlockChain()

	chain.AddBlock("First Block after Genesis")
	assert.Len(t, chain.blocks, 2)

	chain.AddBlock("Second Block after Genesis")
	assert.Len(t, chain.blocks, 3)

	chain.AddBlock("Third Block after Genesis")
	assert.Len(t, chain.blocks, 4)
}

func TestGenesisBlock(t *testing.T) {
	genesis := Genesis()
	assert.NotNil(t, genesis)
	assert.Equal(t, genesis.Data, []byte("Genesis"))
	assert.Empty(t, genesis.PrevHash)
}

func TestInitBlockChain(t *testing.T) {
	chain := InitBlockChain()
	assert.NotNil(t, chain)
	assert.Len(t, chain.blocks, 1)

	genesis := chain.blocks[0]
	assert.Equal(t, genesis.Data, []byte("Genesis"))
	assert.Empty(t, genesis.PrevHash)
}

func TestAPIRoutes(t *testing.T) {
	app := setupTestApp()

	// Test GET /blocks
	req, err := http.NewRequest("GET", "/blocks", nil)
	assert.NoError(t, err)
	resp, err := app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)

	// Test POST /blocks
	req, err = http.NewRequest("POST", "/blocks", nil)
	assert.NoError(t, err)
	resp, err = app.Test(req)
	assert.NoError(t, err)
	assert.Equal(t, resp.StatusCode, 200)
}

func setupTestApp() *fiber.App {
	app := fiber.New()
	chain := InitBlockChain()

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

	return app
}