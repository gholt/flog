package simplelogger

import (
	"fmt"
	"io"
	"sync"
	"time"
)

type SimpleLogger struct {
	outLock sync.Mutex
	out     io.Writer

	errLock sync.Mutex
	err     io.Writer

	criticalPrefixFmt string
	errorPrefixFmt    string
	debugPrefixFmt    string
}

func New(name string, out io.Writer, err io.Writer) *SimpleLogger {
	return &SimpleLogger{
		out:               out,
		err:               err,
		criticalPrefixFmt: fmt.Sprintf("2006-01-02 15:05:05 CRITICAL %s ", name),
		errorPrefixFmt:    fmt.Sprintf("2006-01-02 15:05:05 ERROR %s ", name),
		debugPrefixFmt:    fmt.Sprintf("2006-01-02 15:05:05 DEBUG %s ", name),
	}
}

func (s *SimpleLogger) CriticalPrintf(msg string, args ...interface{}) {
	s.errLock.Lock()
	simpleLoggerPrintf(s.err, s.criticalPrefixFmt, msg, args...)
	s.errLock.Unlock()
}

func (s *SimpleLogger) ErrorPrintf(msg string, args ...interface{}) {
	s.errLock.Lock()
	simpleLoggerPrintf(s.err, s.errorPrefixFmt, msg, args...)
	s.errLock.Unlock()
}

func (s *SimpleLogger) DebugPrintf(msg string, args ...interface{}) {
	s.outLock.Lock()
	simpleLoggerPrintf(s.out, s.debugPrefixFmt, msg, args...)
	s.outLock.Unlock()
}

func simpleLoggerPrintf(out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	out.Write([]byte(time.Now().Format(prefixFmt)))
	m := fmt.Sprintf(msg, args...)
	out.Write([]byte(m))
	if len(m) < 1 || m[len(m)-1] != '\n' {
		out.Write([]byte("\n"))
	}
}
