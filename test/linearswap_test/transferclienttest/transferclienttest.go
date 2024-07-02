package transferclienttest

import (
	"github.com/njmdk/huobi_futures_Golang/config"
	"github.com/njmdk/huobi_futures_Golang/sdk/linearswap/restful"
	"github.com/njmdk/huobi_futures_Golang/sdk/linearswap/restful/response/transfer"
	"github.com/njmdk/huobi_futures_Golang/sdk/log"
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
