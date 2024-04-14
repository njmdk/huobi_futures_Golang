package main

import (
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/accountclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/commonclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/marketclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/orderclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/transferclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/triggerorderclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/wscenternotifyclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/wsindexclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/wsmarketclienttest"
	"github.com/HuobiRDCenter/huobi_futures_Golang/test/coinfutures_test/wsnotifyclienttest"
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
