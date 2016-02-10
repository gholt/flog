package flog

import (
	//"testing"

	"google.golang.org/grpc/grpclog"
)

var writers FlogWriters = Std.(FlogWriters)
var gprcLogger grpclog.Logger = DebugLogger(Std)
