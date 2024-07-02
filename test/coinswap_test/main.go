package main

import (
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/accountclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/commonclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/marketclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/orderclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/transferclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/triggerorderclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/wscenternotifyclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/wsindexclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/wsmarketclienttest"
	"github.com/njmdk/huobi_futures_Golang/test/coinswap_test/wsnotifyclienttest"
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
