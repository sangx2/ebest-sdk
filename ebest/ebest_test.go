package ebest

import (
	"github.com/sangx2/ebest-sdk/impl"
	"github.com/sangx2/ebest-sdk/res"
	"log"
	"sync"
	"testing"
	"time"
)

// fix me
var (
	Id         = ""
	Passwd     = ""
	CertPasswd = ""
	ResPath    = "C:\\eBEST\\xingAPI\\Res\\"
)

func TestEbest(t *testing.T) {
	var err error

	e := NewEBest(Id, Passwd, CertPasswd, ServerVirtual, Port, ResPath)
	if e == nil {
		t.Fatalf("NewEbest is nil")
	}

	if err = e.Connect(); err != nil {
		t.Fatalf("Connect :%s", err.Error())
	}
	defer e.Disconnect()

	if err = e.Login(); err != nil {
		t.Fatalf("Login :%s", err.Error())
	}

	accounts := e.GetAccountList()
	for _, account := range accounts {
		name := e.GetAccountName(account)
		detailName := e.GetAccountDetailName(account)
		nickName := e.GetAccountNickName(account)

		log.Printf("account:%s, name:%s, detailName:%s, nickName:%s\n", account, name, detailName, nickName)
	}

	// query
	t1101 := NewQuery(ResPath, impl.NewT1101())
	if err = t1101.SetInBlock(res.T1101InBlock{Shcode: "005930"}); err != nil {
		t.Fatalf("SetInBlock:%s", err.Error())
	}
	defer t1101.Close()

	ret := t1101.Request(false)
	t.Logf("Request:%d\n", ret)

	msg, err := t1101.GetReceiveMessage()
	if err != nil {
		t.Fatalf("GetReceiveMessage:%s", err.Error())
	}
	t.Logf("GetReceiveMessage:%s\n", msg)

	dataChan := <-t1101.GetReceiveDataChan()
	t.Logf("GetReceiveDataChan:%s\n", dataChan)

	t8436OutBlocks := t1101.GetOutBlocks()
	t.Logf("t8436OutBlock:%+v\n", t8436OutBlocks)

	// real
	nws := NewReal(ResPath, impl.NewNWS())
	if nws == nil {
		t.Fatalf("nws is nil")
	}
	defer nws.Close()

	if err := nws.SetInBlock(res.NWSInBlock{"NWS001"}); err != nil {
		t.Fatalf("SetInBlock:%s", err.Error())
	}

	nws.Start()
	defer nws.Stop("NWS")

	wg := sync.WaitGroup{}
	doneChan := make(chan bool, 1)
	go func(doneChan chan bool, wg *sync.WaitGroup) {
		wg.Add(1)
		for {
			select {
			case <-nws.GetReceivedRealDataChan():
				log.Printf("nws:%+v\n", nws.GetOutBlock().(res.NWSOutBlock))
			case <-doneChan:
				wg.Done()
				return
			}
		}
	}(doneChan, &wg)

	time.Sleep(time.Second * 5)

	doneChan <- true
	close(doneChan)

	wg.Wait()
}
