package gowaves

import "log"

var WNC *WavesNodeClient

var WMC *WavesMatcherClient

func init() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	WNC = initWNC()
	WMC = initWMC()
}
