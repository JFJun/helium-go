package http

import (
	"fmt"
	"log"
	"testing"
	"time"
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
	data, err := hr.GetBlockTransactionByHeight(364896, "")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(len(data.Data))
	fmt.Println(data)

}

func TestHeliumRpc_GetAccountByAddress(t *testing.T) {
	resp, err := hr.GetAccountByAddress("13ZCNF5eHUAsrKcVd9TAfcAWgceCxmuGazqUqbaGEjA8RApaZWG")
	if err != nil {
		log.Fatal(resp)
	}
	balance := resp.Balance
	fmt.Println(balance)
	fmt.Println(time.Now())
	for {
		resp2, err := hr.GetAccountByAddress("13ZCNF5eHUAsrKcVd9TAfcAWgceCxmuGazqUqbaGEjA8RApaZWG")
		if err != nil {
			log.Fatal(resp)
		}
		fmt.Println(balance)
		if balance != resp2.Balance {
			fmt.Println(time.Now())
			break
		}
		time.Sleep(10 * time.Second)
	}
}

func TestHeliumRpc_GetPendingTransactionByTxid(t *testing.T) {
	_, err := hr.GetPendingTransactionByTxid("16oj8uE5DDduoZooCSZDWSDAszfKMNqWredQA5opw9Cxt1jy97")
	if err != nil {
		fmt.Println(err)
	}
}

func TestHeliumRpc_GetTransactionByTxid(t *testing.T) {
	resp, err := hr.GetTransactionByTxid("ptiM25qhPpEM-w0lHsGppVyeRFxrmNwwR5-vuv-aUjs")
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(resp)
}
func TestHeliumRpc_GetVars(t *testing.T) {
	resp, err := hr.GetVars()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.DcPayloadSize)
	fmt.Println(resp.TxnFeeMultiplier)
}
func TestHeliumRpc_GetCurrentPrices(t *testing.T) {
	resp, err := hr.GetCurrentPrices()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(resp.Price)
	fmt.Println(resp.Block)
}
