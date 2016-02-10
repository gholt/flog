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
	simpleLoggerPrintf(s.err, s.criticalPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) ErrorPrintf(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf(s.err, s.errorPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf(s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf2(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf2(s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf3(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf3(s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf4(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf4(s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf5(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf5(s.buf, s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf6(msg string, args ...interface{}) {
	s.lock.Lock()
	simpleLoggerPrintf6(s.buf, s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf7(msg string, args ...interface{}) {
	s.lock.Lock()
	s.simpleLoggerPrintf7(s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func (s *SimpleLogger) DebugPrintf8(msg string, args ...interface{}) {
	s.lock.Lock()
	s.simpleLoggerPrintf8(s.out, s.debugPrefixFmt, msg, args...)
	s.lock.Unlock()
}

func simpleLoggerPrintf(out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	out.Write([]byte(time.Now().Format(prefixFmt)))
	m := fmt.Sprintf(msg, args...)
	out.Write([]byte(m))
	if len(m) < 1 || m[len(m)-1] != '\n' {
		out.Write([]byte("\n"))
	}
}

func simpleLoggerPrintf2(out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	m := time.Now().Format(prefixFmt) + fmt.Sprintf(msg, args...)
	if len(m) == 0 || m[len(m)-1] != '\n' {
		m += "\n"
	}
	out.Write([]byte(m))
}

func simpleLoggerPrintf3(out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	m := time.Now().Format(prefixFmt) + fmt.Sprintf(msg, args...)
	out.Write([]byte(m))
	if len(m) == 0 || m[len(m)-1] != '\n' {
		out.Write([]byte("\n"))
	}
}

func simpleLoggerPrintf4(out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	out.Write([]byte(time.Now().Format(prefixFmt)))
	m := fmt.Sprintf(msg, args...)
	if len(m) == 0 || m[len(m)-1] != '\n' {
		m += "\n"
	}
	out.Write([]byte(m))
}

func simpleLoggerPrintf5(buf []byte, out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	buf = time.Now().AppendFormat(buf[:0], prefixFmt)
	buf = append(buf, []byte(fmt.Sprintf(msg, args...))...)
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
}

func simpleLoggerPrintf6(buf []byte, out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	buf = time.Now().AppendFormat(buf[:0], prefixFmt)
	bufr := bytes.NewBuffer(buf)
	fmt.Fprintf(bufr, msg, args...)
	buf = bufr.Bytes()
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
}

func (s *SimpleLogger) simpleLoggerPrintf7(out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	buf := time.Now().AppendFormat(s.buf[:0], prefixFmt)
	bufr := bytes.NewBuffer(buf)
	fmt.Fprintf(bufr, msg, args...)
	buf = bufr.Bytes()
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
	s.buf = buf
}

func (s *SimpleLogger) simpleLoggerPrintf8(out io.Writer, prefixFmt string, msg string, args ...interface{}) {
	buf := time.Now().AppendFormat(s.buf[:0], prefixFmt)
	buf = append(buf, []byte(fmt.Sprintf(msg, args...))...)
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
	s.buf = buf
}
