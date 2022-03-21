package zrxErc20

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	store "reinvest/demo/examplEth/erc20Tokern"
)

func UseErc20() {

	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	if err != nil {
		log.Fatal(err)
	}
	address := common.HexToAddress("0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599")
	instance, err := store.NewToken(address, client)
	if err != nil {
		log.Fatal(err)
	}
	name, err := instance.Name(nil)

	log.Print(name)

}
