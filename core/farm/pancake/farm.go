package pancake

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"reinvest/core/farm/config"
	"reinvest/printer"
	"reinvest/token"
	"reinvest/utils"
)

type PancakeFarm struct {
	Printer         *printer.Printer
	FarmConfig      *config.FarmConfig
	Client          *ethclient.Client
	FarmInfo        *PoolInfo
	TokenBasic      *token.TokenBasic
	TokenAInfo      *token.Token
	TokenBInfo      *token.Token
	RewardTokenInfo *token.Token
	LpTokenInfo     *token.Token
}

type PendingReward struct {
	Amount      *big.Int
	TokenAmount *big.Int
}

func NewPanckeFarm(farmConfig *config.FarmConfig, client *ethclient.Client, tokenBasic *token.TokenBasic, printer *printer.Printer) *PancakeFarm {
	return &PancakeFarm{
		FarmConfig: farmConfig,
		Client:     client,
		TokenBasic: tokenBasic,
		Printer:    printer,
	}
}
func (c *PancakeFarm) RewardToken() *token.Token {
	return c.RewardTokenInfo
}

func (c *PancakeFarm) Harvest() (*big.Int, error) {

	c.Printer.Info("Harvest Reward...")
	pendingReward, err := c.Pending(c.FarmConfig.NetWork.FarmAddress, c.FarmConfig.Wallet, int(c.FarmConfig.PooID))
	if err != nil {
		return big.NewInt(0), err
	}
	c.Printer.Info(utils.ToDecimal(pendingReward.Amount, int(c.RewardTokenInfo.Decimals)).String() + " need to harvest")
	if pendingReward.Amount.Cmp(big.NewInt(0)) >= 1 {
		tx, err := c.Deposit(c.FarmConfig.NetWork.FarmAddress, big.NewInt(0), int(c.FarmConfig.PooID))
		if err != nil {
			c.Printer.Error("Harvest error: " + err.Error())
			return big.NewInt(0), err
		}
		txStatus, _tx := c.TokenBasic.WaitForBlockCompletation(tx.Hash().String())
		if txStatus != 1 {
			return big.NewInt(0), errors.New("Harvest Err Tx :" + tx.Hash().String())
		}
		sendRewardAmountToWallet, err := c.TokenBasic.GetTxAmount(c.RewardTokenInfo.Address, "", c.FarmConfig.Wallet, _tx)
		if err != nil {
			return big.NewInt(0), fmt.Errorf("Search Real Reward Amount To Wallet Error: %w", err)
		}
		c.Printer.Success(utils.ToDecimal(sendRewardAmountToWallet, int(c.RewardTokenInfo.Decimals)).String() + " " + c.RewardTokenInfo.Symbol + " -> " + c.FarmConfig.Wallet)

		return sendRewardAmountToWallet, nil
	}
	return big.NewInt(0), nil
}

func (c *PancakeFarm) SwapRewardToPairWithRetry(rewardAmount *big.Int, tryCount int) (*big.Int, *big.Int, string, string, error) {
	//	auth.Value = big.NewInt(9000000000000000000)     // in wei
	config.SwapToWBNB()
	swapTokenARealAmountBNB, swapTxHashBNB, err := c.Swap(big.NewInt(3000000000000000000), "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", "0x0E09FaBB73Bd3Ade0a17ECC321fD13a19e81cE82", tryCount)
	log.Print(swapTokenARealAmountBNB, swapTxHashBNB)
	rewardAmount = swapTokenARealAmountBNB

	//ada := "0x55d398326f99059fF775485246999027B3197955"
	//adb := "0xd32d01A43c869EdcD1117C640fBDcfCFD97d9d65"
	//swapa, _, err := c.Swap(big.NewInt(200000000000000000), "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", ada, tryCount)
	//swapb, _, err := c.Swap(big.NewInt(200000000000000000), "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c", adb, tryCount)
	//log.Print(swapa,swapb)
	//saStr := swapa.String()
	////88986860720233918128
	////14725265446135537687
	//sbStr := swapb.String()
	//log.Print(saStr,sbStr)
	//n := new(big.Int)
	//n, ok := n.SetString("13282543398492994102", 10)
	//if !ok {
	//	fmt.Println("SetString: error")
	//}
	//fmt.Println(n)
	//n2 := new(big.Int)
	//n2, ok = n.SetString("13282543398492994102", 10)
	//if !ok {
	//	fmt.Println("SetString: error")
	//}
	//fmt.Println(n2)
	//addLiquidityTx, err := c.addLiquidity(swapa, swapb, ada, adb)
	//log.Print(addLiquidityTx)

	pairAmount := make(map[string]*big.Int)
	avg := new(big.Int).Div(rewardAmount, big.NewInt(2))
	swapTokenARealAmount, swapTxHash, err := c.Swap(avg, c.RewardTokenInfo.Address, c.TokenAInfo.Address, tryCount)
	if err != nil {
		return nil, nil, c.TokenAInfo.Address, c.TokenBInfo.Address, err
	}
	pairAmount[c.TokenAInfo.Address] = swapTokenARealAmount
	c.Printer.Success(fmt.Sprintf(
		"Swap %s %s -> %s %s Tx: %s",
		utils.ToDecimal(avg, int(c.RewardTokenInfo.Decimals)).String(),
		c.RewardTokenInfo.Symbol,
		utils.ToDecimal(swapTokenARealAmount, int(c.TokenAInfo.Decimals)),
		c.TokenAInfo.Symbol,
		swapTxHash,
	))

	swapTokenBRealAmount, swapTxHash, err := c.Swap(avg, c.RewardTokenInfo.Address, c.TokenBInfo.Address, tryCount)
	if err != nil {
		return nil, nil, c.TokenAInfo.Address, c.TokenBInfo.Address, err
	}
	pairAmount[c.TokenBInfo.Address] = swapTokenBRealAmount
	c.Printer.Success(fmt.Sprintf(
		"Swap %s %s -> %s %s Tx: %s",
		utils.ToDecimal(avg, int(c.RewardTokenInfo.Decimals)).String(),
		c.RewardTokenInfo.Symbol,
		utils.ToDecimal(swapTokenBRealAmount, int(c.TokenBInfo.Decimals)),
		c.TokenBInfo.Symbol,
		swapTxHash,
	))

	return swapTokenARealAmount, swapTokenBRealAmount, c.TokenAInfo.Address, c.TokenBInfo.Address, err
}
func (c *PancakeFarm) AddLiquidityWithRetry(wishA *big.Int, wishB *big.Int, tokenAAddress, tokenBAddress string, tryCount int) (string, error) {
	count := 1
	var swapTxHash string
	for {
		if count >= tryCount {
			return swapTxHash, errors.New("Swap  Too Many errors")
		}
		if count > 1 {
			fmt.Printf("Try AddLiquidity %d \n", count)
		}
		//	0x55d398326f99059fF775485246999027B3197955
		// 0xd32d01A43c869EdcD1117C640fBDcfCFD97d9d65
		//344330000000000000000
		//204544830853370076776
		addLiquidityTx, err := c.addLiquidity(wishA, wishB, tokenAAddress, tokenBAddress)
		if err != nil {
			fmt.Printf("addLiquidity Error %s \n", err.Error())
			count++
			continue
		}
		swapTxHash = addLiquidityTx.Hash().String()
		c.Printer.Info(fmt.Sprintf("addLiquidity Tx : %s \n", addLiquidityTx.Hash().String()))
		addLiquidityTxStatus, _ := c.TokenBasic.WaitForBlockCompletation(addLiquidityTx.Hash().String())
		if addLiquidityTxStatus == 1 {
			//c.Printer.Success("%s %s + %s %s => %s Send To %s",utils.ToDecimal())
			return swapTxHash, nil
		}
		count++
	}
}

