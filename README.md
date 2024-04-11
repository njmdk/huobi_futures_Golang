# Huobi Golang SDK For Contracts v3

This is Huobi Golang SDK v3, you can import to your project and use this SDK to query all market data, trading and manage your account. The SDK supports RESTful API invoking, and subscribing the market, account and order update from the WebSocket connection.

If you already use SDK v1 or v2, it is strongly suggested migrate to v3 as we refactor the implementation to make it simpler and easy to maintain. The SDK v3 is completely consistent with the API documentation of the new HTX open platform. Compared to SDK versions v1 and v2, due to changes in parameters of many interfaces, in order to match the latest interface parameter situation, v3 version has made adjustments to parameters of more than 80 interfaces to ensure that requests can be correctly initiated and accurate response data can be obtained. Meanwhile, the v3 version has added over 130 new interfaces available for use, greatly expanding the number of available interfaces.  We will stop the maintenance of v2 in the near future.

## Table of Contents

- [Quick Start](#Quick-Start)

- [Usage](#Usage)

  - [Configuration](#Configuration)
  - [Folder structure](#Folder-Structure)
  - [Client](#Client)
  - [Response](#Response)

- [Request examples](#Request-examples)
  - [Market data](#Market-data)

- [Subscription examples](#Subscription-examples)
  - [Subscribe trade update](#Subscribe-trade-update)



## Quick Start

You can import the source code from remote or local what shuold be dowload first.

The package **sdk** is core code source as SDK API.
The package **test** is a test example what you can find usage of each API interface in.

If you want to create your own application, you can follow below steps:

* Create a client (xxxClient/WSxxxClient) instance.
* Call the method of client instance.

```go
// RESTful api
acClient := restful.AccountClient{}
acClient.Init("AccessKey", "SecretKey", "")

data := make(chan account.GetAccountAssetsResponse)
go acClient.GetAccountAssetsAsync(data, "BTC-USDT", 0)
x, ok := <-data
if !ok || x.Status != "ok" {
	// fail
} else {
	// success
}

// websocker api
wsmkClient := new(ws.WSMarketClient).Init("")
wsmkClient.SubKLine("BTC-USDT", "15min", func(m *market.SubKLineResponse) {
	// show response data with *m
}, "")
time.Sleep(time.Duration(10) * time.Second)
wsmkClient.UnsubKLine("BTC-USDT", "15min", "")
time.Sleep(time.Duration(10) * time.Second)
```

## Usage

### Configuration

The client Init function must set AccessKey/SecretKey two value when you access private data. And it not need to set AccessKey/SecretKey value when you access public data such as market data.

if you want to access private data, you need below additional steps:

1. Create an **API Key** first from Huobi official website
2. Create **key.go** into your **config** folder (package). The purpose of this file is to prevent submitting SecretKey into repository by accident, so this file is already added in the *.gitignore* file.

3. Assign your secret key to string *SecretKey* and Assign your access key to string *AccessKey*

```go
// key.go file
package config

// replace with your API SecretKey
var SecretKey = "xxxx-xxxx-xxxx-xxxx"

// replace with your API AccessKey
var AccessKey = "xxxx-xxxx-xxxx-xxxx"
```

If you don't need to access private data, you can ignore the secret key.

### Folder Structure

This is the folder and namespace structure of SDK source code and the description

- **sdk**: The SDK API package
  - **linearswap**: the linear swap api src inclue RESTful and websocket
  - **coinfutures**: the coin futures api src inclue RESTful and websocket
  - **coinswap**: the coin swap api src inclue RESTful and websocket
  - **requestbuilder**: Responsible to build the request with the signature
  - **log**: The internal logger interface and implementations
  - **wsbase**: The websocket data model
- **test**: The test example package
  - **xxx_test**: The golang test example

You can find all demo example in xxx_test to get/sub contract private/public data, the request data of the example should be filled in according to your actual situation.

### Client

In this SDK, the client is the object to access the Huobi API. All the client are listed in below table. Each client is very small and simple.

| Module      | Access Type | Client               | Privacy        | Data Category             |
| ----------- | ----------- | -------------------- | -------------- | ------------------------- |
| LinearSwap  | RESTful     | AccountClient        | Private        | account info              |
|             |             | CommonClient         | Private        | common info               |
|             |             | MarketClient         | Public         | market info               |
|             |             | OrderClient          | Private        | about order               |
|             |             | TransferClient       | Private        | transfer assets           |
|             |             | TriggerOrderClient   | Private        | about trigger order       |
|             |             | UnifiedAccountClient | Private        | unified account info      |
|             | Websocket   | WSIndexClinet        | Public         | index info                |
|             |             | WSMarketClinet       | Public         | market info               |
|             |             | WSNotifyClinet       | Public/Private | market info/ account info |
|             |             | WSCenterNotifyClient | Public         | heartbeat info            |
| CoinFutures | RestFul     | AccountClient        | Private        | account info              |
|             |             | CommonClient         | Private        | common info               |
|             |             | MarketClient         | Public         | market info               |
|             |             | OrderClient          | Private        | about order               |
|             |             | TransferClient       | Private        | transfer assets           |
|             |             | TriggerOrderClient   | Private        | about trigger order       |
|             | Websocket   | WSIndexClinet        | Public         | index info                |
|             |             | WSMarketClinet       | Public         | market info               |
|             |             | WSNotifyClinet       | Public/Private | market info/ account info |
|             |             | WSCenterNotifyClient | Public         | heartbeat info            |
| CoinSwap    | RestFul     | AccountClient        | Private        | account info              |
|             |             | CommonClient         | Private        | common info               |
|             |             | MarketClient         | Public         | market info               |
|             |             | OrderClient          | Private        | about order               |
|             |             | TransferClient       | Private        | transfer assets           |
|             |             | TriggerOrderClient   | Private        | about trigger order       |
|             | Websocket   | WSIndexClinet        | Public         | index info                |
|             |             | WSMarketClinet       | Public         | market info               |
|             |             | WSNotifyClinet       | Public/Private | market info/ account info |
|             |             | WSCenterNotifyClient | Public         | heartbeat info            |

#### Public vs. Private

There are two types of privacy that is correspondent with privacy of API:

**Public client**: It invokes public API to get public data (Common data and Market data), therefore you can create a new instance without applying an API Key.

```go
mkClient := restful.MarketClient{}
mkClient.Init("")
```

**Private client**: It invokes private API to access private data, you need to follow the API document to apply an API Key first, and pass the API Key to the Init.

```go
acClient := restful.AccountClient{}
acClient.Init("AccessKey", "SecretKey", "")
```

The API key is used for authentication. If the authentication cannot pass, the invoking of private interface will fail.

#### Rest vs. WebSocket

**Rest Client**: It invokes Rest API and get once-off response.

```go
acClient := restful.AccountClient{}
acClient.Init("AccessKey", "SecretKey", "")

data := make(chan account.GetAccountAssetsResponse)
go acClient.GetAccountAssetsAsync(data, "BTC-USDT", 0)
x, ok := <-data
if !ok || x.Status != "ok" {
	// fail
} else {
	// success
}
```

**WebSocket Client**: It establishes WebSocket connection with server and data will be pushed from server actively. There are two types of method for WebSocket client:

- Request method: The method name starts with "Req", it will receive the once-off data after sending the request.
- Subscription: The method name starts with "Sub", it will receive update after sending the subscription.

```go
wsmkClient := new(ws.WSMarketClient).Init("")

// Reqxxx
wsmkClient.ReqKLine("BTC-USDT", "1min", func(m *market.ReqKLineResponse) {
	// show response data with *m
}, 1604395758, 1604396758, "")

// Subxxx
wsmkClient.SubKLine("BTC-USDT", "15min", func(m *market.SubKLineResponse) {
	// show response data with *m
}, "")
time.Sleep(time.Duration(10) * time.Second)
wsmkClient.UnsubKLine("BTC-USDT", "15min", "")
time.Sleep(time.Duration(10) * time.Second)
```

#### Custom host

Each client Init support an optional host parameter, by default it is "api.hbdm.com". If you need to use different host, you can specify the custom host.

```go
acClient := restful.AccountClient{}
acClient.Init("AccessKey", "SecretKey", "Host")

wsmkClient := new(ws.WSMarketClient).Init("Host")
```

### Response

All response data are defined as follow:
- **huobi_futures_Golang.sdk.linearswap.restful.response**: all linearswap RESTful response data
- **huobi_futures_Golang.sdk.linearswap.ws.response**: all linearswap websockt response data
- **huobi_futures_Golang.sdk.coinfutures.restful.response**: all coinfutures RESTful response data
- **huobi_futures_Golang.sdk.linearswap.ws.response**: all coinfutures websockt response data
- **huobi_futures_Golang.sdk.coinswap.restful.response**: all coinswap RESTful response data
- **huobi_futures_Golang.sdk.coinswap.ws.response**: all coinswap websockt response data

## Request Examples
### Market data
```go
mkClient := restful.MarketClient{}
mkClient.Init("")

data := make(chan market.GetContractInfoResponse)

go mkClient.GetContractInfoAsync(data, "")
x, ok := <-data
if !ok || x.Status != "ok" {
	// fail
} else {
	// show data with x
}
```

## Subscription Examples
### Subscribe trade update
```go
wsnfClient := new(ws.WSNotifyClient).Init("AccessKey", "SecretKey", "")

wsnfClient.SubOrders("*", func(m *notify.SubOrdersResponse) {
	// show dat with *m
}, "")
time.Sleep(time.Duration(50) * time.Second)
wsnfClient.UnsubOrders("*", "")
time.Sleep(time.Duration(50) * time.Second)
```
