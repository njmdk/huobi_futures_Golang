package wscenternotifyclienttest

import (
	"huobi_futures_Golang/sdk/coinfutures/ws"
	"huobi_futures_Golang/sdk/coinfutures/ws/response/centernotify"
	"huobi_futures_Golang/sdk/log"
	"huobi_futures_Golang/test/config"
	"time"
)

func RunAllExamples() {
	Heartbeat()
}
func Heartbeat() {
	var wscnfClient = new(ws.WSCenterNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wscnfClient.SubHeartbeat("*", func(m *centernotify.SubHeartbeatResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wscnfClient.UnsubHeartbeat("*", "")
	time.Sleep(time.Duration(500) * time.Second)
}
