package http

import (
	"fmt"
	"log"
	"testing"
)

var hr = NewHeliumRpc("https://api.helium.io")

func TestHeliumRpc_GetLatestBlockHeight(t *testing.T) {
	height, _ := hr.GetLatestBlockHeight()
	fmt.Println(height)
}
func TestHeliumRpc_GetBlock(t *testing.T) {
	hr.GetBlock(304523)
}

func TestHeliumRpc_GetBlockTransactionByHeight(t *testing.T) {
	data, err := hr.GetBlockTransactionByHeight(307458, "")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(len(data.Data))
	fmt.Println(data)

}

func TestHeliumRpc_GetAccountByAddress(t *testing.T) {
	resp, err := hr.GetAccountByAddress("14CEs4NCE8h9QRDY56ud3W7Euk9usuYUNsvTdQpoHGjkUb1GJTB")
	if err != nil {
		log.Fatal(resp)
		return
	}
	fmt.Println(resp.Balance)
}

func TestHeliumRpc_GetPendingTransactionByTxid(t *testing.T) {
	_, err := hr.GetPendingTransactionByTxid("16oj8uE5DDduoZooCSZDWSDAszfKMNqWredQA5opw9Cxt1jy97")
	if err != nil {
		fmt.Println(err)
	}
}

func TestHeliumRpc_GetTransactionByTxid(t *testing.T) {
	resp, err := hr.GetTransactionByTxid("DS6xLdxg4HDVMyntngf7OPPfY3-7wofJer8fD9VzAs8")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(resp)
}
