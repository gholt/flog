package flog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var Std Flog = New("", os.Stderr, os.Stderr, os.Stderr, os.Stdout, nil)

type Flog interface {
	CriticalPrintf(format string, args ...interface{})
	CriticalPrintln(args ...interface{})
	ErrorPrintf(format string, args ...interface{})
	ErrorPrintln(args ...interface{})
	WarningPrintf(format string, args ...interface{})
	WarningPrintln(args ...interface{})
	InfoPrintf(format string, args ...interface{})
	InfoPrintln(args ...interface{})
	DebugPrintf(format string, args ...interface{})
	DebugPrintln(args ...interface{})
}

type FlogWriters interface {
	Flog
	CriticalWriter() io.Writer
	SetCriticalWriter(w io.Writer) io.Writer
	ErrorWriter() io.Writer
	SetErrorWriter(w io.Writer) io.Writer
	WarningWriter() io.Writer
	SetWarningWriter(w io.Writer) io.Writer
	InfoWriter() io.Writer
	SetInfoWriter(w io.Writer) io.Writer
	DebugWriter() io.Writer
	SetDebugWriter(w io.Writer) io.Writer
}

type flog struct {
	lock           sync.Mutex
	criticalWriter io.Writer
	errorWriter    io.Writer
	warningWriter  io.Writer
	infoWriter     io.Writer
	debugWriter    io.Writer
	buf            []byte

	criticalFmt string
	errorFmt    string
	warningFmt  string
	infoFmt     string
	debugFmt    string
}

func New(name string, criticalWriter io.Writer, errorWriter io.Writer, warningWriter io.Writer, infoWriter io.Writer, debugWriter io.Writer) FlogWriters {
	f := &flog{
		criticalWriter: criticalWriter,
		errorWriter:    errorWriter,
		warningWriter:  warningWriter,
		infoWriter:     infoWriter,
		debugWriter:    debugWriter,
	}
	if name == "" {
		f.criticalFmt = fmt.Sprintf("2006-01-02 15:05:05 CRITICAL ")
		f.errorFmt = fmt.Sprintf("2006-01-02 15:05:05 ERROR ")
		f.warningFmt = fmt.Sprintf("2006-01-02 15:05:05 WARNING ")
		f.infoFmt = fmt.Sprintf("2006-01-02 15:05:05 INFO ")
		f.debugFmt = fmt.Sprintf("2006-01-02 15:05:05 DEBUG ")
	} else {
		f.criticalFmt = fmt.Sprintf("2006-01-02 15:05:05 CRITICAL %s ", name)
		f.errorFmt = fmt.Sprintf("2006-01-02 15:05:05 ERROR %s ", name)
		f.warningFmt = fmt.Sprintf("2006-01-02 15:05:05 WARNING %s ", name)
		f.infoFmt = fmt.Sprintf("2006-01-02 15:05:05 INFO %s ", name)
		f.debugFmt = fmt.Sprintf("2006-01-02 15:05:05 DEBUG %s ", name)
	}
	return f
}

func (f *flog) CriticalWriter() io.Writer {
	f.lock.Lock()
	w := f.criticalWriter
	f.lock.Unlock()
	return w
}

func (f *flog) SetCriticalWriter(w io.Writer) io.Writer {
	f.lock.Lock()
	o := f.criticalWriter
	f.criticalWriter = w
	f.lock.Unlock()
	return o
}

func (f *flog) ErrorWriter() io.Writer {
	f.lock.Lock()
	w := f.errorWriter
	f.lock.Unlock()
	return w
}

func (f *flog) SetErrorWriter(w io.Writer) io.Writer {
	f.lock.Lock()
	o := f.errorWriter
	f.errorWriter = w
	f.lock.Unlock()
	return o
}

func (f *flog) WarningWriter() io.Writer {
	f.lock.Lock()
	w := f.warningWriter
	f.lock.Unlock()
	return w
}

func (f *flog) SetWarningWriter(w io.Writer) io.Writer {
	f.lock.Lock()
	o := f.warningWriter
	f.warningWriter = w
	f.lock.Unlock()
	return o
}

func (f *flog) InfoWriter() io.Writer {
	f.lock.Lock()
	w := f.infoWriter
	f.lock.Unlock()
	return w
}

