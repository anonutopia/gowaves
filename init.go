package gowaves

import "log"

var WNC *WavesNodeClient

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	WNC = initWNC()
}
