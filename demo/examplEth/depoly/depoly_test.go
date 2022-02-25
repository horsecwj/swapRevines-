package depoly

import (
	store "reinvest/demo/examplEth/erc20Tokern"
	socket "reinvest/demo/examplEth/socket"
	"testing"
)

func TestA(t *testing.T) {
	depoly()
}

func Test(t *testing.T) {
	loadContract()
}
func Test2(t *testing.T) {
	transStore()
}

func Test3(t *testing.T) {
	store.Opration()
}

//listenToken

func Test4(t *testing.T) {
	socket.ListenToken()
}

//Erc20Listen

func Test5(t *testing.T) {
	socket.Erc20Listen()
}
