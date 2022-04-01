package zrxErc20

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"

	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	exchange "reinvest/demo/examplEth/uniPair" // for demo
)

type UniPairTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

func GetUniTrans() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	if err != nil {
		log.Fatal(err)
	}
	// 0x Protocol Exchange smart contract address
	contractAddress := common.HexToAddress("0xBb2b8038a1640196FbE3e38816F3e67Cba72D940")

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(14197211),
		ToBlock:   big.NewInt(14197211),
		Addresses: []common.Address{
			contractAddress,
		},
	}
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {

		log.Fatal(err)
	}
	contractAbi, err := abi.JSON(strings.NewReader(exchange.UniPairABI))
	if err != nil {
		log.Fatal(err)
	}
	// NOTE: keccak256("LogFill(address,address,address,address,address,uint256,uint256,uint256,uint256,bytes32,bytes32)")
	logFillEvent := common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")
	var allTrans []UniPairTransfer
	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)
		switch vLog.Topics[0].Hex() {
		case logFillEvent.Hex():
			fmt.Printf("Log Name: transfer\n")
			var fillEvent UniPairTransfer
			res1 := make(map[string]interface{})
			err := contractAbi.Events["Transfer"].Inputs.UnpackIntoMap(res1, vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			fillEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
			fillEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
			fillEvent.Value = res1["value"].(*big.Int)
			allTrans = append(allTrans, fillEvent)
		}
	}

	fmt.Printf("解析完成")

}
