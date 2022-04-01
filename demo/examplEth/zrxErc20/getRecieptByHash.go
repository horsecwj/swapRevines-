package zrxErc20

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"strings"
)

func GetLogAndReciept() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	//0xc31d7e7e85cab1d38ce1b8ac17e821ccd47dbde00f9d57f2bd8613bff9428396
	txHash := common.HexToHash("0xc31d7e7e85cab1d38ce1b8ac17e821ccd47dbde00f9d57f2bd8613bff9428396")
	//0x71cc07c759519cdc63c5b2a165fe1578ed26e7d40ebb32bcf687a1e6d93d94bf

	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	rca := receipt.ContractAddress.Hex()
	rcas := receipt.ContractAddress.String()
	rcas = strings.ToLower(rcas)
	//strings.tolo
	//0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f

	add := "0x5c69bee701ef814a2b6a3edd4b1652cb9cc5aa6f"
	add = strings.ToLower(add)
	if len(rca) != 0 && rcas == add {
		log.Print("contract  deploy ")
	}
	log.Print(receipt, rca)

}
