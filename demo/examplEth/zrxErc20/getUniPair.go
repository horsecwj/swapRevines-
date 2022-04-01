package zrxErc20

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	store "reinvest/demo/examplEth/uniPair"
)

func UseUniPair() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0xBb2b8038a1640196FbE3e38816F3e67Cba72D940")
	instance, err := store.NewUniPair(address, client)
	if err != nil {
		log.Fatal(err)
	}
	token1, err := instance.Token0(nil)
	if err != nil {
		log.Fatal(err)
	}
	token2, err := instance.Token1(nil)

	bal, err := instance.BalanceOf(nil, address)
	if err != nil {
		return
	}
	number := bal.Int64()
	addr := token1.Hex()
	log.Print(token1, token2, bal, number, addr)

}
