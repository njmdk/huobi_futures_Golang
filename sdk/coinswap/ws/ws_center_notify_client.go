package ws

import (
	"encoding/json"
	"fmt"
	"github.com/njmdk/huobi_futures_Golang/sdk/coinswap"
	"github.com/njmdk/huobi_futures_Golang/sdk/coinswap/ws/response/centernotify"
	"github.com/njmdk/huobi_futures_Golang/sdk/wsbase"
	"reflect"
)

type WSCenterNotifyClient struct {
	WebSocketOp
}

func (wsNf *WSCenterNotifyClient) Init(accessKey string, secretKey string, host string) *WSCenterNotifyClient {
	if host == "" {
		host = coinswap.COIN_SWAP_DEFAULT_HOST
	}
	wsNf.open("/center-notification", host, accessKey, secretKey, true)
	return wsNf
}

// -------------------------------------------------------------
// heartbeat start
type OnSubHeartbeatResponse func(*centernotify.SubHeartbeatResponse)

func (wsNf *WSCenterNotifyClient) SubHeartbeat(service string, callbackFun OnSubHeartbeatResponse, cid string) {
	if cid == "" {
		cid = coinswap.DEFAULT_CID
	}
	ch := fmt.Sprintf("public.%s.heartbeat", service)
	opData := wsbase.WSOpData{Op: "sub", Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.sub(jdata, ch, callbackFun, reflect.TypeOf(centernotify.SubHeartbeatResponse{}))
}

func (wsNf *WSCenterNotifyClient) UnsubHeartbeat(service string, cid string) {
	if cid == "" {
		cid = coinswap.DEFAULT_CID
	}

	ch := fmt.Sprintf("public.%s.heartbeat", service)
	opData := wsbase.WSOpData{Op: "unsub", Cid: cid, Topic: ch}
	jdata, _ := json.Marshal(opData)

	wsNf.unsub(jdata, ch)
}

// heartbeat end
//-------------------------------------------------------------
