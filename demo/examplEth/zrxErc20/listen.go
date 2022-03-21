package zrxErc20

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/spf13/cast"
	"log"
	"math/big"
	exchange "reinvest/demo/examplEth/contracts_0xprotocol" // for demo
	"strconv"
	"strings"
)

//    event LogFill(
//        address indexed maker,
//        address taker,
//        address indexed feeRecipient,
//        address makerToken,
//        address takerToken,
//        uint filledMakerTokenAmount,
//        uint filledTakerTokenAmount,
//        uint paidMakerFee,
//        uint paidTakerFee,
//        bytes32 indexed tokens, // keccak256(makerToken, takerToken), allows subscribing to a token pair
//        bytes32 orderHash
//    );

// LogFill ...
type LogFill struct {
	Maker                  common.Address
	Taker                  common.Address
	FeeRecipient           common.Address
	MakerToken             common.Address
	TakerToken             common.Address
	FilledMakerTokenAmount *big.Int
	FilledTakerTokenAmount *big.Int
	PaidMakerFee           *big.Int
	PaidTakerFee           *big.Int
	Tokens                 [32]byte
	OrderHash              [32]byte
}

//    function fillOrder(
//          address[5] orderAddresses,
//          uint[6] orderValues,
//          uint fillTakerTokenAmount,
//          bool shouldThrowOnInsufficientBalanceOrAllowance,
//          uint8 v,
//          bytes32 r,
//          bytes32 s)
//          public
//          returns (uint filledTakerTokenAmount);

// LogFill ...

type FillOrder struct {
	OrderAddresses                              common.Address
	OrderValues                                 []big.Int
	FillTakerTokenAmount                        *big.Int
	ShouldThrowOnInsufficientBalanceOrAllowance bool
	V                                           uint8
	R                                           [32]byte
	S                                           [32]byte
}

// LogCancel ...
type LogCancel struct {
	Maker                     common.Address
	FeeRecipient              common.Address
	MakerToken                common.Address
	TakerToken                common.Address
	CancelledMakerTokenAmount *big.Int
	CancelledTakerTokenAmount *big.Int
	Tokens                    [32]byte
	OrderHash                 [32]byte
}

// LogError ...
type LogError struct {
	ErrorID   uint8
	OrderHash [32]byte
}

func decodeTxParams(abi abi.ABI, v map[string]interface{}, data []byte) (map[string]interface{}, error) {
	m, err := abi.MethodById(data[:4])
	if err != nil {
		return map[string]interface{}{}, err
	}

	if err := m.Inputs.UnpackIntoMap(v, data[4:]); err != nil {
		return map[string]interface{}{}, err
	}
	return v, nil
}

