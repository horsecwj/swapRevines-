package pancake

import (
	"context"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"log"
	"math/big"
	"reinvest/core/farm/pancake/contracts"
	"reinvest/utils"
)

type PoolInfo struct {
	FarmContract      *contracts.PancakeFarm
	LpToken           string
	AllocPoint        string
	LastRewardBlock   string
	AccMdxPerShare    string
	AccMultLpPerShare string
	TotalAmount       string
}

func (c *PancakeFarm) GetPoolInfo(farmAddress string, pool int) (*PoolInfo, error) {
	if !utils.IsValidAddress(farmAddress) {
		return nil, errors.New("Farm Address Is InValid!")
	}
	farm, err := contracts.NewPancakeFarm(common.HexToAddress(farmAddress), c.Client)
	if err != nil {
		return nil, fmt.Errorf("Get Farm Error : %w", err)
	}
	poolInfo, err := farm.PoolInfo(&bind.CallOpts{}, new(big.Int).SetInt64(int64(pool)))
	if err != nil {
		return nil, fmt.Errorf("Get Pool Info Error : %w", err)
	}
	return &PoolInfo{
		FarmContract:    farm,
		LpToken:         poolInfo.LpToken.String(),
		AllocPoint:      poolInfo.AllocPoint.String(),
		LastRewardBlock: poolInfo.LastRewardBlock.String(),
	}, nil
}
func (c *PancakeFarm) Swap(amount *big.Int, fromToken string, toToken string, tryCount int) (*big.Int, string, error) {
	if utils.ToValidateAddress(toToken) != utils.ToValidateAddress(fromToken) {
		sendAmountToWallet, swapTxHash, err := c.swapWithRetry(amount, fromToken, toToken, tryCount)
		if err != nil {
			return nil, swapTxHash, fmt.Errorf("Swap error %w Tx: %s", err, swapTxHash)
		}
		return sendAmountToWallet, swapTxHash, nil
	}
	return amount, "", nil
}

//配对
func (c *PancakeFarm) addLiquidity(wishA *big.Int, wishB *big.Int, tokenA, tokenB string) (*types.Transaction, error) {
	if wishA.Cmp(big.NewInt(0)) <= 0 {
		return nil, fmt.Errorf("Error Token A Wish Amount ")
	}
	if wishB.Cmp(big.NewInt(0)) <= 0 {
		return nil, fmt.Errorf("Error Token B Wish Amount ")
	}
	approved, err := c.TokenBasic.Approve(tokenA, c.FarmConfig.NetWork.Router, wishA)
	if err != nil {
		return nil, fmt.Errorf("Approve Token A Error : %w", err)
	}
	if !approved {
		return nil, fmt.Errorf("Approve Token A Fail")
	}
	//fmt.Println(green("Has Approved To :" + router + " for Token A :" + tokenA))
	tokenBApproved, err := c.TokenBasic.Approve(tokenB, c.FarmConfig.NetWork.Router, wishB)
	if err != nil {
		return nil, fmt.Errorf("Approve Token B Error : %w", err)
	}
	if !tokenBApproved {
		return nil, fmt.Errorf("Approve Token B Fail")
	}
	swapRouter, err := NewSwapRouter(c.FarmConfig.NetWork.Router, c.Client, c)
	if err != nil {
		return nil, err
	}
	factory, err := NewSwapFactory(swapRouter.Factory, c.Client, c)
	if err != nil {
		return nil, err
	}
	lpToken, err := NewLpToken(c.LpTokenInfo.Address, c.Client)
	if err != nil {
		return nil, err
	}
	currentCanPairTokenB, err := swapRouter.TokenBPairAmount(tokenA, tokenB, wishA, lpToken)
	if err != nil {
		return nil, err
	}
	minB := factory.Calc(currentCanPairTokenB, 0.99)
	minA := factory.Calc(wishA, 0.99)
	//minA = big.NewInt(0)
	//minB = big.NewInt(0)
	minAStr := minA.String()
	//84763493634115758523
	//98312180835707786
	//saStr := swapa.String()
	//88986860720233918128
	//14725265446135537687
	minBStr := minB.String()
	log.Print(minBStr, minAStr)
	//fmt.Println(green("Has Approved To: " + router + " for Token B: " + tokenB))
	return swapRouter.AddLiquidity(tokenA, tokenB, wishA, wishB, minA, minB)
}

func (c *PancakeFarm) swapWithRetry(amount *big.Int, fromToken string, toToken string, tryCount int) (*big.Int, string, error) {
	count := 1
	keepSwap := true
	var swapTxHash string
	for {
		if count >= tryCount {
			return nil, swapTxHash, errors.New("Swap  Too Many errors")
		}
		if count > 1 {
			fmt.Printf("Try Swap %d \n", count)
		}
		if keepSwap {
			tx, err := c.SwapExactTokenTo(amount, fromToken, toToken)
			if err != nil {
				log.Println("swap error " + err.Error())
				count++
				continue
			}
			swapTxHash = tx.Hash().String()
		}
		swapTxStatus, _tx := c.TokenBasic.WaitForBlockCompletation(swapTxHash)
		if swapTxStatus == 1 {
			keepSwap = false
			sendAmountToWallet, err := c.TokenBasic.GetTxAmount(toToken, "", c.FarmConfig.Wallet, _tx)
			if err != nil {
				log.Println("Swap Error: " + err.Error())
				count++
				continue
			}
			return sendAmountToWallet, swapTxHash, nil
		}

		count++

	}
}

func (c *PancakeFarm) GetCake() {
	//fromAddr := "0xa3b0422fb23d8e0f0eaf243cda405dc12ecf2932"
	//fromAddrs := common.HexToAddress(fromAddr)

}

