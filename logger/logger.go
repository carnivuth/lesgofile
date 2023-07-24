package logger

var Log Logger = CreateLogger(10)

type Logger struct {
	listeners []func(string)
}

func CreateLogger(listenersQueueSize int) Logger {
	var logger Logger
	logger.listeners = make([]func(string), listenersQueueSize)
	return logger
}

func AddListener(logger Logger, callback func(string)) {
	var i int
	for i = 0; i < len(logger.listeners); i++ {
		if logger.listeners[i] == nil {
			logger.listeners[i] = callback
			return
		}
	}
}
func NotifyListeners(callback func(string), message string) {
	go callback(message)

}
func NotifyAll(logger Logger, message string) {
	for _, listener := range logger.listeners {
		if listener != nil {
			NotifyListeners(listener, message)

		}
	}
}

func Emit(logger Logger, message string) {
	NotifyAll(logger, message)
}