func main() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Exchange smart contract address
	contractAddress := common.HexToAddress("0x12459C951127e0c374FF9105DdA097662A027093")
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(6383482),
		ToBlock:   big.NewInt(6383488),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(exchange.ExchangeABI))
	if err != nil {
		log.Fatal(err)
	}

	// NOTE: keccak256("LogFill(address,address,address,address,address,uint256,uint256,uint256,uint256,bytes32,bytes32)")
	logFillEvent := common.HexToHash("0d0b9391970d9a25552f37d436d2aae2925e2bfe1b2a923754bada030c498cb3")

	// NOTE: keccak256("LogCancel(address,address,address,address,uint256,uint256,bytes32,bytes32)")
	logCancelEvent := common.HexToHash("67d66f160bc93d925d05dae1794c90d2d6d6688b29b84ff069398a9b04587131")

	// NOTE: keccak256("LogError(uint8,bytes32)")
	logErrorEvent := common.HexToHash("36d86c59e00bd73dc19ba3adfe068e4b64ac7e92be35546adeddf1b956a87e90")

	for _, vLog := range logs {
		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
		fmt.Printf("Log Index: %d\n", vLog.Index)

		switch vLog.Topics[0].Hex() {
		case logFillEvent.Hex():
			fmt.Printf("Log Name: LogFill\n")

			var fillEvent LogFill
			//contractAbi.Methods["transfer"].Inputs.UnpackValues("")
			rs := contractAbi.Methods[""]
			log.Print(rs)
			tHash := vLog.TxHash.String()
			log.Print(tHash)
			res1, err := contractAbi.Events["LogFill"].Inputs.UnpackValues(vLog.Data)
			log.Print(res1[0].(common.Address).String())
			res, err := contractAbi.Unpack("LogFill", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(res)

			fillEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			fillEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
			fillEvent.Tokens = vLog.Topics[3]

			fmt.Printf("Maker: %s\n", fillEvent.Maker.Hex())
			fmt.Printf("Taker: %s\n", fillEvent.Taker.Hex())
			fmt.Printf("Fee Recipient: %s\n", fillEvent.FeeRecipient.Hex())
			fmt.Printf("Maker Token: %s\n", fillEvent.MakerToken.Hex())
			fmt.Printf("Taker Token: %s\n", fillEvent.TakerToken.Hex())
			fmt.Printf("Filled Maker Token Amount: %s\n", fillEvent.FilledMakerTokenAmount.String())
			fmt.Printf("Filled Taker Token Amount: %s\n", fillEvent.FilledTakerTokenAmount.String())
			fmt.Printf("Paid Maker Fee: %s\n", fillEvent.PaidMakerFee.String())
			fmt.Printf("Paid Taker Fee: %s\n", fillEvent.PaidTakerFee.String())
			fmt.Printf("Tokens: %s\n", hexutil.Encode(fillEvent.Tokens[:]))
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(fillEvent.OrderHash[:]))

		case logCancelEvent.Hex():
			fmt.Printf("Log Name: LogCancel\n")

			var cancelEvent LogCancel

			res, err := contractAbi.Unpack("LogCancel", vLog.Data)
			if err != nil {
				log.Fatal(err)
			}
			log.Print(res)

			cancelEvent.Maker = common.HexToAddress(vLog.Topics[1].Hex())
			cancelEvent.FeeRecipient = common.HexToAddress(vLog.Topics[2].Hex())
			cancelEvent.Tokens = vLog.Topics[3]

			fmt.Printf("Maker: %s\n", cancelEvent.Maker.Hex())
			fmt.Printf("Fee Recipient: %s\n", cancelEvent.FeeRecipient.Hex())
			fmt.Printf("Maker Token: %s\n", cancelEvent.MakerToken.Hex())
			fmt.Printf("Taker Token: %s\n", cancelEvent.TakerToken.Hex())
			fmt.Printf("Cancelled Maker Token Amount: %s\n", cancelEvent.CancelledMakerTokenAmount.String())
			fmt.Printf("Cancelled Taker Token Amount: %s\n", cancelEvent.CancelledTakerTokenAmount.String())
			fmt.Printf("Tokens: %s\n", hexutil.Encode(cancelEvent.Tokens[:]))
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(cancelEvent.OrderHash[:]))

		case logErrorEvent.Hex():
			fmt.Printf("Log Name: LogError\n")

			errorID, err := strconv.ParseInt(vLog.Topics[1].Hex(), 16, 64)
			if err != nil {
				log.Fatal(err)
			}

			errorEvent := &LogError{
				ErrorID:   uint8(errorID),
				OrderHash: vLog.Topics[2],
			}

			fmt.Printf("Error ID: %d\n", errorEvent.ErrorID)
			fmt.Printf("Order Hash: %s\n", hexutil.Encode(errorEvent.OrderHash[:]))
		}

		fmt.Printf("\n\n")
	}
}

func GetLog() {
	client, err := ethclient.Dial("https://mainnet.infura.io/v3/2da8854f387e471f9063be2848f6f9a2")
	if err != nil {
		log.Fatal(err)
	}

	contractAbi, err := abi.JSON(strings.NewReader(exchange.ExchangeABI))
	if err != nil {
		log.Fatal(err)
	}

	// 0x Protocol Exchange smart contract address
	//contractAddress := common.HexToAddress("0x12459C951127e0c374FF9105DdA097662A027093")

	txHash := common.HexToHash("0xa2fc965232d6c54798f153992656e3cefa3a60d42dd96d7fdd9bcf6851478425")
	tx, _, err := client.TransactionByHash(context.Background(), txHash)
	if err != nil {
		log.Fatal(err)
	}
	temp := make(map[string]interface{})
	input := tx.Data()
	decodeTxParams(contractAbi, temp, input)
	temp2 := make(map[string]interface{})
	for key, itme := range temp {
		temp2[key] = cast.ToString(itme)
	}
	log.Print(1)
}
