package flog

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

var Default Flog = New(nil)

func CriticalPrintf(format string, args ...interface{}) {
	Default.CriticalPrintf(format, args...)
}

func CriticalPrintln(args ...interface{}) {
	Default.CriticalPrintln(args...)
}

func ErrorPrintf(format string, args ...interface{}) {
	Default.ErrorPrintf(format, args...)
}

func ErrorPrintln(args ...interface{}) {
	Default.ErrorPrintln(args...)
}

func WarningPrintf(format string, args ...interface{}) {
	Default.WarningPrintf(format, args...)
}

func WarningPrintln(args ...interface{}) {
	Default.WarningPrintln(args...)
}

func InfoPrintf(format string, args ...interface{}) {
	Default.InfoPrintf(format, args...)
}

func InfoPrintln(args ...interface{}) {
	Default.InfoPrintln(args...)
}

func DebugPrintf(format string, args ...interface{}) {
	Default.DebugPrintf(format, args...)
}

func DebugPrintln(args ...interface{}) {
	Default.DebugPrintln(args...)
}

func Sub(c *Config) Flog {
	return Default.Sub(c)
}

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
	Sub(c *Config) Flog
}

type nilWriter struct {
}

func (n *nilWriter) Write(p []byte) (int, error) {
	return len(p), nil
}

var NilWriter = &nilWriter{}

type Config struct {
	Name           string
	CriticalWriter io.Writer
	ErrorWriter    io.Writer
	WarningWriter  io.Writer
	InfoWriter     io.Writer
	DebugWriter    io.Writer
}

func resolveConfig(c *Config, f *flog) *Config {
	cfg := &Config{}
	if c != nil {
		*cfg = *c
	}
	if f != nil && f.name != "" {
		if cfg.Name != "" {
			cfg.Name = f.name + " " + cfg.Name
		} else {
			cfg.Name = f.name
		}
	}
	if cfg.CriticalWriter == nil {
		if f != nil {
			cfg.CriticalWriter = f.criticalWriter
		} else {
			cfg.CriticalWriter = os.Stderr
		}
	} else if _, ok := cfg.CriticalWriter.(*nilWriter); ok {
		cfg.CriticalWriter = nil
	}
	if cfg.ErrorWriter == nil {
		if f != nil {
			cfg.ErrorWriter = f.errorWriter
		} else {
			cfg.ErrorWriter = os.Stderr
		}
	} else if _, ok := cfg.ErrorWriter.(*nilWriter); ok {
		cfg.ErrorWriter = nil
	}
	if cfg.WarningWriter == nil {
		if f != nil {
			cfg.WarningWriter = f.warningWriter
		} else {
			cfg.WarningWriter = os.Stderr
		}
	} else if _, ok := cfg.WarningWriter.(*nilWriter); ok {
		cfg.WarningWriter = nil
	}
	if cfg.InfoWriter == nil {
		if f != nil {
			cfg.InfoWriter = f.infoWriter
		} else {
			cfg.InfoWriter = os.Stdout
		}
	} else if _, ok := cfg.InfoWriter.(*nilWriter); ok {
		cfg.InfoWriter = nil
	}
	if cfg.DebugWriter == nil {
		if f != nil {
			cfg.DebugWriter = f.debugWriter
		} else {
			cfg.DebugWriter = nil
		}
	} else if _, ok := cfg.DebugWriter.(*nilWriter); ok {
		cfg.DebugWriter = nil
	}
	return cfg
}

type flog struct {
	lock           sync.Mutex
	criticalWriter io.Writer
	errorWriter    io.Writer
	warningWriter  io.Writer
	infoWriter     io.Writer
	debugWriter    io.Writer
	buf            []byte

	name           string
	criticalFormat string
	errorFormat    string
	warningFormat  string
	infoFormat     string
	debugFormat    string
}

func New(c *Config) Flog {
	cfg := resolveConfig(c, nil)
	f := &flog{
		name:           cfg.Name,
		criticalWriter: cfg.CriticalWriter,
		errorWriter:    cfg.ErrorWriter,
		warningWriter:  cfg.WarningWriter,
		infoWriter:     cfg.InfoWriter,
		debugWriter:    cfg.DebugWriter,
	}
	if f.name == "" {
		f.criticalFormat = fmt.Sprintf("2006-01-02 15:05:05 CRITICAL ")
		f.errorFormat = fmt.Sprintf("2006-01-02 15:05:05 ERROR ")
		f.warningFormat = fmt.Sprintf("2006-01-02 15:05:05 WARNING ")
		f.infoFormat = fmt.Sprintf("2006-01-02 15:05:05 INFO ")
		f.debugFormat = fmt.Sprintf("2006-01-02 15:05:05 DEBUG ")
	} else {
		f.criticalFormat = fmt.Sprintf("2006-01-02 15:05:05 CRITICAL %s ", f.name)
		f.errorFormat = fmt.Sprintf("2006-01-02 15:05:05 ERROR %s ", f.name)
		f.warningFormat = fmt.Sprintf("2006-01-02 15:05:05 WARNING %s ", f.name)
		f.infoFormat = fmt.Sprintf("2006-01-02 15:05:05 INFO %s ", f.name)
		f.debugFormat = fmt.Sprintf("2006-01-02 15:05:05 DEBUG %s ", f.name)
	}
	return f
}

func (f *flog) CriticalPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.criticalWriter, f.criticalFormat, format, args...)
	f.lock.Unlock()
}

func (f *flog) ErrorPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.errorWriter, f.errorFormat, format, args...)
	f.lock.Unlock()
}

func (f *flog) WarningPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.warningWriter, f.warningFormat, format, args...)
	f.lock.Unlock()
}

func (f *flog) InfoPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.infoWriter, f.infoFormat, format, args...)
	f.lock.Unlock()
}

func (f *flog) DebugPrintf(format string, args ...interface{}) {
	f.lock.Lock()
	flogPrintf(f.buf, f.debugWriter, f.debugFormat, format, args...)
	f.lock.Unlock()
}

func (f *flog) CriticalPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.criticalWriter, f.criticalFormat, args...)
	f.lock.Unlock()
}

func (f *flog) ErrorPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.errorWriter, f.errorFormat, args...)
	f.lock.Unlock()
}

func (f *flog) WarningPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.warningWriter, f.warningFormat, args...)
	f.lock.Unlock()
}

func (f *flog) InfoPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.infoWriter, f.infoFormat, args...)
	f.lock.Unlock()
}

func (f *flog) DebugPrintln(args ...interface{}) {
	f.lock.Lock()
	flogPrintln(f.buf, f.debugWriter, f.debugFormat, args...)
	f.lock.Unlock()
}

func flogPrintf(buf []byte, out io.Writer, prefixFormat string, format string, args ...interface{}) {
	if out == nil {
		return
	}
	buf = time.Now().AppendFormat(buf[:0], prefixFormat)
	bufr := bytes.NewBuffer(buf)
	fmt.Fprintf(bufr, format, args...)
	buf = bufr.Bytes()
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
}

func flogPrintln(buf []byte, out io.Writer, prefixFormat string, args ...interface{}) {
	if out == nil {
		return
	}
	buf = time.Now().AppendFormat(buf[:0], prefixFormat)
	bufr := bytes.NewBuffer(buf)
	fmt.Fprintln(bufr, args...)
	buf = bufr.Bytes()
	if len(buf) == 0 || buf[len(buf)-1] != '\n' {
		buf = append(buf, '\n')
	}
	out.Write(buf)
}

func (f *flog) Sub(c *Config) Flog {
	cfg := resolveConfig(c, f)
	return New(cfg)
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
