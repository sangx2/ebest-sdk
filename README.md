# 소개

ebest 는 이베스트 투자증권 COM버전을 이용한 라이브러리 패키지 입니다. 

내부적으로 COM을 위한 "github.com/go-ole/go-ole"와 ORM을 위한 "github.com/jinzhu/gorm" 패키지를 사용합니다.

### 주의사항

1. golang은 windows x86(x86-64는 동작하지 않음)에 해당하는 파일을 설치해야 합니다.
2. 요청 수 제한에 대한 처리는 하지 않습니다.
3. COM에서는 이벤트를 맺은 윈도우 쓰레드에 메시지를 전달하기 때문에 go-scheduler를 통하지 않고 직접 쓰레드를 생성하는 API를 구현해 개발했습니다.
4. 추가적인 RES 파일은 점차 추가할 예정입니다.

### 설치

```bash
go get -u github.com/sangx2/ebest
```

## Getting started

### 가상 서버 접속 및 테스트

```go
package main

import (
	"os"
	"os/signal"
	"syscall"

    "github.com/sangx2/ebest"
	"github.com/sangx2/ebest/model"
)

func TestEbest() {
	// fix me
	resPath := "C:\\eBEST\\xingAPI\\Res\\"
	ebest := NewEbest("", "", "", SERVER_VIRTUAL, PORT, nil)
	if ebest == nil {
		fmt.Println("NewEbest error")
	}

	ebest.Connect()

	if e := ebest.Login(); e != nil {
		fmt.Printf("Login error : %s", e)
	}

	accounts := ebest.GetAccountList()
	fmt.Printf("GetAccountList : %s", accounts)
	fmt.Printf("GetAccountName : %s", ebest.GetAccountName(accounts[0]))
	fmt.Printf("GetAccountDetailName : %s", ebest.GetAccountDetailName(accounts[0]))
	fmt.Printf("GetAccountNickName : %s", ebest.GetAccountNickName(accounts[0]))

	// Query
	t0424 := NewQuery(resPath, T0424)
	if t0424 == nil {
		fmt.Println("NewQuery(\"주식 잔고2\") error")
	}

	if e := t0424.SetInBlock(model.T0424InBlock{Accno: accounts[0]}, nil); e != nil {
		fmt.Printf("Query.SetInblock error : %s", e)
	}

	if ret := t0424.Request(); ret < 0 {
		fmt.Printf("Query.Request error : %s", ebest.GetErrorMessage(ret))
	}

	if msg, e := t0424.GetReceiveMessage(); e != nil {
		fmt.Printf("Query.GetReceiveMessage error : %s", e)
	} else {
		fmt.Printf("Query.GetReceiveMessage : %s", msg)
		fmt.Printf("Query.GetReceiveData : %s", <-t0424.GetReceiveDataChan())
	}

	out, out1, _, _, _, _ := t0424.GetOutBlock()
	t0424Outblock, ok := out.(model.T0424OutBlock)
	if !ok {
		fmt.Printf("T0424Outblock error %+v", t0424Outblock)
	} else {
		fmt.Printf("T0424Outblock : %+v", t0424Outblock)
	}

	t0424Outblock1s, ok := out1.([]model.T0424OutBlock1)
	if !ok {
		fmt.Println("T0424OutBlock1 error")
	} else {
		fmt.Printf("T0424Outblock1 : %+v", t0424Outblock1s)
	}

	t0424.Close()

	// Real
	news := NewReal(resPath, NWS)
	if news == nil {
		fmt.Println("NewReal(\"실시간 뉴스 제목 패킷\") error")
	}
	fmt.Println("NewReal (\"실시간 뉴스 제목 패킷\")")

	e := news.SetInBlock(model.NWSInBlock{NWcode: "NWS001"})
	if e != nil {
		fmt.Printf("Real.SetInBlock error : %s", e)
	}

	news.Start()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	select {
	case data := <-news.GetReceivedRealDataChan():
		fmt.Printf("Real.GetReceiveRealData : %s", data)

		newsOut, ok := news.GetOutBlock().(model.NWSOutBlock)
		if !ok {
			fmt.Println("NWSOutBlock error")
		} else {
			fmt.Printf("NWSOutBlock : %+v", newsOut)
		}
	case <-interruptChan:
		break
	}

	news.Stop("")

	news.Close()

	ebest.Disconnect()
}

```