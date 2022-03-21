package zrxErc20

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	store "reinvest/demo/examplEth/contractFactory"
)

func UseUniFac() {

	address := common.HexToAddress("0x5C69bEe701ef814a2B6a3EDD4B1652CB9cc5aA6f")

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	if err != nil {
		log.Fatal(err)
	}
	instance, err := store.NewUniFac(address, client)
	if err != nil {
		log.Fatal(err)
	}
	length, err := instance.AllPairsLength(nil)
	if err != nil {
		log.Fatal(err)
	}
	pair, err := instance.AllPairs(nil, big.NewInt(1000))
	addr := pair.Hex()
	log.Print(addr, length)

}
