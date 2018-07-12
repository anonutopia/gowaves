package gowaves

import (
// "log"
)

var WNC *WavesNodeClient

func init() {
	WNC = &WavesNodeClient{}

	WNC.Host = DEFAULT_HOST
	WNC.Port = DEFAULT_PORT
}
