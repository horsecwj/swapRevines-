package store

import (
	rz "reinvest/demo/examplEth/zrxErc20"
	"testing"
)

func TestParse(t *testing.T) {
	//clineEth()
	//creatwallet()
	//creatKeystore()
	//
	//	trans()
	//	transfer()
	erc20Transfer()
	//deployTest()
}
func TestGetLog(t *testing.T) {
	GetLog()
}
func TestGetLog2(t *testing.T) {
	rz.GetLog()
}
func TestGetLog3(t *testing.T) {
	rz.UseUniFac()
}

func TestGetLog4(t *testing.T) {
	rz.UseUniPair()
}
func TestGetLog5(t *testing.T) {
	rz.UseErc20()
}
func TestGetLog6(t *testing.T) {
	rz.GetUniTrans()
}

//GetLogAndReciept
func TestGetLog7(t *testing.T) {
	rz.GetLogAndReciept()
}
