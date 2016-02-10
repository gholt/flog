package simplelogger

import (
	"os"
	"testing"
)

func BenchmarkDebugPrintf(b *testing.B) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	defer devNull.Close()
	logger := New("benchmark", devNull, devNull)
	for n := 0; n < b.N; n++ {
		logger.DebugPrintf("testing log line %016x with some formatting %p", n, logger)
	}
}

func BenchmarkDebugPrintf2(b *testing.B) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	defer devNull.Close()
	logger := New("benchmark", devNull, devNull)
	for n := 0; n < b.N; n++ {
		logger.DebugPrintf2("testing log line %016x with some formatting %p", n, logger)
	}
}

func BenchmarkDebugPrintf3(b *testing.B) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	defer devNull.Close()
	logger := New("benchmark", devNull, devNull)
	for n := 0; n < b.N; n++ {
		logger.DebugPrintf3("testing log line %016x with some formatting %p", n, logger)
	}
}

func BenchmarkDebugPrintf4(b *testing.B) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	defer devNull.Close()
	logger := New("benchmark", devNull, devNull)
	for n := 0; n < b.N; n++ {
		logger.DebugPrintf4("testing log line %016x with some formatting %p", n, logger)
	}
}

func BenchmarkDebugPrintf5(b *testing.B) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	defer devNull.Close()
	logger := New("benchmark", devNull, devNull)
	for n := 0; n < b.N; n++ {
		logger.DebugPrintf5("testing log line %016x with some formatting %p", n, logger)
	}
}

func BenchmarkDebugPrintf6(b *testing.B) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	defer devNull.Close()
	logger := New("benchmark", devNull, devNull)
	for n := 0; n < b.N; n++ {
		logger.DebugPrintf6("testing log line %016x with some formatting %p", n, logger)
	}
}

func BenchmarkDebugPrintf7(b *testing.B) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	defer devNull.Close()
	logger := New("benchmark", devNull, devNull)
	for n := 0; n < b.N; n++ {
		logger.DebugPrintf7("testing log line %016x with some formatting %p", n, logger)
	}
}

func BenchmarkDebugPrintf8(b *testing.B) {
	devNull, err := os.Open(os.DevNull)
	if err != nil {
		b.Fatal(err)
	}
	defer devNull.Close()
	logger := New("benchmark", devNull, devNull)
	for n := 0; n < b.N; n++ {
		logger.DebugPrintf8("testing log line %016x with some formatting %p", n, logger)
	}
}
