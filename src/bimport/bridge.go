package bimport

import "vk-film-library/internal/bridge"

type Bridge struct {
	Info bridge.Info
}

type TestBridge struct {
	Info *bridge.MockInfo
}
