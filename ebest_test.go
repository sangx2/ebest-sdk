package ebest

import (
	"os"
	"os/signal"
	"syscall"
	"testing"

	"github.com/sangx2/ebest/model"
)

func TestEbest(t *testing.T) {
	// fix me
	resPath := "C:\\eBEST\\xingAPI\\Res\\"
	ebest := NewEbest("", "", "", SERVER_VIRTUAL, PORT, nil)
	if ebest == nil {
		t.Error("NewEbest error")
	}

	ebest.Connect()

	if e := ebest.Login(); e != nil {
		t.Errorf("Login error : %s", e)
	}

	accounts := ebest.GetAccountList()
	t.Logf("GetAccountList : %s", accounts)
	t.Logf("GetAccountName : %s", ebest.GetAccountName(accounts[0]))
	t.Logf("GetAccountDetailName : %s", ebest.GetAccountDetailName(accounts[0]))
	t.Logf("GetAccountNickName : %s", ebest.GetAccountNickName(accounts[0]))

	// Query
	t0424 := NewQuery(resPath, T0424)
	if t0424 == nil {
		t.Error("NewQuery(\"주식 잔고2\") error")
	}

	if e := t0424.SetInBlock(model.T0424InBlock{Accno: accounts[0]}, nil); e != nil {
		t.Errorf("Query.SetInblock error : %s", e)
	}

	if ret := t0424.Request(); ret < 0 {
		t.Errorf("Query.Request error : %s", ebest.GetErrorMessage(ret))
	}

	if msg, e := t0424.GetReceiveMessage(); e != nil {
		t.Errorf("Query.GetReceiveMessage error : %s", e)
	} else {
		t.Logf("Query.GetReceiveMessage : %s", msg)
		t.Logf("Query.GetReceiveData : %s", <-t0424.GetReceiveDataChan())
	}

	out, out1, _, _, _, _ := t0424.GetOutBlock()
	t0424Outblock, ok := out.(model.T0424OutBlock)
	if !ok {
		t.Errorf("T0424Outblock error %+v", t0424Outblock)
	} else {
		t.Logf("T0424Outblock : %+v", t0424Outblock)
	}

	t0424Outblock1s, ok := out1.([]model.T0424OutBlock1)
	if !ok {
		t.Errorf("T0424OutBlock1 error")
	} else {
		t.Logf("T0424Outblock1 : %+v", t0424Outblock1s)
	}

	t0424.Close()

	// Real
	news := NewReal(resPath, NWS)
	if news == nil {
		t.Error("NewReal(\"실시간 뉴스 제목 패킷\") error")
	}
	t.Log("NewReal (\"실시간 뉴스 제목 패킷\")")

	e := news.SetInBlock(model.NWSInBlock{NWcode: "NWS001"})
	if e != nil {
		t.Errorf("Real.SetInBlock error : %s", e)
	}

	news.Start()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case data := <-news.GetReceivedRealDataChan():
		t.Logf("Real.GetReceiveRealData : %s", data)

		newsOut, ok := news.GetOutBlock().(model.NWSOutBlock)
		if !ok {
			t.Errorf("NWSOutBlock error")
		} else {
			t.Logf("NWSOutBlock : %+v", newsOut)
		}
	case <-interruptChan:
		break
	}

	news.Stop("")

	news.Close()

	ebest.Disconnect()
}
