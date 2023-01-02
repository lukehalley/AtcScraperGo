package web3

import (
	"github.com/ethereum/go-ethereum/common"
)

func GetHexAddress(Address interface{}) string {

	PairAddressCast, _ := Address.(common.Address)

	return PairAddressCast.Hex()
}
