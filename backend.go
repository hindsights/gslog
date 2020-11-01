package gslog

var theBackend Backend

type Backend interface {
	GetLogger(name string) Logger
}

func GetLogger(name string) Logger {
	return theBackend.GetLogger(name)
}

func SetBackend(backend Backend) {
	theBackend = backend
}
