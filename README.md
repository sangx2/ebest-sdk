# 소개

ebest 는 이베스트 투자증권 COM버전을 이용한 라이브러리 패키지

## 주의사항

- golang은 windows x86(x86-64는 동작하지 않음) 플랫폼 파일 설치

## 참고사항
- COM에서는 이벤트를 맺은 윈도우 쓰레드에 메시지를 전달하기 때문에 go-scheduler를 통하지 않고 직접 쓰레드를 생성하는 API를 구현해 개발
- 서버는 [ebest](https://github.com/sangx2/ebest)에 구현 중
