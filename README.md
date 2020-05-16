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

ebest_test.go 코드 참조