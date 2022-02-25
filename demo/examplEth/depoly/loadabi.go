package depoly

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	store "reinvest/demo/examplEth/contracts"
)

func loadContract() {

	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/c561706d7070475ab1b59071ee4684b0")
	if err != nil {
		log.Fatal(err)
	}
	//	privateKey, err := crypto.HexToECDSA("956fb3f29e34a14c603f458ffdee4b526a7f6b6b918f6d5f3a9ea7c533fa6b6b")
	address := common.HexToAddress("0x52001daE367173D09f1Ddf435d6B2776BCA0920C")
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "
}
