package interfaces

type Logger interface {
	Info(msg string)
	Error(msg string)
}
