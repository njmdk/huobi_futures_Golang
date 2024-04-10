package main

import (
	"huobi_futures_Golang/test/coinfutures_test/accountclienttest"
	"huobi_futures_Golang/test/coinfutures_test/commonclienttest"
	"huobi_futures_Golang/test/coinfutures_test/marketclienttest"
	"huobi_futures_Golang/test/coinfutures_test/orderclienttest"
	"huobi_futures_Golang/test/coinfutures_test/transferclienttest"
	"huobi_futures_Golang/test/coinfutures_test/triggerorderclienttest"
	"huobi_futures_Golang/test/coinfutures_test/wscenternotifyclienttest"
	"huobi_futures_Golang/test/coinfutures_test/wsindexclienttest"
	"huobi_futures_Golang/test/coinfutures_test/wsmarketclienttest"
	"huobi_futures_Golang/test/coinfutures_test/wsnotifyclienttest"
)

func main() {
	runAll()
}

// Run all examples
func runAll() {
	accountclienttest.RunAllExamples()
	commonclienttest.RunAllExamples()
	marketclienttest.RunAllExamples()
	orderclienttest.RunAllExamples()
	transferclienttest.RunAllExamples()
	triggerorderclienttest.RunAllExamples()
	wsindexclienttest.RunAllExamples()
	wsmarketclienttest.RunAllExamples()
	wsnotifyclienttest.RunAllExamples()
	wscenternotifyclienttest.RunAllExamples()
}
