package main

import (
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/accountclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/commonclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/marketclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/orderclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/transferclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/triggerorderclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/unifiedaccountclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/wscenternotifyclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/wsindexclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/wsmarketclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/linearswap_test/wsnotifyclienttest"
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