func (c *PancakeFarm) SwapExactTokenTo(rewardAmount *big.Int, from, to string) (*types.Transaction, error) {
	approved, err := c.TokenBasic.Approve(from, c.FarmConfig.NetWork.Router, rewardAmount)
	if err != nil {
		return nil, fmt.Errorf("Approve Swap Token Error : %w", err)
	}
	if !approved {
		return nil, fmt.Errorf("Approve Swap Token Fail")
	}
	//fmt.Println(green("Has Approved To :" + router + " for :" + from))
	swapRouter, err := NewSwapRouter(c.FarmConfig.NetWork.Router, c.Client, c)
	if err != nil {
		return nil, err
	}
	factory, err := NewSwapFactory(swapRouter.Factory, c.Client, c)

	wishAmount, err := swapRouter.WishExchange(rewardAmount, from, to)
	minExchange := factory.Calc(wishAmount[1], 0.005)

	tx, err := swapRouter.SwapExactTokenTo(from, to, rewardAmount, minExchange)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (c *PancakeFarm) SwapExactEthToToken(rewardAmount *big.Int, from, to string) (*types.Transaction, error) {

	approved, err := c.TokenBasic.Approve(from, c.FarmConfig.NetWork.Router, rewardAmount)
	if err != nil {
		return nil, fmt.Errorf("Approve Swap Token Error : %w", err)
	}
	if !approved {
		return nil, fmt.Errorf("Approve Swap Token Fail")
	}
	//fmt.Println(green("Has Approved To :" + router + " for :" + from))
	swapRouter, err := NewSwapRouter(c.FarmConfig.NetWork.Router, c.Client, c)
	if err != nil {
		return nil, err
	}
	factory, err := NewSwapFactory(swapRouter.Factory, c.Client, c)

	wishAmount, err := swapRouter.WishExchange(rewardAmount, from, to)
	minExchange := factory.Calc(wishAmount[1], 0.005)

	tx, err := swapRouter.SwapExactTokenTo(from, to, rewardAmount, minExchange)
	if err != nil {
		return nil, err
	}
	return tx, nil

}
func (c *PancakeFarm) Pending(farmAddress string, wallet string, pool int) (*PendingReward, error) {
	if !utils.IsValidAddress(farmAddress) {
		return &PendingReward{
			Amount: big.NewInt(0),
		}, errors.New("Farm Address Is InValid!")
	}
	if !utils.IsValidAddress(wallet) {
		return &PendingReward{
			Amount: big.NewInt(0),
		}, errors.New("Wallet Address Is InValid!")
	}
	//质押池信息
	poolInfo, err := c.GetPoolInfo(farmAddress, pool)
	if err != nil {
		return &PendingReward{
			Amount: big.NewInt(0),
		}, fmt.Errorf("Get Pool Info Error : %v", err)
	}
	amount, err := poolInfo.FarmContract.PendingCake(&bind.CallOpts{}, new(big.Int).SetInt64(int64(pool)), common.HexToAddress(wallet))
	if err != nil {
		return &PendingReward{
			Amount: big.NewInt(0),
		}, fmt.Errorf("Get Pool Pending Reward  Error : %v", err)

	}

	return &PendingReward{
		Amount: amount,
	}, nil
}
func (c *PancakeFarm) Deposit(farmAddress string, amount *big.Int, pool int) (*types.Transaction, error) {

	if !utils.IsValidAddress(farmAddress) {
		return nil, errors.New("Farm Address Is InValid!")
	}
	poolInfo, err := c.GetPoolInfo(farmAddress, pool)
	if err != nil {
		return nil, fmt.Errorf("Get Pool Info Error : %w", err)
	}
	wallet := c.FarmConfig.Wallet
	gasPrice, err := c.Client.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, fmt.Errorf("Get Gas Price Error %w", err)
	}
	nonce, err := c.Client.PendingNonceAt(context.Background(), common.HexToAddress(wallet))
	if err != nil {
		return nil, fmt.Errorf("Get Nonce Error %w", err)
	}
	auth, err := c.TokenBasic.CreateTx()
	if err != nil {
		return nil, fmt.Errorf("Get Create Transaction Error %w", err)
	}
	auth.GasPrice = gasPrice
	auth.From = common.HexToAddress(wallet)
	auth.GasLimit = uint64(300000)
	auth.Context = context.Background()
	auth.Nonce = big.NewInt(int64(nonce))
	tx, err := poolInfo.FarmContract.Deposit(auth, big.NewInt(int64(pool)), amount)
	if err != nil {
		return nil, fmt.Errorf("Deposit Err %w", err)
	}
	return tx, nil
}

type FarmUserInfo struct {
	Amount           *big.Int
	RewardDebt       string
	MultLpRewardDebt string
}

//获取我的信息
func (c *PancakeFarm) GetFarmUserInfo(farmAddress string, wallet string, pool int) (*FarmUserInfo, error) {
	if !utils.IsValidAddress(farmAddress) {
		return nil, errors.New("Farm Address Is InValid!")
	}
	if !utils.IsValidAddress(wallet) {
		return nil, errors.New("Wallet Address Is InValid!")
	}
	//质押池信息
	poolInfo, err := c.GetPoolInfo(farmAddress, pool)
	if err != nil {
		return nil, fmt.Errorf("Get Pool Info  Error : %v", err)
	}

	userInfo, err := poolInfo.FarmContract.UserInfo(&bind.CallOpts{}, new(big.Int).SetInt64(int64(pool)), common.HexToAddress(wallet))
	if err != nil {
		return nil, fmt.Errorf("Get User Info  Error : %v", err)
	}

	return &FarmUserInfo{
		Amount:     userInfo.Amount,
		RewardDebt: userInfo.RewardDebt.String(),
	}, nil
}
