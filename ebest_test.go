package ebest

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/sangx2/ebest/res"
)

func TestEbest(t *testing.T) {
	// fix me
	resPath := "C:\\eBEST\\xingAPI\\Res\\"
	eBest := NewEbest("", "", "", SERVER_VIRTUAL, PORT, resPath)
	if eBest == nil {
		t.Error("NewEbest error")
		return
	}

	if e := eBest.Connect(); e != nil {
		t.Errorf("Login error : %s", e)
		return
	}

	if e := eBest.Login(); e != nil {
		t.Errorf("Login error : %s", e)
		return
	}

	accounts := eBest.GetAccountList()
	t.Logf("GetAccountList : %s", accounts)
	t.Logf("GetAccountName : %s", eBest.GetAccountName(accounts[0]))
	t.Logf("GetAccountDetailName : %s", eBest.GetAccountDetailName(accounts[0]))
	t.Logf("GetAccountNickName : %s", eBest.GetAccountNickName(accounts[0]))

	// Query
	t0424 := eBest.CreateQuery(T0424)
	if t0424 == nil {
		t.Error("NewQuery(\"주식 잔고2\") error")
		return
	}

	if e := t0424.SetInBlock(res.T0424InBlock{Accno: accounts[0]}); e != nil {
		t.Errorf("Query.SetInblock error : %s", e)
		return
	}

	if ret := t0424.Request(); ret < 0 {
		t.Errorf("Query.Request error : %s", eBest.GetErrorMessage(ret))
		return
	}

	if msg, e := t0424.GetReceiveMessage(); e != nil {
		t.Errorf("Query.GetReceiveMessage error : %s", e)
		return
	} else {
		t.Logf("Query.GetReceiveMessage : %s", msg)
		t.Logf("Query.GetReceiveData : %s", <-t0424.GetReceiveDataChan())
	}

	outBlocks := t0424.GetOutBlocks()
	if t0424OutBlock, ok := outBlocks[0].(res.T0424OutBlock); !ok {
		t.Errorf("T0424Outblock error %+v", t0424OutBlock)
		return
	} else {
		t.Logf("T0424Outblock : %+v", t0424OutBlock)
	}

	if t0424OutBlock1s, ok := outBlocks[1].([]res.T0424OutBlock1); !ok {
		t.Errorf("T0424OutBlock1 error")
		return
	} else {
		t.Logf("T0424Outblock1 : %+v", t0424OutBlock1s)
	}

	// Real
	news := eBest.CreateReal(NWS)
	if news == nil {
		t.Error("NewReal(\"실시간 뉴스 제목 패킷\") error")
		return
	}
	t.Log("NewReal (\"실시간 뉴스 제목 패킷\")")

	if e := news.SetInBlock(res.NWSInBlock{NWcode: "NWS001"}); e != nil {
		t.Errorf("Real.SetInBlock error : %s", e)
		return
	}

	news.Start()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case data := <-news.GetReceivedRealDataChan():
		t.Logf("Real.GetReceiveRealData : %s", data)
		if newsOut, ok := news.GetOutBlock().(res.NWSOutBlock); !ok {
			t.Errorf("NWSOutBlock error")
			return
		} else {
			t.Logf("NWSOutBlock : %+v", newsOut)
		}
	case <-interruptChan:
		break
	}

	t0424.Close()
	news.Close()

	eBest.Disconnect()
}
