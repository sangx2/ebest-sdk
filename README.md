# 소개

ebest 는 이베스트 투자증권 COM버전을 이용한 라이브러리 패키지.

이를 이용한 api 서버는 [ebest](https://github.com/sangx2/ebest) 에 구현

## 주의사항

- golang은 windows x86(x86-64는 동작하지 않음) 플랫폼 파일 설치

## 구현 정보

항목이 추가될 경우 impl, res 폴더 내 해당 파일 추가
- interfaces : impl 폴더 내의 trade(query/real) 구현을 위한 인터페이스
- impl : trade 구현체
- res : trade 에 사용되는 inBlock/outBlock 구조체 정의


## 참고사항

- COM에서는 이벤트를 맺은 윈도우 쓰레드에 메시지를 전달하기 때문에 go-scheduler를 통하지 않고 직접 쓰레드를 생성하는 API를 구현해 개발