package bimport

import "grpc-test/internal/bridge"

type Bridge struct {
	Info bridge.Info
}

type TestBridge struct {
	Info *bridge.MockInfo
}
