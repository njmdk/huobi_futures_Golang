package wsnotifyclienttest

import (
	"github.com/njmdk/huobi_futures_Golang/config"
	"github.com/njmdk/huobi_futures_Golang/sdk/coinswap/ws"
	"github.com/njmdk/huobi_futures_Golang/sdk/coinswap/ws/response/notify"
	"github.com/njmdk/huobi_futures_Golang/sdk/log"
	"time"
)

func RunAllExamples() {
	Orders()
	Accounts()
	Positions()
	MathOrders()
	LiquidationOrders()
	FundingRate()
	ContractInfo()
	TriggerOrder()
}

func Orders() {
	var wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wsnfClient.IsolatedSubOrders("*", func(m *notify.SubOrdersResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wsnfClient.IsolatedUnsubOrders("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func Accounts() {
	var wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wsnfClient.IsolatedSubAcounts("*", func(m *notify.SubAccountsResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(50) * time.Second)
	wsnfClient.IsolatedUnsubAccounts("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func Positions() {
	var wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wsnfClient.IsolatedSubPositions("xrp-usdt", func(m *notify.SubPositionsResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(50) * time.Second)
	wsnfClient.IsolatdUnsubPositions("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func MathOrders() {
	var wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wsnfClient.IsolatedSubMatchOrders("*", func(m *notify.SubOrdersResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wsnfClient.IsolatedUnsubMathOrders("*", "")
	time.Sleep(time.Duration(50) * time.Second)
}

func LiquidationOrders() {
	var wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wsnfClient.SubLiquidationOrders("*", func(m *notify.SubLiquidationOrdersResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(500) * time.Second)
	wsnfClient.UnsubLiquidationOrders("*", "")
	time.Sleep(time.Duration(500) * time.Second)
}

func FundingRate() {
	var wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wsnfClient.SubFundingRate("BTC-USDT", func(m *notify.SubFundingRateResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(10) * time.Second)
	wsnfClient.UnsubFundingRate("BTC-USDT", "")
	time.Sleep(time.Duration(10) * time.Second)
}

func ContractInfo() {
	var wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wsnfClient.SubContractInfo("*", func(m *notify.SubContractInfoResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(10) * time.Second)
	wsnfClient.UnsubContractInfo("*", "")
	time.Sleep(time.Duration(10) * time.Second)
}

func TriggerOrder() {
	var wsnfClient = new(ws.WSNotifyClient).Init(config.AccessKey, config.SecretKey, "")
	wsnfClient.IsolatedSubTriggerOrder("*", func(m *notify.SubTriggerOrderResponse) {
		log.Info("%v", *m)
	}, "")
	time.Sleep(time.Duration(100) * time.Second)
	wsnfClient.IsolatedUnsubTriggerOrder("*", "")
	time.Sleep(time.Duration(100) * time.Second)
}
