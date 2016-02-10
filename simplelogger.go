package simplelogger

import (
	"bytes"
	"fmt"
	"io"
	"sync"
	"time"
)

type SimpleLogger struct {
	lock sync.Mutex
	out  io.Writer
	err  io.Writer
	buf  []byte

	criticalPrefixFmt string
	errorPrefixFmt    string
	debugPrefixFmt    string
}

func New(name string, out io.Writer, err io.Writer) *SimpleLogger {
	return &SimpleLogger{
		out:               out,
		err:               err,
		criticalPrefixFmt: fmt.Sprintf("2006-01-02 15:05:05 CRITICAL %s ", name),
		buf:               make([]byte, 1024),
		errorPrefixFmt:    fmt.Sprintf("2006-01-02 15:05:05 ERROR %s ", name),
		debugPrefixFmt:    fmt.Sprintf("2006-01-02 15:05:05 DEBUG %s ", name),
	}
}

func (s *SimpleLogger) CriticalPrintf(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf(s.buf, s.err, s.criticalPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) ErrorPrintf(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf(s.buf, s.err, s.errorPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf(s.buf, s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func simpleLoggerPrintf(buf []byte, out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	buf = time.Now().AppendFormat(buf[:0], prefixFmt)
	bufr := bytes.NewBuffer(buf)
	fmt.Fprintf(bufr, msg, args...)
	buf = bufr.Bytes()
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
}
