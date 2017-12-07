package noop

type Logger struct{}

func NewErrorLogger() *Logger { return &Logger{} }

func (l *Logger) Capture(err error, opts map[string]interface{}) error {
	return nil
}
func (l *Logger) Close()                             {}
func (l *Logger) IgnoreErrors(_ ...func(error) bool) {}
