package transferclienttest

import (
	"github.com/njmdk/huobi_futures_Golang/config"
	"github.com/njmdk/huobi_futures_Golang/sdk/coinswap/restful"
	"github.com/njmdk/huobi_futures_Golang/sdk/coinswap/restful/response/transfer"
	"github.com/njmdk/huobi_futures_Golang/sdk/log"
)

func RunAllExamples() {
	AccountTransferAsync()
}

func AccountTransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp := make(chan transfer.AccountTransferResponse)
	go client.AccountTransferAsync(resp, "", "", "", "")
	x := <-resp
	log.Info("%v", x)
}