func (f *flog) SetInfoWriter(w io.Writer) io.Writer {
	f.lock.Lock()
	o := f.infoWriter
	f.infoWriter = w
	f.lock.Unlock()
	return o
}

func (f *flog) DebugWriter() io.Writer {
	f.lock.Lock()
	w := f.debugWriter
	f.lock.Unlock()
	return w
}

func (f *flog) SetDebugWriter(w io.Writer) io.Writer {
	f.lock.Lock()
	o := f.debugWriter
	f.debugWriter = w
	f.lock.Unlock()
	return o
}

func (f *flog) CriticalPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.criticalWriter, f.criticalFmt, format, args...)
	f.lock.Unlock()
}

func (f *flog) ErrorPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.errorWriter, f.errorFmt, format, args...)
	f.lock.Unlock()
}

func (f *flog) WarningPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.warningWriter, f.warningFmt, format, args...)
	f.lock.Unlock()
}

func (f *flog) InfoPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.infoWriter, f.infoFmt, format, args...)
	f.lock.Unlock()
}

func (f *flog) DebugPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.debugWriter, f.debugFmt, format, args...)
	f.lock.Unlock()
}

func (f *flog) CriticalPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.criticalWriter, f.criticalFmt, args...)
	f.lock.Unlock()
}

func (f *flog) ErrorPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.errorWriter, f.errorFmt, args...)
	f.lock.Unlock()
}

func (f *flog) WarningPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.warningWriter, f.warningFmt, args...)
	f.lock.Unlock()
}

func (f *flog) InfoPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.infoWriter, f.infoFmt, args...)
	f.lock.Unlock()
}

func (f *flog) DebugPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.debugWriter, f.debugFmt, args...)
	f.lock.Unlock()
}

func flogPrintf(buf []byte, out io.Writer, prefixFmt string, format string, args ...interface{}) {
	if out == nil {
		return
	}
	buf = time.Now().AppendFormat(buf[:0], prefixFmt)
	bufr := bytes.NewBuffer(buf)
	fmt.Fprintf(bufr, format, args...)
	buf = bufr.Bytes()
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
}

func flogPrintln(buf []byte, out io.Writer, prefixFmt string, args ...interface{}) {
	if out == nil {
		return
	}
	buf = time.Now().AppendFormat(buf[:0], prefixFmt)
	bufr := bytes.NewBuffer(buf)
	fmt.Fprintln(bufr, args...)
	buf = bufr.Bytes()
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
}

type Logger interface {
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	Fatalln(args ...interface{})
	Print(args ...interface{})
	Printf(format string, args ...interface{})
	Println(args ...interface{})
}

type wrapper struct {
	printf  func(format string, args ...interface{})
	println func(args ...interface{})
}

func (w *wrapper) Fatal(args ...interface{}) {
	w.println(args)
	os.Exit(1)
}

func (w *wrapper) Fatalf(format string, args ...interface{}) {
	w.printf(format, args)
	os.Exit(1)
}

func (w *wrapper) Fatalln(args ...interface{}) {
	w.println(args)
	os.Exit(1)
}

func (w *wrapper) Print(args ...interface{}) {
	w.println(args)
}

func (w *wrapper) Printf(format string, args ...interface{}) {
	w.printf(format, args)
}

func (w *wrapper) Println(args ...interface{}) {
	w.println(args)
}

func CriticalLogger(f Flog) Logger {
	return &wrapper{
		printf:  f.CriticalPrintf,
		println: f.CriticalPrintln,
	}
}

func ErrorLogger(f Flog) Logger {
	return &wrapper{
		printf:  f.ErrorPrintf,
		println: f.ErrorPrintln,
	}
}

func WarningLogger(f Flog) Logger {
	return &wrapper{
		printf:  f.WarningPrintf,
		println: f.WarningPrintln,
	}
}

func InfoLogger(f Flog) Logger {
	return &wrapper{
		printf:  f.InfoPrintf,
		println: f.InfoPrintln,
	}
}

func DebugLogger(f Flog) Logger {
	return &wrapper{
		printf:  f.DebugPrintf,
		println: f.DebugPrintln,
	}
}
