package main

import (
	"C"
	sharedfrp "github.com/fatedier/frp/sharefrp/mobile"
)

//export LoadServerAssets
func LoadServerAssets() {
	sharedfrp.LoadServerAssets()
}

//export RunServer
func RunServer(iniContent *C.char) C.int {
	err := sharedfrp.RunServer(C.GoString(iniContent))
	if err == nil {
		return C.int(1)
	} else {
		return C.int(0)
	}
}

//export GetServerStatus
func GetServerStatus() C.int {
	return C.int(sharedfrp.GetServerStatus())
}
