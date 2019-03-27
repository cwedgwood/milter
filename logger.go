package milter

// Logger is a interface to inject a custom logger
type Logger interface {
	Printf(format string, v ...interface{})
}
