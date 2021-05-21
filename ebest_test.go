package ebestsdk

import (
	"github.com/sangx2/ebestsdk/impl"
	"github.com/sangx2/ebestsdk/res"
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
		log.Fatalf("NewEbest is nil")
	}

	if err = e.Connect(); err != nil {
		log.Fatalf("Connect :%s", err.Error())
	}

	if err = e.Login(); err != nil {
		log.Fatalf("Login :%s", err.Error())
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

	t1101.Close()

	// real
	nws := NewReal(ResPath, impl.NewNWS())
	if nws == nil {
		t.Fatalf("nws is nil")
	}

	if err := nws.SetInBlock(res.NWSInBlock{"NWS001"}); err != nil {
		t.Fatalf("SetInBlock:%s", err.Error())
	}

	nws.Start()

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
	nws.Close()

	doneChan <- true
	close(doneChan)

	wg.Wait()

	e.Disconnect()
}
