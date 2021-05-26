package impl

import (
	"github.com/sangx2/ebest-sdk/wrapper"
)

type Session struct {
	LoginChan      chan wrapper.XaSessionLogin
	LogoutChan     chan bool
	DisConnectChan chan bool
}

func NewSession() *Session {
	return &Session{
		LoginChan:      make(chan wrapper.XaSessionLogin, 1),
		LogoutChan:     make(chan bool, 1),
		DisConnectChan: make(chan bool, 1),
	}
}

func (s Session) GetSessionLoginChan() chan wrapper.XaSessionLogin {
	return s.LoginChan
}

func (s Session) ReceivedLogin(ew *wrapper.EBestWrapper, x wrapper.XaSessionLogin) {
	s.LoginChan <- x
}

func (s Session) ReceivedLogout(ew *wrapper.EBestWrapper) {
	s.LogoutChan <- true
}

func (s Session) ReceivedDisConnect(ew *wrapper.EBestWrapper) {
	s.DisConnectChan <- true
}
