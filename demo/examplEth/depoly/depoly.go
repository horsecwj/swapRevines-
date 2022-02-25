package depoly

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	store "reinvest/demo/examplEth/contracts"
)

func depoly() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/c561706d7070475ab1b59071ee4684b0")
	if err != nil {
		log.Fatal(err)
	}
	//	address := common.HexToAddress("0x52001daE367173D09f1Ddf435d6B2776BCA0920C")
	privateKey, err := crypto.HexToECDSA("956fb3f29e34a14c603f458ffdee4b526a7f6b6b918f6d5f3a9ea7c533fa6b6b")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	input := "1.0"
	address, tx, instance, err := store.DeployStore(auth, client, input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // 0x147B8eb97fD247D06C4006D269c90C1908Fb5D54
	fmt.Println(tx.Hash().Hex()) // 0xdae8ba5444eefdc99f4d45cd0c4f24056cba6a02cefbf78066ef9f4188ff7dc0
	//0x52001daE367173D09f1Ddf435d6B2776BCA0920C
	//0x6a6b2a8ea5b826730d3d3cc83337f2c53fbbeb719b505c1ca593234611deff32
	_ = instance
}
