package transferclienttest

import (
	"huobi_futures_Golang/config"
	"huobi_futures_Golang/sdk/linearswap/restful"
	"huobi_futures_Golang/sdk/linearswap/restful/response/transfer"
	"huobi_futures_Golang/sdk/log"
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
