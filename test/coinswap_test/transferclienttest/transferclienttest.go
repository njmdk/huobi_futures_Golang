package transferclienttest

import (
	"huobi_futures_Golang/sdk/coinswap/restful"
	"huobi_futures_Golang/sdk/coinswap/restful/response/transfer"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/test/config"
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
