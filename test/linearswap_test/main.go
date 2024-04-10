package main

import (
	"huobi_futures_Golang/test/linearswap_test/accountclienttest"
	"huobi_futures_Golang/test/linearswap_test/commonclienttest"
	"huobi_futures_Golang/test/linearswap_test/marketclienttest"
	"huobi_futures_Golang/test/linearswap_test/orderclienttest"
	"huobi_futures_Golang/test/linearswap_test/transferclienttest"
	"huobi_futures_Golang/test/linearswap_test/triggerorderclienttest"
	"huobi_futures_Golang/test/linearswap_test/unifiedaccountclienttest"
	"huobi_futures_Golang/test/linearswap_test/wscenternotifyclienttest"
	"huobi_futures_Golang/test/linearswap_test/wsindexclienttest"
	"huobi_futures_Golang/test/linearswap_test/wsmarketclienttest"
	"huobi_futures_Golang/test/linearswap_test/wsnotifyclienttest"
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
	unifiedaccountclienttest.RunAllExamples()
	wsindexclienttest.RunAllExamples()
	wsmarketclienttest.RunAllExamples()
	wsnotifyclienttest.RunAllExamples()
	wscenternotifyclienttest.RunAllExamples()
}
