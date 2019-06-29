package interfaces

import "github.com/sangx2/ebest/wrapper"

type Session interface {
	// event
	GetSessionLoginChan() chan wrapper.XaSessionLogin

	// callback func
	ReceivedLogin(*wrapper.Ebest, wrapper.XaSessionLogin)
	ReceivedLogout(*wrapper.Ebest)
	ReceivedDisConnect(*wrapper.Ebest)
}