func (c *PancakeFarm) Reinvest() (*big.Int, string, error) {
	LpTokenBalance, err := c.TokenBasic.GetMyTokenInfo(c.LpTokenInfo.Address)
	if err != nil {
		c.Printer.Error(fmt.Sprintf("Get My LP Token Info  Err  %s \n", err))

		return big.NewInt(0), "", fmt.Errorf("Get My LP Token Info  Err  %w", err)
	}
	lpTokenApprove, err := c.TokenBasic.Approve(c.LpTokenInfo.Address, c.FarmConfig.NetWork.FarmAddress, LpTokenBalance.Balance)
	if err != nil {

		return big.NewInt(0), "", fmt.Errorf("Approve LP Token Error: %w", err)
	}
	if !lpTokenApprove {
		c.Printer.Error("Approve LP Token  Fail")

		return big.NewInt(0), "", errors.New("Approve LP Token  Fail")
	}
	depTx, err := c.Deposit(c.FarmConfig.NetWork.FarmAddress, LpTokenBalance.Balance, int(c.FarmConfig.PooID))
	if err != nil {

		return big.NewInt(0), "", fmt.Errorf("Deposit Error: %w", err)
	}
	depTxStatus, _ := c.TokenBasic.WaitForBlockCompletation(depTx.Hash().String())
	if depTxStatus != 1 {
		return big.NewInt(0), "", errors.New("Deposit Err Txh :" + depTx.Hash().String())
	}

	return LpTokenBalance.Balance, depTx.Hash().String(), nil
}
func (c *PancakeFarm) LpToken() *token.Token {
	return c.LpTokenInfo
}
func (c *PancakeFarm) Start() error {

	farmInfo, err := c.GetPoolInfo(c.FarmConfig.NetWork.FarmAddress, int(c.FarmConfig.PooID))
	if err != nil {
		return fmt.Errorf("Get LpToken Info Err %w ", err)
	}
	LpToken, err := NewLpToken(farmInfo.LpToken, c.Client)
	if err != nil {
		return fmt.Errorf("Get LpToken Info Err %w ", err)

	}
	address0, address1, err := LpToken.Pair()
	if err != nil {
		return fmt.Errorf("Get Token Pair Info Err %w ", err)

	}
	LpTokenInfo, err := c.TokenBasic.TokenInfo(farmInfo.LpToken)
	if err != nil {
		return fmt.Errorf("Get LpToken  Info Err %w  ", err)

	}

	rewardTokenInfo, err := c.TokenBasic.TokenInfo(c.FarmConfig.NetWork.RewardToken)
	if err != nil {
		return fmt.Errorf("Get Reward Token Info Err %w   ", err)
	}

	tokenAInfo, err := c.TokenBasic.TokenInfo(address0)

	if err != nil {
		return fmt.Errorf("Get My Token A Info  Err %w ", err)
	}
	tokenBInfo, err := c.TokenBasic.TokenInfo(address1)
	if err != nil {
		return fmt.Errorf("Get My Token B Info  Er %w ", err)
	}
	c.TokenAInfo = tokenAInfo
	c.TokenBInfo = tokenBInfo
	c.RewardTokenInfo = rewardTokenInfo
	c.LpTokenInfo = LpTokenInfo
	return nil
}
