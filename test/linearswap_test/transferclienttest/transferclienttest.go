package transferclienttest

import (
	"huobi_futures_Golang/sdk/linearswap/restful"
	"huobi_futures_Golang/sdk/linearswap/restful/response/transfer"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/test/config"
)

func RunAllExamples() {
	TransferAsync()
}

func TransferAsync() {
	client := new(restful.TransferClient).Init(config.AccessKey, config.SecretKey, config.Host)
	resp := make(chan transfer.TransferResponse)
	go client.TransferAsync(resp, "", "", 0, "", "")
	x := <-resp
	log.Info("%v", x)
}
